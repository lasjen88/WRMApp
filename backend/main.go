package main

import (
	"net/http"

	mongo "./v1/mongo"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

const (
	URL     = "localhost"
	DB_NAME = "wrm"
)

func setRouteHandles(router *mux.Router) *mux.Router {
	router = characterservice.setRouteHandles(router)
	router = itemservice.setRouteHandles(router)
	router = initiativeservice.setRouteHandles(router)
	return router
}

func main() {
	//Database handling
	mongoSession := mongo.GetSession(URL)
	DB := mongo.Use(mongoSession, DB_NAME)
	log.Infof("Databases: ")
	mongo.PrintDBNames(mongoSession)
	log.Infof("Collections and Documents: ")
	//mongo.PrintCollectionNames(DB)
	mongo.PrintCollections(DB)
	defer mongoSession.Close()

	//Route handling
	router := mux.NewRouter()
	router = setRouteHandles(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
