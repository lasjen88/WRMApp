package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"./v1/models"

	"github.com/gorilla/mux"
)

var characters []models.Character

func getCharacters(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(characters)
}

func createCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var character models.Character
	_ = json.NewDecoder(request.Body).Decode(&character)
	character.ID = strconv.Itoa(rand.Intn(1000000)) // @TODO FIX
	characters = append(characters, character)
	json.NewEncoder(writer).Encode(character)
}

func getCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range characters {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&models.Character{})
}

func deleteCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range characters {
		if item.ID == params["id"] {
			characters = append(characters[:index], characters[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(characters)
}

func updateCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range characters {
		if item.ID == params["id"] {
			var character models.Character
			_ = json.NewDecoder(request.Body).Decode(&character)
			characters = append(characters[:index], characters[index+1:]...)
			characters = append(characters, character)
			break
		}
	}
	json.NewEncoder(writer).Encode(characters)
}

func main() {

	router := mux.NewRouter()

	// Some data for testing @TODO -- Mongo
	characters = append(characters, models.Character{ID: "1", Name: "John"})
	characters = append(characters, models.Character{ID: "2", Name: "Smith"})

	// Router Handlers
	router.HandleFunc("/v1/characters", getCharacters).Methods("GET")
	router.HandleFunc("/v1/characters", createCharacter).Methods("POST")
	router.HandleFunc("/v1/characters/{id}", getCharacter).Methods("GET")
	router.HandleFunc("/v1/characters/{id}", updateCharacter).Methods("PUT")
	router.HandleFunc("/v1/characters/{id}", deleteCharacter).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
