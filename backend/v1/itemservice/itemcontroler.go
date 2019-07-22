package itemservice

import (
	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//SetRouteHandles sets the item and spell handles on the router
func SetRouteHandles(router *mux.Router) *mux.Router {
	itemCollection := mongo.ItemCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.ItemCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	loadInitialItems(itemCollection)
	spellCollection := mongo.SpellCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.SpellCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	loadInitialSpells(spellCollection)
	itemHandle := ItemHandle{ItemCollection: itemCollection}
	spellHandle := SpellHandle{SpellCollection: spellCollection}
	return defineAPI(router, itemHandle, spellHandle)
}

func defineAPI(router *mux.Router, itemHandle ItemHandle, spellHandle SpellHandle) *mux.Router {
	router.HandleFunc("/v1/items", itemHandle.GetItems).Methods("GET")
	router.HandleFunc("/v1/items", itemHandle.CreateItem).Methods("POST")
	router.HandleFunc("/v1/spells", spellHandle.GetSpells).Methods("GET")
	router.HandleFunc("/v1/spells", spellHandle.CreateSpell).Methods("POST")
	router.HandleFunc("/v1/spells/{circle}", spellHandle.GetCircleSpells).Methods("GET")
	return router
}

func loadInitialSpells(spellCollection mongo.SpellCollection) {
	err := mongo.InitializeSpells(spellCollection.Session)
	if err != nil {
		logrus.Fatal(err)
	}
}

func loadInitialItems(itemCollection mongo.ItemCollection) {
	err := mongo.InitializeEquipment(itemCollection.Session)
	if err != nil {
		logrus.Fatal(err)
	}
}
