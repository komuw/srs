package srs

import (
	"bytes"
	"encoding/gob"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/pkg/errors"
	"github.com/pkg/xattr"
)

// has to start with "user."
const attrName = "user.algo"

// Deck represents a collection of the cards to review.
type Deck struct {
	Cards []Card
}

// NewDeck creates a new pack of Cards
func NewDeck() *Deck {
	return &Deck{}
}

// Card represents a single card in a Deck.
type Card struct {
	Version   uint32
	Question  string
	FilePath  string
	Algorithm SRSalgorithm
}

// NewCard returns a new Card
func NewCard(filepath string) (*Card, error) {
	md, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read file %v", filepath)

	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	mainNode := parser.Parse(md)
	question, err := getQuestion(mainNode)
	if err != nil {
		// the error is already annotated
		return nil, err
	}
	cardAttribute, err := getExtendedAttrs(filepath)
	if err != nil {
		return nil, err
	}

	card := &Card{
		Version:   1,
		Question:  question,
		FilePath:  filepath,
		Algorithm: NewSupermemo2(),
	}
	if len(cardAttribute) > 0 {
		// if cardAttribute exists, then this is not a new card and we should
		// bootstrap the Algorithm to use from the cardAttribute
		// else, use the newly created card(up there)
		newCard, err := card.Decode(bytes.NewReader(cardAttribute))
		if err != nil {
			return nil, err
		}
		card = newCard
	}
	return card, nil

}

// Encode encodes the Card value into the encoder.
func (c Card) Encode() error {
	// The Card.Algorithm concrete type(eg Supermemo2) has to be registered
	// using gob.Register else this function will fail
	// We registered it in main.

	var w bytes.Buffer

	enc := gob.NewEncoder(&w)
	err := enc.Encode(&c)
	if err != nil {
		return errors.Wrapf(err, "unable to encode card %v", c)
	}

	err = c.setExtendedAttrs(w.Bytes())
	if err != nil {
		return err
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

// Rate a Card.
func (c *Card) Rate(uInput float64) {
	// the receiver needs to be a pointer so that the changes
	// can propagate back
	sm := c.Algorithm.Advance(uInput)
	c.Algorithm = sm
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

// SetExtendedAttrs sets the files extra metadata
func (c Card) setExtendedAttrs(algoEncoded []byte) error {
	err := xattr.Set(c.FilePath, attrName, algoEncoded)
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
