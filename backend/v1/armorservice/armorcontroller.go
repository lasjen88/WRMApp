package armorservice

import (
	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//SetRouteHandles sets the talent handles on the router
func SetRouteHandles(router *mux.Router) *mux.Router {
	armorhandle := getArmorHandle()
	shieldhandle := getShieldHandle()
	router.HandleFunc("/v1/armors", armorhandle.GetArmors).Methods("GET")
	router.HandleFunc("/v1/shields", shieldhandle.GetShields).Methods("GET")
	return router
}

func getArmorHandle() ArmorHandle {
	collection := mongo.WrmCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.ArmorCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	armorCollection := mongo.ArmorCollection{Wrmcollection: collection}
	loadInitialArmors(armorCollection)
	armorhandle := ArmorHandle{ArmorCollection: armorCollection}
	return armorhandle
}

func getShieldHandle() ShieldHandle {
	collection := mongo.WrmCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.ShieldCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	shieldCollection := mongo.ShieldCollection{Wrmcollection: collection}
	loadInitialShields(shieldCollection)
	shieldhandle := ShieldHandle{ShieldCollection: shieldCollection}
	return shieldhandle
}

func loadInitialArmors(armorCollection mongo.ArmorCollection) {
	err := mongo.InitializeArmors(armorCollection)
	if err != nil {
		logrus.Fatal(err)
	}
}

func loadInitialShields(shieldCollection mongo.ShieldCollection) {
	err := mongo.InitializeShields(shieldCollection)
	if err != nil {
		logrus.Fatal(err)
	}
}
