package main

import (
	"bytes"

	"fmt"
	"log"

	"github.com/sanity-io/litter"
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

	fmt.Println("NextReviewAt() 1: ", card.Algorithm.NextReviewAt())

	// TODO: this is where we rely on user input for them to rate this card
	// then we call advance based on user input.
	// Remember to validate the user input
	// review and rate a card
	sm := card.Algorithm.Advance(0.8)
	card.Algorithm = sm
	fmt.Println("NextReviewAt() 2: ", card.Algorithm.NextReviewAt())

	// After the user has rated the card and we have updated the card struct with the new metadata
	// We need to  persist that on the markdown files' extended attributes
	// update the card attributes with new algo
	var network bytes.Buffer
	err = card.Encode(&network)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	err = setExtendedAttrs(filepath, network.Bytes())
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	fmt.Println("card when saving")
	litter.Dump(card)

	// TODO: finally we display the answer to the user
	err = card.Display()
	if err != nil {
		log.Fatalf("error: %+v", err)

	}

}
