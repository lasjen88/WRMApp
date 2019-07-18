package main

import (
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/lasjen88/WRMApp/backend/v1/characterservice"
	"github.com/lasjen88/WRMApp/backend/v1/initiativeservice"
	"github.com/lasjen88/WRMApp/backend/v1/itemservice"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"

	"github.com/gorilla/mux"
)

const (
	databaseURL             = "localhost"
	databaseName            = "wrm"
	itemCollectionName      = "equipment"
	spellCollectionName     = "spell"
	characterCollectionName = "equipment"
)

func setupItemServiceRoute(router *mux.Router, mongoSession *mgo.Session) *mux.Router {
	itemCollection := mongo.ItemCollection{
		DatabaseName:   databaseName,
		CollectionName: itemCollectionName,
		Session:        mongoSession,
	}
	err := mongo.InitializeEquipment(mongoSession)
	if err != nil {
		log.Fatal(err)
	}
	spellCollection := mongo.SpellCollection{
		DatabaseName:   databaseName,
		CollectionName: spellCollectionName,
		Session:        mongoSession,
	}
	err = mongo.InitializeSpells(mongoSession)
	if err != nil {
		log.Fatal(err)
	}
	itemHandle := itemservice.ItemHandle{ItemCollection: itemCollection}
	spellHandle := itemservice.SpellHandle{SpellCollection: spellCollection}
	router = itemservice.SetRouteHandles(router, itemHandle, spellHandle)
	return router
}

func setRouteHandles(router *mux.Router, mongoSession *mgo.Session) *mux.Router {
	router = setupItemServiceRoute(router, mongoSession)
	router = characterservice.SetRouteHandles(router)
	router = initiativeservice.SetRouteHandles(router)
	return router
}

func main() {
	//Route handling
	mongoSession := mongo.GetSession(databaseURL)
	router := mux.NewRouter()
	router = setRouteHandles(router, mongoSession)
	defer mongoSession.Close()
	log.Fatal(http.ListenAndServe(":8000", router))
}
