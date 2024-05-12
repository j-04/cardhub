package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/j-04/cardhub/service"
	"github.com/j-04/cardhub/types"
)

type RequestHandler struct {
	service *service.Service
}

func NewRequestHandler(service *service.Service) *RequestHandler {
	return &RequestHandler{
		service,
	}
}

func (handler *RequestHandler) HandleGreetings(res http.ResponseWriter, req *http.Request) error {
	res.Write([]byte("Greetings!"))
	return nil
}

func (handler *RequestHandler) HandleGetDecks(res http.ResponseWriter, req *http.Request) error {
	context := context.Background()
	decks, err := handler.service.GetDecks(context)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(decks)
	if err != nil {
		return err
	}

	writeResponse(bytes, res)
	return nil
}

func (handler *RequestHandler) HandleGetDeck(res http.ResponseWriter, req *http.Request) error {
	context := context.Background()

	deckId, err := parseIdUrlParam(req, "deckId")
	if err != nil {
		return err
	}

	deck, err := handler.service.GetDeck(context, deckId)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(deck)
	if err != nil {
		return err
	}

	writeResponse(bytes, res)
	return nil
}

func (handler *RequestHandler) HandleSaveDeck(res http.ResponseWriter, req *http.Request) error {
	context := context.Background()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var deck types.Deck

	err = json.Unmarshal(body, &deck)
	if err != nil {
		return err
	}

	err = handler.service.SaveDeck(context, deck)
	if err != nil {
		return err
	}
	return nil
}

func (handler *RequestHandler) HandlePutWordsInDeck(res http.ResponseWriter, req *http.Request) error {
	context := context.Background()

	deckId, err := parseIdUrlParam(req, "deckId")
	if err != nil {
		return err
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	var words []types.Word
	err = json.Unmarshal(body, &words)
	if err != nil {
		return err
	}

	return handler.service.PutWordsInDeck(context, words, deckId)
}

func (handler *RequestHandler) HandleDeleteDeck(res http.ResponseWriter, req *http.Request) error {
	context := context.Background()

	deckId, err := parseIdUrlParam(req, "deckId")
	if err != nil {
		return err
	}

	return handler.service.DeleteDeck(context, deckId)
}

func (handler *RequestHandler) HandleDeleteWordInPeck(res http.ResponseWriter, req *http.Request) error {
	context := context.Background()

	deckId, err := parseIdUrlParam(req, "deckId")
	if err != nil {
		return err
	}
	wordId, err := parseIdUrlParam(req, "wordId")
	if err != nil {
		return err
	}

	return handler.service.DeleteWordInDeck(context, deckId, wordId)
}

func (handler *RequestHandler) HandleGetWords(res http.ResponseWriter, req *http.Request) error {
	context := context.Background()

	page, size, err := validateGetWordsRequest(req)
	if err != nil {
		return err
	}

	words, err := handler.service.GetWords(context, page, size)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(words)
	if err != nil {
		return types.ValidationErr{Msg: err.Error()}
	}

	writeResponse(bytes, res)
	return nil
}

func (handler *RequestHandler) HandleSaveWord(res http.ResponseWriter, req *http.Request) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	words := make([]types.Word, 5)

	err = json.Unmarshal(body, &words)
	if err != nil {
		return err
	}

	context := context.Background()

	err = handler.service.SaveWords(context, words)
	if err != nil {
		return err
	}

	return nil
}

func (handler *RequestHandler) HandlerUpdateWord(res http.ResponseWriter, req *http.Request) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	wordId, err := parseIdUrlParam(req, "wordId")
	if err != nil {
		return err
	}

	word := &types.Word{}

	err = json.Unmarshal(body, word)
	if err != nil {
		return err
	}

	context := context.Background()

	err = handler.service.UpdateWord(context, wordId, *word)
	if err != nil {
		return err
	}

	return nil
}

func validateGetWordsRequest(req *http.Request) (page int, size int, err error) {
	page, errPage := strconv.Atoi(req.URL.Query().Get("page"))
	size, errSize := strconv.Atoi(req.URL.Query().Get("size"))
	if errPage != nil || errSize != nil {
		log.Printf("Validation failed. Page: %v, size: %v", errPage, errSize)

		err = types.ValidationErr{
			Msg: "page or size param are invalid. Params should be integer value and not empty.",
		}

		return
	}
	return
}

func parseIdUrlParam(req *http.Request, param string) (string, error) {
	id := chi.URLParam(req, param)
	if id == "" {
		return "", types.ValidationErr{
			Msg: fmt.Sprintf("%s attribute is empty or not a number.", param),
		}
	}
	return id, nil
}

func writeResponse(data []byte, res http.ResponseWriter) {
	res.Header().Add("Content-Type", "application/json")
	res.Write(data)
}
