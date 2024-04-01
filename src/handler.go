package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/j-04/cardhub/service"
)

type RequestHandler struct {
	service *service.Service
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{
		service: service.NewService(),
	}
}

func (handler *RequestHandler) HandleGreetings(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Greetings!"))
}

func (handler *RequestHandler) HandleGetWords(res http.ResponseWriter, req *http.Request) {
	page, errPage := strconv.Atoi(req.URL.Query().Get("page"))
	size, errSize := strconv.Atoi(req.URL.Query().Get("size"))

	if errPage != nil || errSize != nil {
		log.Println(errPage)
		log.Println(errSize)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("page or size query param are invalid. Params should be integer value and not empty"))
		return
	}

	bytes, err := json.Marshal(handler.service.GetWords(page, size))

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	res.Write(bytes)
}

func (handler *RequestHandler) HandleSaveWord(res http.ResponseWriter, req *http.Request) {

}
