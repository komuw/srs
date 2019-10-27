package srs

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func tempfile() string {
	f, err := ioutil.TempFile("", "srs-testDB-")
	if err != nil {
		panic(err)
	}
	if err := f.Close(); err != nil {
		panic(err)
	}
	if err := os.Remove(f.Name()); err != nil {
		panic(err)
	}
	return f.Name()
}

func dbRemoverUtil(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

func TestCardCreation(t *testing.T) {
	dbPath := tempfile()
	defer dbRemoverUtil(dbPath)
	db, err := OpenDb(dbPath)
	if err != nil {
		log.Fatal("unable to open test db at:: ", dbPath, err)
	}
	defer db.Close()

	cardPath := "./testdata/test1.md"
	_, err = NewCard(cardPath, db)
	if err != nil {
		t.Errorf("\nCalled NewCard(%#+v) \ngot %s \nwanted %#+v", cardPath, err, nil)
	}

}

func TestCardReviewAt(t *testing.T) {
	dbPath := tempfile()
	defer dbRemoverUtil(dbPath)
	db, err := OpenDb(dbPath)
	if err != nil {
		log.Fatal("unable to open test db at:: ", dbPath, err)
	}
	defer db.Close()

	cardPath := "./testdata/test1.md"
	card, err := NewCard(cardPath, db)
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
	dbPath := tempfile()
	defer dbRemoverUtil(dbPath)
	db, err := OpenDb(dbPath)
	if err != nil {
		log.Fatal("unable to open test db at:: ", dbPath, err)
	}
	defer db.Close()

	t.Run("rating less than ratingSuccess", func(t *testing.T) {
		// rating less than `ratingSuccess`
		rating := 0.4
		cardPath := "./testdata/test1.md"
		card, err := NewCard(cardPath, db)
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
		card, err := NewCard(cardPath, db)
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

	// t.Run("multiple rating less than ratingSuccess", func(t *testing.T) {
	// 	// multiple rating less than `ratingSuccess`
	// 	rating := 0.4
	// 	cardPath := "./testdata/test1.md"
	// 	card, err := NewCard(cardPath)
	// 	if err != nil {
	// 		t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", cardPath, err, nil)
	// 	}

	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)

	// 	now := time.Now()
	// 	nextReview := card.Algorithm.NextReviewAt()
	// 	if nextReview.Day() != now.Day()+1 {
	// 		t.Errorf("\nCalled card.Algorithm.Rate() \ngot %#+v \nwanted %#+v", nextReview.Day(), now.Day()+1)
	// 	}

	// })

	// t.Run("multiple rating greater than ratingSuccess", func(t *testing.T) {
	// 	// multiple rating greater than `ratingSuccess`
	// 	rating := 0.8
	// 	cardPath := "./testdata/test1.md"
	// 	card, err := NewCard(cardPath)
	// 	if err != nil {
	// 		t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", cardPath, err, nil)
	// 	}

	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)
	// 	card.Rate(rating)

	// 	now := time.Now()
	// 	nextReview := card.Algorithm.NextReviewAt()
	// 	if nextReview.Year() != now.Year()+10 {
	// 		t.Errorf("\nCalled card.Algorithm.Rate() \ngot %#+v \nwanted %#+v", nextReview.Year(), now.Year()+10)
	// 	}
	// })
}
