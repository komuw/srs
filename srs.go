package main

import "time"

// SRSalgorithm is the interface that all posible srs algo's should satisfy
type SRSalgorithm interface {
	NextReviewAt() time.Time
	Advance(rating float64) SRSalgorithm
}
