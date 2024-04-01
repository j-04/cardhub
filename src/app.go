package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	root.Get("/", handler.HandleGreetings)

	api := chi.NewRouter()
	api.Get("/words", handler.HandleGetWords)
	api.Post("/words", handler.HandleSaveWord)

	root.Mount("/api/v1", api)

	log.Println("server is up and running")
	err := http.ListenAndServe(config.GetHostAndPort(), root)
	if err != nil {
		log.Fatalln(err.Error())
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

	bytes, err := os.ReadFile(fileName)
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
