package armorservice

import (
	"encoding/json"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/router"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//ShieldHandle Rest handle for armor
type ShieldHandle struct {
	ShieldCollection mongo.ShieldCollection
}

//GetShields fetches all shields
func (s *ShieldHandle) GetShields(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	shields, err := s.ShieldCollection.GetAllShields()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the shields"))
		logrus.Error(err)
		return
	}
	logrus.Infof("Found %d shields", len(shields))
	json.NewEncoder(writer).Encode(shields)
}
