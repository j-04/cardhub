package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/j-04/cardhub/types"
)

type DatabaseReader interface {
	GetDecks(context context.Context) ([]*types.Deck, error)
	GetDeck(context context.Context, deckId int64) (*types.Deck, error)
	PutWordsInDeck(context context.Context, words []types.Word, deckId int64) error
	SaveDeck(context context.Context, deck types.Deck) error
	DeleteDeck(context context.Context, deckId int64) error
	DeleteWordInDeck(context context.Context, deckId int64, wordId int64) error
	GetWords(context context.Context, pageSize int, pageNumber int) ([]types.Word, error)
	SaveWords(context context.Context, words []types.Word) error
	UpdateWord(context context.Context, wordId int64, newWord types.Word) error
}

type StubDatabase struct {
	mu          sync.Mutex
	counter     int64
	deckCounter int64
	data        map[int64]types.Word
	decks       map[int64]*types.Deck
	_           struct{}
}

func NewStubDatabse() *StubDatabase {
	m := make(map[int64]types.Word)
	m[0] = types.Word{
		Id:    0,
		Front: "ždravo",
		Back:  "hello",
	}

	bytes, err := os.ReadFile("resources/test-data.json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	var decks map[int64]*types.Deck

	err = json.Unmarshal(bytes, &decks)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &StubDatabase{
		mu:          sync.Mutex{},
		counter:     2,
		deckCounter: 1,
		data:        m,
		decks:       decks,
	}
}

func (db *StubDatabase) GetDecks(context context.Context) ([]*types.Deck, error) {
	decks := make([]*types.Deck, 0, len(db.decks))
	for _, v := range db.decks {
		decks = append(decks, v)
	}

	return decks, nil
}

func (db *StubDatabase) GetDeck(context context.Context, deckId int64) (*types.Deck, error) {
	deck, ok := db.decks[deckId]
	if !ok {
		return nil, types.NotFoundErr{
			Msg: fmt.Sprintf("could not find deck by id %d", deckId),
		}
	}

	return deck, nil
}

func (db *StubDatabase) PutWordsInDeck(context context.Context, words []types.Word, deckId int64) error {
	deck, ok := db.decks[deckId]
	if !ok {
		return types.NotFoundErr{
			Msg: fmt.Sprintf("could not find deck by id %d", deckId),
		}
	}

	for _, word := range words {
		word.Id = db.counter
		db.counter++
	}

	log.Println(words)

	deck.Words = append(deck.Words, words...)
	return nil
}

func (db *StubDatabase) SaveDeck(context context.Context, deck types.Deck) error {
	(&deck).Id = db.deckCounter
	db.decks[db.deckCounter] = &deck
	return nil
}

func (db *StubDatabase) DeleteDeck(context context.Context, deckId int64) error {
	delete(db.decks, deckId)
	return nil
}

func (db *StubDatabase) DeleteWordInDeck(context context.Context, deckId int64, wordId int64) error {
	deck, ok := db.decks[deckId]
	if !ok {
		return types.NotFoundErr{
			Msg: fmt.Sprintf("could not find deck by id %d", deckId),
		}
	}
	var wordIndex int

	words := deck.Words
	for i, word := range words {
		if word.Id == wordId {
			wordIndex = i
			break
		}
	}
	deck.Words[wordIndex] = deck.Words[len(deck.Words)-1]
	deck.Words = deck.Words[:len(deck.Words)-1]
	return nil
}

func (db *StubDatabase) GetWords(context context.Context, pageSize int, pageNumber int) ([]types.Word, error) {
	words := make([]types.Word, 0, len(db.data))
	for _, v := range db.data {
		words = append(words, v)
	}

	return words, nil
}

func (db *StubDatabase) SaveWords(context context.Context, words []types.Word) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	for _, word := range words {
		word.Id = db.counter
		db.data[db.counter] = word
		db.counter++
	}
	return nil
}

func (db *StubDatabase) UpdateWord(context context.Context, wordId int64, newWord types.Word) error {
	value, ok := db.data[wordId]
	if !ok {
		return types.NotFoundErr{
			Msg: fmt.Sprintf("could not find word by id %d", wordId),
		}
	}
	value.Front = newWord.Front
	value.Back = newWord.Back
	db.data[wordId] = value
	return nil
}
