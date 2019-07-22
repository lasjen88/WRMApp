package characterservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/v1/mongo"

	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/v1/models"
	log "github.com/sirupsen/logrus"
)

//CharacterHandle Rest handle for characters
type CharacterHandle struct {
	CharacterCollection mongo.CharacterCollection
}

const (
	//ContentTypenHeader is the http header key for Content-Type
	ContentTypenHeader string = "Content-Type"
	//ContentTypenHeaderValue is the http header value for Content-Type
	ContentTypenHeaderValue string = "application/json"
	//AccessControlAllowOriginHeader is the http header key for Access-Control-Allow-Origin
	AccessControlAllowOriginHeader string = "Access-Control-Allow-Origin"
	//AccessControlAllowOriginHeaderValue is the http header value for Access-Control-Allow-Origin
	AccessControlAllowOriginHeaderValue string = "*"
)

//GetCharacters fetches all characters
func (c *CharacterHandle) GetCharacters(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(ContentTypenHeader, ContentTypenHeaderValue)
	writer.Header().Set(AccessControlAllowOriginHeader, AccessControlAllowOriginHeaderValue)
	characters, err := c.CharacterCollection.GetAllCharacters()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the charcters"))
		log.Error(err)
		return
	}
	json.NewEncoder(writer).Encode(characters)
}

//CreateCharacter creates a new character
func (c *CharacterHandle) CreateCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(ContentTypenHeader, ContentTypenHeaderValue)
	writer.Header().Set(AccessControlAllowOriginHeader, AccessControlAllowOriginHeaderValue)
	var character models.Character
	parseError := json.NewDecoder(request.Body).Decode(&character)
	if parseError != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("400 - Could not parse json body to character"))
		log.Error(parseError)
		return
	}
	c.CharacterCollection.PutCharacter(character)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(character)
}

//GetCharacter fetches the character specified in the parameter
func (c *CharacterHandle) GetCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(ContentTypenHeader, ContentTypenHeaderValue)
	writer.Header().Set(AccessControlAllowOriginHeader, AccessControlAllowOriginHeaderValue)
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
	writer.Header().Set(ContentTypenHeader, ContentTypenHeaderValue)
	writer.Header().Set(AccessControlAllowOriginHeader, AccessControlAllowOriginHeaderValue)
	params := mux.Vars(request)
	err := c.CharacterCollection.DeleteCharacter(params["id"])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Could not delete character."))
		log.Error(err)
		return
	}
}

//UpdateCharacter updates the character
func (c *CharacterHandle) UpdateCharacter(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(ContentTypenHeader, ContentTypenHeaderValue)
	writer.Header().Set(AccessControlAllowOriginHeader, AccessControlAllowOriginHeaderValue)
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
}
