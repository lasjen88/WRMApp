package weaponservice

import (
	"encoding/json"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/router"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//WeaponHandle Rest handle for weapons
type WeaponHandle struct {
	WeaponCollection mongo.WeaponCollection
}

//GetWeapons fetches all weapons
func (w *WeaponHandle) GetWeapons(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	weapons, err := w.WeaponCollection.GetAllWeapons()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the weapons"))
		logrus.Error(err)
		return
	}
	logrus.Infof("Found %d weapons", len(weapons))
	json.NewEncoder(writer).Encode(weapons)
}
