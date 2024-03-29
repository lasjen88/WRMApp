package characterservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/router"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"

	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/v1/models"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

//CharacterHandle Rest handle for characters
type CharacterHandle struct {
	CharacterCollection mongo.CharacterCollection
}

//GetCharacters fetches all characters
func (c *CharacterHandle) GetCharacters(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	characters, err := c.CharacterCollection.GetAllCharacters()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the charcters"))
		log.Error(err)
		return
	}
	logrus.Infof("Found %d charcters", len(characters))
	json.NewEncoder(writer).Encode(characters)
}

//CreateCharacter creates a new character
func (c *CharacterHandle) CreateCharacter(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	var character models.Character
	parseError := json.NewDecoder(request.Body).Decode(&character)
	if parseError != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("400 - Could not parse json body to character"))
		log.Error(parseError)
		return
	}
	err := c.CharacterCollection.PutCharacter(character)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Could not character save charcter to database"))
		log.Error(err)
		return
	}
	logrus.Infof("Created character: %s", character.CharacterName)
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Character created"))
}

//GetCharacter fetches the character specified in the parameter
func (c *CharacterHandle) GetCharacter(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	params := mux.Vars(request)
	characters, err := c.CharacterCollection.GetAllCharacters()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the characters."))
		log.Error(err)
		return
	}
	for _, character := range characters {
		if character.ID == params["id"] {
			json.NewEncoder(writer).Encode(character)
			return
		}
	}
	writer.WriteHeader(http.StatusNotFound)
	notFoundMessage := fmt.Sprintf("404 - Could not find character with id [%s].", params["id"])
	writer.Write([]byte(notFoundMessage))
}

//DeleteCharacter deletes the character
func (c *CharacterHandle) DeleteCharacter(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	params := mux.Vars(request)
	err := c.CharacterCollection.DeleteCharacter(params["id"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Could not delete character."))
		log.Error(err)
		return
	}
	writer.Write([]byte("Character deleted"))
}

//UpdateCharacter updates the character
func (c *CharacterHandle) UpdateCharacter(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	var character models.Character
	parseError := json.NewDecoder(request.Body).Decode(&character)
	if parseError != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("400 - Could not parse json body to character"))
		log.Error(parseError)
		return
	}
	params := mux.Vars(request)
	err := c.CharacterCollection.UpdateCharacter(params["id"], character)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Could not delete character."))
		log.Error(err)
		return
	}
	writer.Write([]byte("Character updated"))
}
