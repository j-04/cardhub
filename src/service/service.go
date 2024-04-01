package service

import (
	"github.com/j-04/cardhub/model"
	"github.com/j-04/cardhub/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService() *Service {
	return &Service{
		repository: repository.NewRepository(),
	}
}

func (service *Service) GetWords(pageSize int, pageNumber int) []model.Word {
	return service.repository.GetWords(pageSize, pageNumber)
}
