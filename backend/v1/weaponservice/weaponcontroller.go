package weaponservice

import (
	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//SetRouteHandles sets the talent handles on the router
func SetRouteHandles(router *mux.Router) *mux.Router {
	collection := mongo.WrmCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.WeaponCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	weaponCollection := mongo.WeaponCollection{Wrmcollection: collection}
	loadInitialWeapons(weaponCollection)
	weaponhandle := WeaponHandle{WeaponCollection: weaponCollection}
	router.HandleFunc("/v1/weapons", weaponhandle.GetWeapons).Methods("GET")
	return router
}

func loadInitialWeapons(weaponCollection mongo.WeaponCollection) {
	collection := mongo.WrmCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.SkillCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	skillCollection := mongo.SkillCollection{Wrmcollection: collection}
	err := mongo.InitializeWeapons(weaponCollection, skillCollection)
	if err != nil {
		logrus.Fatal(err)
	}
}
