package database

import (
	"github.com/j-04/cardhub/model"
)

type DatabaseReader interface {
	GetWords(pageSize int, pageNumber int) []model.Word
}

type StubDatabase struct {
}

func NewStubDatabse() *StubDatabase {
	return &StubDatabase{}
}

func (db *StubDatabase) GetWords(pageSize int, pageNumber int) []model.Word {
	return []model.Word{
		{
			Front: "ždravo",
			Back:  "hello",
		},
	}
}
