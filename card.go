package main

import (
	"encoding/gob"
	"io"
	"os"
	"os/exec"

	"github.com/gomarkdown/markdown/ast"
	"github.com/pkg/errors"
	"github.com/pkg/xattr"
)

// has to start with "user."
const attrName = "user.algo"

// Card represents a single card in a Deck.
type Card struct {
	Version   uint32
	Question  string
	FilePath  string
	Algorithm SRSalgorithm
}

// Encode encodes the Card value into the encoder.
func (c Card) Encode(w io.Writer) error {
	// The Card.Algorithm concrete type(eg Supermemo2) has to be registered
	// using gob.Register else this function will fail
	// We registered it in main.

	enc := gob.NewEncoder(w)
	err := enc.Encode(&c)
	if err != nil {
		return errors.Wrapf(err, "unable to encode card %v", c)
	}
	return nil
}

// Decode decodes the next Card value from the stream and returns it.
func (c Card) Decode(r io.Reader) (*Card, error) {
	dec := gob.NewDecoder(r)
	err := dec.Decode(&c)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode card %v", c)
	}
	return &c, nil
}

// Display shows cards content to terminal
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

func setExtendedAttrs(filepath string, algoEncoded []byte) error {
	err := xattr.Set(filepath, attrName, algoEncoded)
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
