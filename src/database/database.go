package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/j-04/cardhub/types"
)

type DatabaseReader interface {
	GetWords(context context.Context, pageSize int, pageNumber int) ([]types.Word, error)
	SaveWords(context context.Context, words []types.Word) error
	UpdateWord(context context.Context, wordId int64, newWord types.Word) error
}

type StubDatabase struct {
	mu      sync.Mutex
	counter int64
	data    map[int64]types.Word
}

func NewStubDatabse() *StubDatabase {
	m := make(map[int64]types.Word)
	m[0] = types.Word{
		Id:    0,
		Front: "ždravo",
		Back:  "hello",
	}

	return &StubDatabase{
		mu:      sync.Mutex{},
		counter: 1,
		data:    m,
	}
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
