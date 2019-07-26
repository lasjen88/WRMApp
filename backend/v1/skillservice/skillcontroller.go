package skillservice

import (
	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//SetRouteHandles sets the skill handles on the router
func SetRouteHandles(router *mux.Router) *mux.Router {
	collection := mongo.WrmCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.SkillCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	skillCollection := mongo.SkillCollection{Wrmcollection: collection}
	loadInitialSkills(skillCollection)
	skillhandle := SkillHandle{SkillCollection: skillCollection}
	router.HandleFunc("/v1/skills", skillhandle.GetSkills).Methods("GET")
	return router
}

func loadInitialSkills(skillCollection mongo.SkillCollection) {
	err := mongo.InitializeSkills(skillCollection)
	if err != nil {
		logrus.Fatal(err)
	}
}
