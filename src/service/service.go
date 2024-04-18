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

func (service *Service) GetDecks(context context.Context) ([]*types.Deck, error) {
	return service.repository.GetDecks(context)
}

func (service *Service) GetDeck(context context.Context, deckId int64) (*types.Deck, error) {
	return service.repository.GetDeck(context, deckId)
}

func (service *Service) SaveDeck(context context.Context, deck types.Deck) error {
	return service.repository.SaveDeck(context, deck)
}

func (service *Service) PutWordsInDeck(context context.Context, words []types.Word, deckId int64) error {
	return service.repository.PutWordsInDeck(context, words, deckId)
}

func (service *Service) DeleteDeck(context context.Context, deckId int64) error {
	return service.repository.DeleteDeck(context, deckId)
}

func (service *Service) DeleteWordInDeck(context context.Context, deckId int64, wordId int64) error {
	return service.repository.DeleteWordinDeck(context, deckId, wordId)
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
