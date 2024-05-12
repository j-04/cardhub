package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/j-04/cardhub/types"
)

type StubDatabase struct {
	mu    sync.Mutex
	data  map[string]types.Word
	decks map[string]*types.Deck
	_     struct{}
}

func NewStubDatabse() *StubDatabase {
	m := make(map[string]types.Word)
	uuid := uuid.New().String()
	m[uuid] = types.Word{
		Id:    uuid,
		Front: "ždravo",
		Back:  "hello",
	}

	bytes, err := os.ReadFile("resources/test-data.json")
	if err != nil {
		log.Fatalln(err.Error())
	}

	var decks map[string]*types.Deck

	err = json.Unmarshal(bytes, &decks)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return &StubDatabase{
		mu:    sync.Mutex{},
		data:  m,
		decks: decks,
	}
}

func (db *StubDatabase) Close() {

}

func (db *StubDatabase) GetDecks(context context.Context) ([]*types.Deck, error) {
	decks := make([]*types.Deck, 0, len(db.decks))
	for _, v := range db.decks {
		decks = append(decks, v)
	}

	return decks, nil
}

func (db *StubDatabase) GetDeck(context context.Context, deckId string) (*types.Deck, error) {
	deck, ok := db.decks[deckId]
	if !ok {
		return nil, types.NotFoundErr{
			Msg: fmt.Sprintf("could not find deck by id %s", deckId),
		}
	}

	return deck, nil
}

func (db *StubDatabase) PutWordsInDeck(context context.Context, words []types.Word, deckId string) error {
	deck, ok := db.decks[deckId]
	if !ok {
		return types.NotFoundErr{
			Msg: fmt.Sprintf("could not find deck by id %s", deckId),
		}
	}

	uuid := uuid.New().String()

	for _, word := range words {
		word.Id = uuid
	}

	log.Println(words)

	deck.Words = append(deck.Words, words...)
	return nil
}

func (db *StubDatabase) SaveDeck(context context.Context, deck types.Deck) error {
	uuid := uuid.New().String()
	(&deck).Id = uuid
	db.decks[uuid] = &deck
	return nil
}

func (db *StubDatabase) DeleteDeck(context context.Context, deckId string) error {
	delete(db.decks, deckId)
	return nil
}

func (db *StubDatabase) DeleteWordInDeck(context context.Context, deckId string, wordId string) error {
	deck, ok := db.decks[deckId]
	if !ok {
		return types.NotFoundErr{
			Msg: fmt.Sprintf("could not find deck by id %s", deckId),
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

	uuid := uuid.New().String()
	for _, word := range words {
		db.data[uuid] = word
	}
	return nil
}

func (db *StubDatabase) UpdateWord(context context.Context, wordId string, newWord types.Word) error {
	value, ok := db.data[wordId]
	if !ok {
		return types.NotFoundErr{
			Msg: fmt.Sprintf("could not find word by id %s", wordId),
		}
	}
	value.Front = newWord.Front
	value.Back = newWord.Back
	db.data[wordId] = value
	return nil
}
