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

	filename := "test1.md"
	cardDir := "./testdata/"
	_, err = NewCard(filename, cardDir, db)
	// NewCard(filename string, cardDir string, db *bbolt.DB)
	if err != nil {
		t.Errorf("\nCalled NewCard(%#+v) \ngot %s \nwanted %#+v", filename, err, nil)
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

	filename := "test1.md"
	cardDir := "./testdata/"
	card, err := NewCard(filename, cardDir, db)
	if err != nil {
		t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", filename, err, nil)
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
		filename := "test1.md"
		cardDir := "./testdata/"
		card, err := NewCard(filename, cardDir, db)
		if err != nil {
			t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", filename, err, nil)
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
		filename := "test1.md"
		cardDir := "./testdata/"
		card, err := NewCard(filename, cardDir, db)
		if err != nil {
			t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", filename, err, nil)
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
	// 	filename := "test1.md"
	// 	card, err := NewCard(filename)
	// 	if err != nil {
	// 		t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", filename, err, nil)
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
	// 	filename := "test1.md"
	// 	card, err := NewCard(filename)
	// 	if err != nil {
	// 		t.Errorf("\nCalled NewCard(%#+v) \ngot %#+v \nwanted %#+v", filename, err, nil)
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
