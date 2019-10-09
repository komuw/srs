package srs

import (
	"math"
	"time"
)

// Supermemo2 calculates review intervals using SM2 algorithm
type Supermemo2 struct {
	LastReviewedAt time.Time
	Interval       float64
	Easiness       float64
	Correct        int
	Total          int
}

// NewSupermemo2 returns a new Supermemo2 instance
func NewSupermemo2() Supermemo2 {
	return Supermemo2{
		LastReviewedAt: time.Now(),
		Interval:       0,
		Easiness:       2.5,
		Correct:        0,
		Total:          0,
	}
}

// NextReviewAt returns next review timestamp for a card.
func (sm Supermemo2) NextReviewAt() time.Time {
	return sm.LastReviewedAt.Add(time.Duration(24*sm.Interval) * time.Hour)
}

// Advance advances supermemo state for a card.
func (sm Supermemo2) Advance(rating float64) SRSalgorithm {
	newSm := sm
	newSm.Total++
	newSm.LastReviewedAt = time.Now()
	newSm.Easiness += 0.1 - (1-rating)*(0.4+(1-rating)*0.5)
	newSm.Easiness = math.Max(newSm.Easiness, 1.3)

	interval := 1.0
	if rating >= ratingSuccess {
		if newSm.Total == 2 {
			interval = 6
		} else if newSm.Total > 2 {
			interval = math.Round(newSm.Interval * newSm.Easiness)
		}
		newSm.Correct++
	} else {
		newSm.Correct = 0
	}
	newSm.Interval = interval

	return newSm
}
