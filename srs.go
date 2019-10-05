package main

import "time"

type SRSalgorithm interface {
	NextReviewAt() time.Time
	Advance(rating float64) SRSalgorithm
}
