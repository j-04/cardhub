package database

import (
	"context"

	"github.com/j-04/cardhub/types"
)

type DatabaseReader interface {
	GetDecks(context context.Context) ([]*types.Deck, error)
	GetDeck(context context.Context, deckId string) (*types.Deck, error)
	PutWordsInDeck(context context.Context, words []types.Word, deckId string) error
	SaveDeck(context context.Context, deck types.Deck) error
	DeleteDeck(context context.Context, deckId string) error
	DeleteWordInDeck(context context.Context, deckId string, wordId string) error
	GetWords(context context.Context, pageSize int, pageNumber int) ([]types.Word, error)
	SaveWords(context context.Context, words []types.Word) error
	UpdateWord(context context.Context, wordId string, newWord types.Word) error
	Close()
}
