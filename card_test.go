package srs

import (
	"testing"
	"time"
)

//TODO: We should remove all file extended attributes before every test

func TestCardCreation(t *testing.T) {

	cardPath := "./testdata/test1.md"
	_, err := NewCard(cardPath)
	if err != nil {
		t.Errorf("\nCalled NewCard(%#+v) \ngot %s \nwanted %#+v", cardPath, err, nil)
	}

}

func TestCardReviewAt(t *testing.T) {
	cardPath := "./testdata/test1.md"
	card, err := NewCard(cardPath)
	if err != nil {
		t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", cardPath, err, nil)
	}

	now := time.Now()
	nextReview := card.Algorithm.NextReviewAt()

	if nextReview.Day() != now.Day() {
		t.Errorf("\nCalled card.Algorithm.NextReviewAt() \ngot %#+v \nwanted %#+v", nextReview.Day(), now.Day())
	}

}

func TestCardAdvance(t *testing.T) {

	t.Run("rating less than ratingSuccess", func(t *testing.T) {
		// rating less than `ratingSuccess`
		rating := 0.4
		cardPath := "./testdata/test1.md"
		card, err := NewCard(cardPath)
		if err != nil {
			t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", cardPath, err, nil)
		}

		card.Rate(rating)
		now := time.Now()
		nextReview := card.Algorithm.NextReviewAt()
		if nextReview.Day() != now.Day()+1 {
			t.Errorf("\nCalled card.Algorithm.Rate() \ngot %#+v \nwanted %#+v", nextReview.Day(), now.Day()+1)
		}

	})

	t.Run("rating greater than ratingSuccess", func(t *testing.T) {
		// rating greater than `ratingSuccess`
		rating := 0.9
		cardPath := "./testdata/test1.md"
		card, err := NewCard(cardPath)
		if err != nil {
			t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", cardPath, err, nil)
		}

		card.Rate(rating)
		now := time.Now()
		nextReview := card.Algorithm.NextReviewAt()
		if nextReview.Day() != now.Day()+1 {
			t.Errorf("\nCalled card.Algorithm.Rate() \ngot %#+v \nwanted %#+v", nextReview.Day(), now.Day()+1)
		}

	})

	t.Run("multiple rating less than ratingSuccess", func(t *testing.T) {
		// multiple rating less than `ratingSuccess`
		rating := 0.4
		cardPath := "./testdata/test1.md"
		card, err := NewCard(cardPath)
		if err != nil {
			t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", cardPath, err, nil)
		}

		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)

		now := time.Now()
		nextReview := card.Algorithm.NextReviewAt()
		if nextReview.Day() != now.Day()+1 {
			t.Errorf("\nCalled card.Algorithm.Rate() \ngot %#+v \nwanted %#+v", nextReview.Day(), now.Day()+1)
		}

	})

	t.Run("multiple rating greater than ratingSuccess", func(t *testing.T) {
		// multiple rating greater than `ratingSuccess`
		rating := 0.8
		cardPath := "./testdata/test1.md"
		card, err := NewCard(cardPath)
		if err != nil {
			t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", cardPath, err, nil)
		}

		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)
		card.Rate(rating)

		now := time.Now()
		nextReview := card.Algorithm.NextReviewAt()
		if nextReview.Year() != now.Year()+10 {
			t.Errorf("\nCalled card.Algorithm.Rate() \ngot %#+v \nwanted %#+v", nextReview.Year(), now.Year()+10)
		}
	})
}
