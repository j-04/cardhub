package repository

import (
	"context"

	"github.com/j-04/cardhub/database"
	"github.com/j-04/cardhub/types"
)

type Repository struct {
	database database.DatabaseReader
}

func NewRepository() *Repository {
	return &Repository{
		database: database.NewStubDatabse(),
	}
}

func (repo *Repository) GetDecks(context context.Context) ([]*types.Deck, error) {
	decks, err := repo.database.GetDecks(context)
	if err != nil {
		return nil, err
	}
	for _, deck := range decks {
		deck.Size = len(deck.Words)
	}
	return decks, nil
}

func (repo *Repository) GetDeck(context context.Context, deckId int64) (*types.Deck, error) {
	return repo.database.GetDeck(context, deckId)
}

func (repo *Repository) PutWordsInDeck(context context.Context, words []types.Word, deckId int64) error {
	return repo.database.PutWordsInDeck(context, words, deckId)
}

func (repo *Repository) SaveDeck(context context.Context, deck types.Deck) error {
	return repo.database.SaveDeck(context, deck)
}

func (repo *Repository) DeleteDeck(context context.Context, deckId int64) error {
	return repo.database.DeleteDeck(context, deckId)
}

func (repo *Repository) DeleteWordinDeck(context context.Context, deckId int64, wordId int64) error {
	return repo.database.DeleteWordInDeck(context, deckId, wordId)
}

func (repo *Repository) GetWords(context context.Context, pageSize int, pageNumber int) ([]types.Word, error) {
	return repo.database.GetWords(context, pageSize, pageNumber)
}

func (repo *Repository) SaveWords(context context.Context, words []types.Word) error {
	return repo.database.SaveWords(context, words)
}

func (repo *Repository) UpdateWord(context context.Context, wordId int64, newWord types.Word) error {
	return repo.database.UpdateWord(context, wordId, newWord)
}
