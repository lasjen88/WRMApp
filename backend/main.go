package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Character struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var characters []Character

func getCharacters(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(characters)
}

func createCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var character Character
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
	json.NewEncoder(writer).Encode(&Character{})
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
			var character Character
			_ = json.NewDecoder(request.Body).Decode(&character)
			characters = append(characters[:index], characters[index+1:]...)
			characters = append(characters, character)
			break
		}
	}
	json.NewEncoder(writer).Encode(characters)
}

func main() {

	//Init Router
	router := mux.NewRouter()

	// Some data for testing @TODO -- Mongo
	characters = append(characters, Character{ID: "1", Name: "John"})
	characters = append(characters, Character{ID: "2", Name: "Smith"})

	// Router Handlers
	router.HandleFunc("/api/characters", getCharacters).Methods("GET")
	router.HandleFunc("/api/characters", createCharacter).Methods("POST")
	router.HandleFunc("/api/characters/{id}", getCharacter).Methods("GET")
	router.HandleFunc("/api/characters/{id}", updateCharacter).Methods("PUT")
	router.HandleFunc("/api/characters/{id}", deleteCharacter).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
