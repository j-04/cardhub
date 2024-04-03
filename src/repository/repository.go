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

func (repo *Repository) GetWords(context context.Context, pageSize int, pageNumber int) ([]types.Word, error) {
	return repo.database.GetWords(context, pageSize, pageNumber)
}

func (repo *Repository) SaveWords(context context.Context, words []types.Word) error {
	return repo.database.SaveWords(context, words)
}

func (repo *Repository) UpdateWord(context context.Context, wordId int64, newWord types.Word) error {
	return repo.database.UpdateWord(context, wordId, newWord)
}
