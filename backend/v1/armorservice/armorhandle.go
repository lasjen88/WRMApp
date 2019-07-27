package armorservice

import (
	"encoding/json"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/router"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//ArmorHandle Rest handle for armor
type ArmorHandle struct {
	ArmorCollection mongo.ArmorCollection
}

//GetArmors fetches all armors
func (a *ArmorHandle) GetArmors(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	armors, err := a.ArmorCollection.GetAllArmor()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the armors"))
		logrus.Error(err)
		return
	}
	logrus.Infof("Found %d armors", len(armors))
	json.NewEncoder(writer).Encode(armors)
}
