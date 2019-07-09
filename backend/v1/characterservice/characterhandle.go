package characterservice

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/v1/models"
)

var characters []models.Character

func setupTestData() {
	characters = append(characters, models.Character{ID: "1", PlayerName: "John"})
	characters = append(characters, models.Character{ID: "2", PlayerName: "Smith"})
}

//GetCharacters fetches all characters
func GetCharacters(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	setupTestData()
	json.NewEncoder(writer).Encode(characters)
}

//CreateCharacter creates a new character
func CreateCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	setupTestData()
	var character models.Character
	_ = json.NewDecoder(request.Body).Decode(&character)
	character.ID = strconv.Itoa(rand.Intn(1000000)) // @TODO FIX
	characters = append(characters, character)
	json.NewEncoder(writer).Encode(character)
}

//GetCharacter fetches the character specified in the parameter
func GetCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	setupTestData()
	params := mux.Vars(request)
	for _, item := range characters {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&models.Character{})
}

//DeleteCharacter deletes the character
func DeleteCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	setupTestData()
	params := mux.Vars(request)
	for index, item := range characters {
		if item.ID == params["id"] {
			characters = append(characters[:index], characters[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(characters)
}

//UpdateCharacter updates the character
func UpdateCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	setupTestData()
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
