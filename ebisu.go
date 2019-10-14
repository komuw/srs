package srs

import (
	"math"
	"time"
)

type model struct {
	Alpha float64
	Beta  float64
	T     float64
}

// Ebisu implements ebisu SSR algorithm.
type Ebisu struct {
	LastReviewedAt time.Time
	Alpha          float64
	Beta           float64
	Interval       float64
}

// NewEbisu consturcts a new Ebisu instance.
func NewEbisu() Ebisu {
	return Ebisu{
		LastReviewedAt: time.Now().Add(-24 * time.Hour),
		Alpha:          3,
		Beta:           3,
		Interval:       24,
	}
}

// NextReviewAt returns next review timestamp for a card.
func (eb Ebisu) NextReviewAt() time.Time {
	return eb.LastReviewedAt.Add(time.Duration(eb.Interval) * time.Hour)
}

// Advance advances supermemo state for a card.
func (eb Ebisu) Advance(rating float64) SRSalgorithm {
	newEb := eb

	model := &model{newEb.Alpha, newEb.Beta, newEb.Interval}
	elapsed := float64(time.Since(newEb.LastReviewedAt)) / float64(time.Hour)
	proposed := updateRecall(model, rating >= ratingSuccess, float64(elapsed), true, newEb.Interval)

	newEb.Alpha = proposed.Alpha
	newEb.Beta = proposed.Beta
	newEb.Interval = proposed.T
	newEb.LastReviewedAt = time.Now()

	return newEb
}

func rebalanceModel(prior *model, result bool, tnow float64, proposed *model) *model {
	if proposed.Alpha > 2*proposed.Beta || proposed.Beta > 2*proposed.Alpha {
		roughHalflife := modelToPercentileDecay(proposed, 0.5)
		return updateRecall(prior, result, tnow, false, roughHalflife)
	}

	return proposed
}

func updateRecall(prior *model, result bool, tnow float64, rebalance bool, tback float64) *model {
	dt := tnow / prior.T
	et := tnow / tback

	var sig2, mean float64
	if result {
		if tback == prior.T {
			proposed := &model{prior.Alpha + dt, prior.Beta, prior.T}
			if rebalance {
				return rebalanceModel(prior, result, tnow, proposed)
			}

			return proposed
		}

		logDenominator := betaln(prior.Alpha+dt, prior.Beta)
		logmean := betaln(prior.Alpha+dt/et*(1+et), prior.Beta) - logDenominator
		logm2 := betaln(prior.Alpha+dt/et*(2+et), prior.Beta) - logDenominator
		mean = math.Exp(logmean)
		sig2 = subexp(logm2, 2*logmean)
	} else {
		logDenominator := logsumexp(
			[2]float64{betaln(prior.Alpha, prior.Beta), betaln(prior.Alpha+dt, prior.Beta)},
			[2]float64{1, -1},
		)
		mean = subexp(
			betaln(prior.Alpha+dt/et, prior.Beta)-logDenominator,
			betaln(prior.Alpha+(dt/et)*(et+1), prior.Beta)-logDenominator,
		)
		m2 := subexp(
			betaln(prior.Alpha+2*dt/et, prior.Beta)-logDenominator,
			betaln(prior.Alpha+dt/et*(et+2), prior.Beta)-logDenominator,
		)

		if m2 <= 0 {
			panic("invalid second moment found")
		}

		sig2 = m2 - math.Pow(mean, 2)
	}

	if mean <= 0 {
		panic("invalid mean found")
	}
	if sig2 <= 0 {
		panic("invalid variance found")
	}

	newAlpha, newBeta := meanVarToBeta(mean, sig2)
	proposed := &model{newAlpha, newBeta, tback}
	if rebalance {
		return rebalanceModel(prior, result, tnow, proposed)
	}

	return proposed
}

func modelToPercentileDecay(model *model, percentile float64) float64 {
	if percentile < 0 || percentile > 1 {
		panic("percentiles must be between (0, 1) exclusive")
	}
	alpha := model.Alpha
	beta := model.Beta
	t0 := model.T

	logBab := betaln(alpha, beta)
	logPercentile := math.Log(percentile)
	f := func(lndelta float64) float64 {
		logMean := betaln(alpha+math.Exp(lndelta), beta) - logBab
		return logMean - logPercentile
	}

	bracketWidth := 1.0
	blow := -bracketWidth / 2.0
	bhigh := bracketWidth / 2.0
	flow := f(blow)
	fhigh := f(bhigh)
	for {
		if flow < 0 || fhigh < 0 {
			break
		}

		// Move the bracket up.
		blow = bhigh
		flow = fhigh
		bhigh += bracketWidth
		fhigh = f(bhigh)
	}

	for {
		if flow > 0 || fhigh > 0 {
			break
		}

		// Move the bracket down.
		bhigh = blow
		fhigh = flow
		blow -= bracketWidth
		flow = f(blow)
	}

	if !(flow > 0 && fhigh < 0) {
		panic("failed to bracket")
	}

	return (math.Exp(blow) + math.Exp(bhigh)) / 2 * t0
}

func meanVarToBeta(mean, v float64) (float64, float64) {
	tmp := mean*(1-mean)/v - 1
	return mean * tmp, (1 - mean) * tmp
}

func subexp(x, y float64) float64 {
	maxval := math.Max(x, y)
	return math.Exp(maxval) * (math.Exp(x-maxval) - math.Exp(y-maxval))
}

func logsumexp(a, b [2]float64) float64 {
	aMax := math.Max(a[0], a[1])
	sum := b[0] * math.Exp(a[0]-aMax)
	sum += b[1] * math.Exp(a[1]-aMax)
	return math.Log(sum) + aMax
}

// betaln returns natural logarithm of the Beta function.
func betaln(a, b float64) float64 {
	// B(x,y) = Γ(x)Γ(y) / Γ(x+y)
	// Therefore log(B(x,y)) = log(Γ(x)) + log(Γ(y)) - log(Γ(x+y))
	la, _ := math.Lgamma(a)
	lb, _ := math.Lgamma(b)
	lab, _ := math.Lgamma(a + b)
	return la + lb - lab
}
