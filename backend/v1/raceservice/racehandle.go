package raceservice

import (
	"encoding/json"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/router"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//RaceHandle Rest handle for races
type RaceHandle struct {
	RaceCollection mongo.RaceCollection
}

//GetRaces fetches all races
func (r *RaceHandle) GetRaces(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	races, err := r.RaceCollection.GetAllRaces()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the races"))
		logrus.Error(err)
		return
	}
	logrus.Infof("Found %d races", len(races))
	json.NewEncoder(writer).Encode(races)
}
