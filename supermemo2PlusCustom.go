package srs

import (
	"math"
	"time"
)

// Supermemo2PlusCustom calculates review intervals using altered SM2+ algorithm
type Supermemo2PlusCustom struct {
	LastReviewedAt time.Time
	Difficulty     float64
	Interval       float64
}

// NewSupermemo2PlusCustom returns a new Supermemo2PlusCustom instance
// func NewSupermemo2PlusCustom() Supermemo2PlusCustom {
// 	return Supermemo2PlusCustom{
// 		LastReviewedAt: time.Now().Add(-4 * time.Hour),
// 		Difficulty:     0.3,
// 		Interval:       0.2,
// 	}
// }

// NextReviewAt returns next review timestamp for a card.
func (sm Supermemo2PlusCustom) NextReviewAt() time.Time {
	return sm.LastReviewedAt.Add(time.Duration(24*sm.Interval) * time.Hour)
}

// Advance advances supermemo state for a card.
func (sm Supermemo2PlusCustom) Advance(rating float64) SRSalgorithm {
	newSm := sm
	success := rating >= ratingSuccess
	percentOverdue := float64(1)
	if success {
		percentOverdue = newSm.PercentOverdue()
	}

	newSm.Difficulty += percentOverdue / 35 * (8 - 9*rating)
	newSm.Difficulty = math.Max(0, math.Min(1, newSm.Difficulty))
	difficultyWeight := 3.5 - 1.7*newSm.Difficulty

	minInterval := math.Min(1.0, newSm.Interval)
	factor := minInterval / math.Pow(difficultyWeight, 2)
	if success {
		minInterval = 0.2
		factor = minInterval + (difficultyWeight-1)*percentOverdue
	}

	newSm.LastReviewedAt = time.Now()
	newSm.Interval = math.Max(minInterval, math.Min(newSm.Interval*factor, 300))
	return newSm
}

// PercentOverdue returns corresponding SM2+ value for a Card.
func (sm Supermemo2PlusCustom) PercentOverdue() float64 {
	percentOverdue := time.Since(sm.LastReviewedAt).Hours() / float64(24*sm.Interval)
	return math.Min(2, percentOverdue)
}
