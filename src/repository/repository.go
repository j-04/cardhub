package repository

import (
	"github.com/j-04/cardhub/database"
	"github.com/j-04/cardhub/model"
)

type Repository struct {
	database database.DatabaseReader
}

func NewRepository() *Repository {
	return &Repository{
		database: database.NewStubDatabse(),
	}
}

func (repo *Repository) GetWords(pageSize int, pageNumber int) []model.Word {
	return repo.database.GetWords(pageSize, pageNumber)
}
