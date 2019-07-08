package initiativeservice

import (
	"encoding/json"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/v1/models"
	"github.com/lasjen88/WRMApp/wrmrules"
	log "github.com/sirupsen/logrus"
)

//GetInitiative reads the body of characters provided and returns the characters with a calculated initiative, and sorted in initiative order.
func GetInitiative(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var characters []models.Character
	decodeErr := json.NewDecoder(request.Body).Decode(&characters)
	if decodeErr != nil {
		log.Error(decodeErr)
		return
	}
	log.Debugf("Parsing in %d characters", len(characters))
	initiatives, initiativeErr := wrmrules.GetCharacterInitiatives(characters)
	if initiativeErr != nil {
		log.Error(initiativeErr)
		return
	}
	log.Debugf("Returning %d initiatives", initiatives.Len())
	json.NewEncoder(writer).Encode(initiatives)
}
