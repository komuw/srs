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

	"github.com/komuw/srs"
)

/*
run as:
    go run cmd/main.go -d myCards/
*/

func init() {
	// init funcs are bad
	srs.AlgoRegistration()
}

func main() {
	var cardDir string
	flag.StringVar(
		&cardDir,
		"d",
		"",
		"path to directory containing cards.")
	flag.Parse()

	cardDirAbs, err := filepath.Abs(cardDir)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
	deck := srs.NewDeck()
	err = filepath.Walk(cardDirAbs, walkFnClosure(cardDirAbs, deck))
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	for k, card := range deck.Cards {
		divider := fmt.Sprintf("\n\t##################### question: %d #####################\n", k+1)

		fmt.Printf(divider)
		fmt.Printf("\n\t %s \n\n", card.Question)
		fmt.Print("Rate your answer between 1-10:")

		uInput, err := getuserInput()
		if err != nil {
			log.Fatalf("error: %+v", err)
		}

		// rate card
		card.Rate(uInput)

		// persist new metadata
		err = card.Encode()
		if err != nil {
			log.Fatalf("error: %+v", err)
		}

		// display the answer to the user
		err = card.Display()
		if err != nil {
			log.Fatalf("error: %+v", err)

		}
		fmt.Printf("The next reviewed is at: %s", card.Algorithm.NextReviewAt().Format("02 Jan 2006"))
		fmt.Printf(divider)
		time.Sleep(3 * time.Second)
	}

}

func walkFnClosure(src string, deck *srs.Deck) filepath.WalkFunc {
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

		card, err := srs.NewCard(path)
		if err != nil {
			return err
		}
		deck.Cards = append(deck.Cards, *card)

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

			userInput = float64(uInputInt)
			break
		}
	}
	err = scanner.Err()
	if err != nil {
		return userInput, errors.Wrapf(err, "scanner error")
	}

	return userInput, nil
}
