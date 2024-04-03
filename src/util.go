package main

import (
	"encoding/json"
	"log"
)

func writeJson(data interface{}) (bytes []byte) {
	log.Println("Marshalling the data ", data)
	bytes, _ = json.Marshal(data)
	return bytes
}
