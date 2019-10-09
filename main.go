package main

import (
	"bytes"

	"fmt"
	"io/ioutil"
	"log"

	"github.com/gomarkdown/markdown/parser"
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
	md, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	mainNode := parser.Parse(md)
	question, err := getQuestion(mainNode)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
	cardAttribute, err := getExtendedAttrs(filepath)
	if err != nil {
		log.Fatalf("error: %+v", err)

	}
	fmt.Println("cardAttribute:")
	litter.Dump(string(cardAttribute))

	// if cardAttribute exists, then this is not a new card and we should
	// bootstrap the Algorithm to use from the cardAttribute
	// else, create a card with a new Algorithm
	card := Card{
		Version:   1,
		Question:  question,
		FilePath:  filepath,
		Algorithm: NewSupermemo2(),
	}
	if len(cardAttribute) > 0 {
		newCard, err := card.Decode(
			bytes.NewReader(cardAttribute))
		if err != nil {
			log.Fatalf("error: %+v", err)
		}
		fmt.Println("card from file")
		litter.Dump(newCard)

		card = *newCard
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
