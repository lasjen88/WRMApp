package raceservice

import (
	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//SetRouteHandles sets the talent handles on the router
func SetRouteHandles(router *mux.Router) *mux.Router {
	collection := mongo.WrmCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.RaceCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	raceCollection := mongo.RaceCollection{Wrmcollection: collection}
	loadInitialRaces(raceCollection)
	racehandle := RaceHandle{RaceCollection: raceCollection}
	router.HandleFunc("/v1/races", racehandle.GetRaces).Methods("GET")
	return router
}

func loadInitialRaces(weaponCollection mongo.RaceCollection) {
	collection := mongo.WrmCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.TalentCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	talentCollection := mongo.TalentCollection{Wrmcollection: collection}
	err := mongo.InitializRaces(weaponCollection, talentCollection)
	if err != nil {
		logrus.Fatal(err)
	}
}
