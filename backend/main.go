package main

import (
	"net/http"

	mongo "./v1/mongo"
	"github.com/lasjen88/WRMApp/backend/v1/characterservice"
	"github.com/lasjen88/WRMApp/backend/v1/itemservice"
	"github.com/lasjen88/WRMApp/backend/v1/initiativeservice"
	"github.com/lasjen88/WRMApp/backend/v1/models"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

const (
	URL     = "localhost"
	DB_NAME = "wrm"
)

func setRouteHandles(route *Router) *Router {
	router.HandleFunc("/v1/characters", characterservice.GetCharacters).Methods("GET")
	router.HandleFunc("/v1/characters", characterservice.CreateCharacter).Methods("POST")
	router.HandleFunc("/v1/characters/{id}", characterservice.GetCharacter).Methods("GET")
	router.HandleFunc("/v1/characters/{id}", characterservice.UpdateCharacter).Methods("PUT")
	router.HandleFunc("/v1/characters/{id}", characterservice.DeleteCharacter).Methods("DELETE")
	router.HandleFunc("/v1/items", itemservice.GetItems).Methods("GET")
	router.HandleFunc("/v1/items", itemservice.CreateItem).Methods("POST")
}

func main() {
	router := mux.NewRouter()
	mongoSession := mongo.GetSession(URL)
	DB := mongo.Use(mongoSession, DB_NAME)
	log.Infof("Databases: ")
	mongo.PrintDBNames(mongoSession)
	log.Infof("Collections and Documents: ")
	//mongo.PrintCollectionNames(DB)
	mongo.PrintCollections(DB)
	defer mongoSession.Close()

	// Router Handlers
	router = setRouteHandles()

	router.HandleFunc("/v1/initiative", initiativeservice.GetInitiative).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
