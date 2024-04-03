package main

import (
	"context"
	"encoding/json"
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

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{
		service: service.NewService(),
	}
}

func (handler *RequestHandler) HandleGreetings(res http.ResponseWriter, req *http.Request) error {
	res.Write([]byte("Greetings!"))
	return nil
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

	wordId, err := strconv.ParseInt(chi.URLParam(req, "wordId"), 0, 64)
	if err != nil {
		return types.ValidationErr{
			Msg: "wordId attribute is empty or not a number.",
		}
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

func writeResponse(data []byte, res http.ResponseWriter) {
	res.Header().Add("Content-Type", "application/json")
	res.Write(data)
}
