package main

import (
	"encoding/json"

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
func main() {
	// TODO: accept a directory and loop through all the markdown files in that directory
	filepath := "/home/komuw/mystuff/srs/pol.md"
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
		err = json.Unmarshal(cardAttribute, &card)
		if err != nil {
			log.Fatalf("error: %+v", err)
		}
		fmt.Println("card from file")
		litter.Dump(card)
	}

	fmt.Println("NextReviewAt() 1: ", card.Algorithm.NextReviewAt())
	// review and rate a card
	sm := card.Algorithm.Advance(0.8)
	card.Algorithm = sm
	fmt.Println("NextReviewAt() 2: ", card.Algorithm.NextReviewAt())

	// update the card attributes with new algo
	algoJson, err := json.Marshal(card)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
	err = setExtendedAttrs(filepath, algoJson)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	fmt.Println("card when saving")
	litter.Dump(card)

	err = card.Display()
	if err != nil {
		log.Fatalf("error: %+v", err)

	}

}
