package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	mongo "./v1/mongo"
	"github.com/lasjen88/WRMApp/backend/v1/initiativeservice"
	"github.com/lasjen88/WRMApp/backend/v1/models"
	log "github.com/sirupsen/logrus"

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

const (
	URL     = "localhost"
	DB_NAME = "wrm"
)

func main() {

	router := mux.NewRouter()

	mongoSession := mongo.GetSession(URL)
	DB := mongo.Use(mongoSession, DB_NAME)
	log.Infof("Databases: ")
	mongo.PrintDBNames(mongoSession)
	log.Infof("Collections and Documents: ")
	//mongo.PrintCollectionNames(DB)
	mongo.PrintCollections(DB)
	defer mongoSession.Close()

	// Router Handlers
	router.HandleFunc("/v1/characters", getCharacters).Methods("GET")
	router.HandleFunc("/v1/characters", createCharacter).Methods("POST")
	router.HandleFunc("/v1/characters/{id}", getCharacter).Methods("GET")
	router.HandleFunc("/v1/characters/{id}", updateCharacter).Methods("PUT")
	router.HandleFunc("/v1/characters/{id}", deleteCharacter).Methods("DELETE")

	router.HandleFunc("/v1/initiative", initiativeservice.GetInitiative).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
