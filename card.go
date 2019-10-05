package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"time"

	"github.com/gomarkdown/markdown/ast"
	"github.com/pkg/errors"
	"github.com/pkg/xattr"
)

const attrName = "user.algo" // has to start with "user."

// Card represents a single card in a Deck.
type Card struct {
	Version  uint32
	Question string
	// FileContents []byte
	FilePath  string
	Algorithm SRSalgorithm
}

// UnmarshalJSON implements json.Unmarshaler for Supermemo2
func (c *Card) UnmarshalJSON(b []byte) error {
	var objMap map[string]interface{}
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return errors.Wrapf(err, "unable to Unmarshal")
	}

	c.Version = uint32(objMap["Version"].(float64))
	c.Question = objMap["Question"].(string)
	c.FilePath = objMap["FilePath"].(string)

	var objMapAlgo = objMap["Algorithm"].(map[string]interface{})
	myAlg := NewSupermemo2()
	myAlg.Interval = objMapAlgo["Interval"].(float64)
	myAlg.Easiness = objMapAlgo["Easiness"].(float64)
	myAlg.Correct = int(objMapAlgo["Correct"].(float64))
	myAlg.Total = int(objMapAlgo["Total"].(float64))
	LastReviewedAtLayout := "2006-01-02T15:04:05Z07:00"
	ReviewedAt, err := time.Parse(LastReviewedAtLayout, objMapAlgo["LastReviewedAt"].(string))
	if err != nil {
		return errors.Wrapf(err, "unable to Parse LastReviewedAt")
	}
	myAlg.LastReviewedAt = ReviewedAt

	c.Algorithm = myAlg
	return nil
}

func (c Card) Display() error {
	cmd := exec.Command("bat", "-p", c.FilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return errors.Wrapf(err, "unable to open %v for reading", c.FilePath)
	}
	return nil
}

func getQuestion(node ast.Node) (string, error) {
	for _, child := range node.GetChildren() {
		switch thisNode := child.(type) {
		case *ast.Heading:
			question := thisNode.HeadingID
			return question, nil
		default:
			// unknown Node
		}
	}
	return "", errors.New("The markdown file does not contain a question")
}

func setExtendedAttrs(filepath string, algoJson []byte) error {
	err := xattr.Set(filepath, attrName, algoJson)
	if err != nil {
		return errors.Wrapf(err, "unable to set extended file attributes")
	}
	return nil
}

func getExtendedAttrs(filepath string) ([]byte, error) {
	attribute, err := xattr.Get(filepath, attrName)
	if len(attribute) > 0 && err != nil {
		return []byte(""), errors.Wrapf(err, "unable to get extended file attributes")
	}
	return attribute, nil
}
