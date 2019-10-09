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

const attrName = "user.algo" // has to start with "user."

// Card represents a single card in a Deck.
type Card struct {
	Version  uint32
	Question string
	// FileContents []byte
	FilePath  string
	Algorithm SRSalgorithm
}

// Encode encodes the interface value into the encoder.
func (c Card) Encode(w io.Writer) error {
	// The encode will fail unless the concrete type has been
	// registered. We registered it in the calling function.

	// Pass pointer to interface so Encode sees (and hence sends) a value of
	// interface type. If we passed p directly it would see the concrete type instead.
	// See the blog post, "The Laws of Reflection" for background.

	enc := gob.NewEncoder(w)
	err := enc.Encode(&c)
	if err != nil {
		return errors.Wrapf(err, "unable to encode card %v", c)
	}
	return nil
}

// Decode decodes the next interface value from the stream and returns it.
func (c Card) Decode(r io.Reader) (*Card, error) {
	// The decode will fail unless the concrete type on the wire has been
	// registered. We registered it in the calling function.

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

func setExtendedAttrs(filepath string, algoJSON []byte) error {
	err := xattr.Set(filepath, attrName, algoJSON)
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
