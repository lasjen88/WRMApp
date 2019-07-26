package talentservice

import (
	"encoding/json"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/router"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//TalentHandle Rest handle for talents
type TalentHandle struct {
	TalentCollection mongo.TalentCollection
}

//GetTalents fetches all talents
func (t *TalentHandle) GetTalents(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	talents, err := t.TalentCollection.GetAllTalents()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the talents"))
		logrus.Error(err)
		return
	}
	logrus.Infof("Found %d talents", len(talents))
	json.NewEncoder(writer).Encode(talents)
}
