package main

import (
	"bufio"
	"flag"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"fmt"
	"log"

	"github.com/pkg/errors"
	"go.etcd.io/bbolt"

	"github.com/komuw/srs"
)

/*
run as:
    go run cmd/main.go -d myCards/
*/

var version string

func init() {
	// init funcs are bad
	srs.AlgoRegistration()
}

func main() {
	var showVersion bool
	var cardDir string
	var dbPath string
	flag.BoolVar(
		&showVersion,
		"v",
		false,
		"Show version information.")
	flag.StringVar(
		&cardDir,
		"d",
		"",
		"path to directory containing cards.")
	flag.StringVar(
		&dbPath,
		"db",
		"",
		"db path.")
	flag.Parse()

	if showVersion {
		fmt.Println("Srs version: ", version)
		os.Exit(0)
	}
	if cardDir == "" {
		fmt.Println("You should provide a path to your cards collection")
		os.Exit(1)
	}
	if dbPath == "" {
		fmt.Println("You should provide a path to your db")
		os.Exit(1)
	}

	db, err := srs.OpenDb(dbPath)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
	defer db.Close()

	cardDirAbs, err := filepath.Abs(cardDir)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
	deck := srs.NewDeck()
	err = filepath.Walk(cardDirAbs, walkFnClosure(cardDirAbs, deck, db))
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	if len(deck.Cards) == 0 {
		fmt.Printf("There are no cards to review today in: %s\n", cardDirAbs)
	}
	for k, card := range deck.Cards {
		newCard := card //iteration bug
		divider := fmt.Sprintf("\n\t##################### question: %d #####################", k+1)

		fmt.Println(divider)
		fmt.Printf("\n\t %s \n\n", newCard.Question)
		fmt.Print("Rate your answer between 1-10:")

		uInput, err := getuserInput()
		if err != nil {
			log.Fatalf("error: %+v", err)
		}

		// rate card
		newCard.Rate(uInput)

		// persist new metadata
		err = newCard.Encode(db)
		if err != nil {
			log.Fatalf("error: %+v", err)
		}

		// display the answer to the user
		err = newCard.Display()
		if err != nil {
			log.Fatalf("error: %+v", err)

		}
		fmt.Printf("The next reviewed is at: %s", newCard.Algorithm.NextReviewAt().Format("02 Jan 2006"))
		fmt.Println(divider)
		time.Sleep(3 * time.Second)
	}

}

func walkFnClosure(src string, deck *srs.Deck, db *bbolt.DB) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// todo: maybe we should return nil
			return err
		}

		if info.Mode().IsDir() {
			// return on
			return nil
		}
		if !info.Mode().IsRegular() {
			// return on non-regular files
			return nil
		}
		if strings.ToLower(filepath.Ext(path)) != ".md" {
			// non-markdown
			// TODO: support other markdown extensions like .mkd
			return nil
		}

		card, err := srs.NewCard(path, db)
		if err != nil {
			return err
		}

		// If the next review date for a card is not today;
		// it should not be added to the deck.
		// However if next review date is in the past;
		// it should be added.
		now := time.Now()
		nextReview := card.Algorithm.NextReviewAt()
		if now.Sub(nextReview) >= 0 {
			// the duration `now - nextReview`
			deck.Cards = append(deck.Cards, *card)
		}
		return nil
	}
}

func getuserInput() (float64, error) {
	var userInput float64
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInputStr := scanner.Text()
		if len(userInputStr) > 0 {
			fmt.Println("you entered: ", userInputStr)
			uInputInt, err := strconv.Atoi(userInputStr)
			if err != nil {
				err = errors.New("user input should be between 1-10, try again")
				log.Println(err)
				continue
			}
			if uInputInt < 0 {
				err = errors.New("user input should be between 1-10, try again")
				log.Println(err)
				continue
			} else if uInputInt > 10 {
				err = errors.New("user input should be between 1-10, try again")
				log.Println(err)
				continue
			}

			// rating should be between 0.1 - 1.0
			userInput = float64(uInputInt) / 10
			break
		}
	}
	err = scanner.Err()
	if err != nil {
		return userInput, errors.Wrapf(err, "scanner error")
	}

	return userInput, nil
}
