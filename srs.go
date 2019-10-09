package srs

import (
	"encoding/gob"
	"time"
)

// SRSalgorithm is the interface that all posible srs algo's should satisfy
type SRSalgorithm interface {
	NextReviewAt() time.Time
	Advance(rating float64) SRSalgorithm
}

// AlgoRegistration registers all SRS algos for gob encoding/decoding
func AlgoRegistration() {
	gob.Register(Supermemo2{})
}
