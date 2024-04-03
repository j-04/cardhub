package service

import (
	"context"

	"github.com/j-04/cardhub/repository"
	"github.com/j-04/cardhub/types"
)

type Service struct {
	repository *repository.Repository
}

func NewService() *Service {
	return &Service{
		repository: repository.NewRepository(),
	}
}

func (service *Service) GetWords(context context.Context, pageSize int, pageNumber int) ([]types.Word, error) {
	return service.repository.GetWords(context, pageSize, pageNumber)
}

func (service *Service) SaveWords(context context.Context, words []types.Word) error {
	return service.repository.SaveWords(context, words)
}

func (service *Service) UpdateWord(context context.Context, wordId int64, newWord types.Word) error {
	return service.repository.UpdateWord(context, wordId, newWord)
}
