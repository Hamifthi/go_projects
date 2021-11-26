package main

import (
	"choose_adventure"
	"choose_adventure/handlers"
	"choose_adventure/jsonReader"
	"flag"
	"log"
	"net/http"
)

func main() {
	filepath := flag.String("file", "./gopher.json", "json file to read stories.")
	flag.Parse()

	storage := make(map[string]adventure.Story)
	reader := jsonReader.Reader{
		TempStorage: storage,
	}
	err := reader.ReadInput(*filepath)
	if err != nil {
		log.Fatalf("Got error in reading %s file", filepath)
	}

	handler := handlers.HTTPHandler{Storage: storage}

	http.HandleFunc("/", handler.GetFirstPage)
	http.HandleFunc("/story", handler.GetStory)
	http.ListenAndServe(":8080", nil)
}
