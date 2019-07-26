package talentservice

import (
	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//SetRouteHandles sets the talent handles on the router
func SetRouteHandles(router *mux.Router) *mux.Router {
	collection := mongo.WrmCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.TalentCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	talentCollection := mongo.TalentCollection{Wrmcollection: collection}
	loadInitialTalents(talentCollection)
	talenthandle := TalentHandle{TalentCollection: talentCollection}
	router.HandleFunc("/v1/talents", talenthandle.GetTalents).Methods("GET")
	return router
}

func loadInitialTalents(talentCollection mongo.TalentCollection) {
	err := mongo.InitializeTalents(talentCollection)
	if err != nil {
		logrus.Fatal(err)
	}
}
