package srs

import (
	"bytes"
	"encoding/gob"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/pkg/errors"

	"github.com/alecthomas/chroma/quick"

	"go.etcd.io/bbolt"
)

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
	FileName  string
	CardDir   string
	Algorithm SRSalgorithm
}

// NewCard returns a new Card
func NewCard(filename string, cardDir string, db *bbolt.DB) (*Card, error) {
	filepath := filepath.Join(cardDir, filename)
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

	cardAttribute, err := getCard(db, question)
	if err != nil {
		return nil, err
	}

	card := &Card{
		Version:  1,
		Question: question,
		FileName: filename,

		Algorithm: NewEbisu(), // NewSupermemo2(),
	}
	if len(cardAttribute) > 0 {
		// if cardAttribute exists, then this is not a new card and we should
		// bootstrap the Algorithm to use from the cardAttribute
		// else, use the newly created card(up there)
		newCard, err := card.Decode(bytes.NewReader(cardAttribute))
		card = newCard
		if err != nil {
			return nil, err
		}
	}
	card.CardDir = cardDir

	return card, nil
}

// Path returns the absolute path to a card on disk.
func (c Card) Path() string {
	return filepath.Join(c.CardDir, c.FileName)
}

// Encode encodes the Card value into the encoder.
func (c Card) Encode(db *bbolt.DB) error {
	// The Card.Algorithm concrete type(eg Supermemo2) has to be registered
	// using gob.Register else this function will fail
	// We registered it in main.

	var w bytes.Buffer

	enc := gob.NewEncoder(&w)
	err := enc.Encode(&c)
	if err != nil {
		return errors.Wrapf(err, "unable to encode card %v", c)
	}

	err = saveCard(db, c.Question, w.Bytes())
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
	filepath := c.Path()
	source, err := ioutil.ReadFile(filepath)
	if err != nil {
		return errors.Wrapf(err, "unable to read file %v", filepath)

	}
	err = quick.Highlight(os.Stdout, string(source), "markdown", "terminal16m", "pygments") // some other good styles: paraiso-dark, native, fruity, rrt
	if err != nil {
		return errors.Wrapf(err, "unable to render %s on screen", filepath)
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

func OpenDb(dbPath string) (*bbolt.DB, error) {
	// caller should close the database
	db, err := bbolt.Open(dbPath, 0600, &bbolt.Options{Timeout: 3 * time.Second})

	if err != nil {
		return nil, errors.Wrapf(err, "unable to open DB at path %v", dbPath)
	}
	return db, nil
}

func saveCard(db *bbolt.DB, card string, data []byte) error {
	var err error
	db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(card))
		if err != nil {
			return errors.Wrapf(err, "unable to create bucket: %v", card)
		}

		return b.Put([]byte(card), data)
	})
	return err
}

func getCard(db *bbolt.DB, card string) ([]byte, error) {
	var err error
	var data []byte
	db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(card))
		if err != nil {
			return errors.Wrapf(err, "unable to create bucket: %v", card)
		}
		data = b.Get([]byte(card))
		return nil
	})
	return data, err
}
