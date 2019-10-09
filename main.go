package main

import (
	"bufio"
	"os"
	"strconv"

	"fmt"
	"log"

	"github.com/pkg/errors"
)

/*
run as:
    go run .
*/

func init() {
	// init funcs are bad
	AlgoRegistration()
}

func main() {
	// TODO: accept a directory and loop through all the markdown files in that directory
	filepath := "/Users/komuw/mystuff/srs/pol.md"
	card, err := NewCard(filepath)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

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
