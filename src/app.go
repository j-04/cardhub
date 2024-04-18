package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/j-04/cardhub/types"
	"gopkg.in/yaml.v3"
)

func main() {
	config := loadConfig()

	handler := NewRequestHandler()

	root := chi.NewRouter()

	log.Println("adding middleware...")

	root.Use(middleware.Logger)
	root.Use(middleware.Heartbeat("/ping"))

	log.Println("creating routes...")
	root.Get("/", handleError(handler.HandleGreetings))

	api := chi.NewRouter()

	api.Get("/decks", handleError(handler.HandleGetDecks))
	api.Get("/decks/{deckId}", handleError(handler.HandleGetDeck))
	api.Post("/decks", handleError(handler.HandleSaveDeck))
	api.Put("/decks/{deckId}", handleError(handler.HandlePutWordsInDeck))
	api.Delete("/decks/{deckId}", handleError(handler.HandleDeleteDeck))
	api.Delete("/decks/{deckId}/words/{wordId}", handleError(handler.HandleDeleteWordInPeck))

	api.Get("/words", handleError(handler.HandleGetWords))
	api.Post("/words", handleError(handler.HandleSaveWord))
	api.Put("/words/{wordId}", handleError(handler.HandlerUpdateWord))

	root.Mount("/api/v1", api)

	log.Println("server is up and running")
	err := http.ListenAndServe(config.GetHostAndPort(), root)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type HandlerWithErrorFunc func(res http.ResponseWriter, req *http.Request) error

func handleError(handler HandlerWithErrorFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			if err, ok := err.(types.ValidationErr); ok {
				log.Println("Validation was failed with msg:", err.Msg)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(writeJson(map[string]string{"error": err.Msg}))
				return
			}

			if err, ok := err.(types.NotFoundErr); ok {
				log.Println("Something was not found:", err.Msg)
				w.WriteHeader(http.StatusNotFound)
				w.Write(writeJson(map[string]string{"error": err.Msg}))
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(writeJson(map[string]string{"error": err.Error()}))
			return
		}
	}
}

func loadConfig() *Config {
	profile := os.Getenv("APP_PROFILE")

	var fileName string

	switch profile {
	default:
		profile = "local"
		fileName = "local.yaml"
	}

	log.Printf("application profile %s detected\n", profile)
	log.Printf("loading file %s", fileName)

	bytes, err := os.ReadFile("resources/" + fileName)
	if err != nil {
		log.Fatalf("Couldn't read file %s. Error: %s", fileName, err.Error())
	}

	log.Println("config file loaded succesfully")

	config := &Config{}

	yaml.Unmarshal(bytes, config)

	bytes, _ = json.Marshal(config)
	log.Println(string(bytes))
	return config
}
