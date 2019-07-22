package main

import (
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/lasjen88/WRMApp/backend/v1/characterservice"
	"github.com/lasjen88/WRMApp/backend/v1/initiativeservice"
	"github.com/lasjen88/WRMApp/backend/v1/itemservice"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

const (
	databaseURL             = "localhost"
	databaseName            = "wrm"
	itemCollectionName      = "equipment"
	spellCollectionName     = "spell"
	characterCollectionName = "character"
	backendPort             = "8000"
)

func setupItemServiceRoute(router *mux.Router, mongoSession *mgo.Session) *mux.Router {
	itemCollection := mongo.ItemCollection{
		DatabaseName:   databaseName,
		CollectionName: itemCollectionName,
		Session:        mongoSession,
	}
	err := mongo.InitializeEquipment(mongoSession)
	if err != nil {
		logrus.Fatal(err)
	}
	spellCollection := mongo.SpellCollection{
		DatabaseName:   databaseName,
		CollectionName: spellCollectionName,
		Session:        mongoSession,
	}
	err = mongo.InitializeSpells(mongoSession)
	if err != nil {
		logrus.Fatal(err)
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
	mongoSession := mongo.GetSession(mongo.DatabaseURL)
	router := mux.NewRouter()
	router = setRouteHandles(router, mongoSession)
	defer mongoSession.Close()
	logrus.Fatal(http.ListenAndServe(":"+backendPort, router))
}
