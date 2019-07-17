package main

import (
	"net/http"

	mongogo "./v1/mongo"

	"github.com/lasjen88/WRMApp/backend/v1/characterservice"
	"github.com/lasjen88/WRMApp/backend/v1/initiativeservice"
	"github.com/lasjen88/WRMApp/backend/v1/itemservice"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

const (
	URL     = "localhost"
	DB_NAME = "wrm"
)

func setRouteHandles(router *mux.Router) *mux.Router {
	router = characterservice.SetRouteHandles(router)
	router = itemservice.SetRouteHandles(router)
	router = initiativeservice.SetRouteHandles(router)
	return router
}

func main() {
	//Database handling
	mongoSession := mongogo.GetSession(URL)
	//DB := mongogo.Use(mongoSession, DB_NAME)
	log.Infof("Databases: ")
	mongogo.PrintDBNames(mongoSession)
	log.Infof("Collections and Documents: ")
	//mongo.PrintCollectionNames(DB)
	//mongogo.PrintCollections(DB)
	err := mongogo.InitializeEquipment(mongoSession)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoSession.Close()

	//Route handling
	router := mux.NewRouter()
	router = setRouteHandles(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
