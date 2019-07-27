package languageservice

import (
	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//SetRouteHandles sets the talent handles on the router
func SetRouteHandles(router *mux.Router) *mux.Router {
	collection := mongo.WrmCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.LanguageCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	languageCollection := mongo.LanguageCollection{Wrmcollection: collection}
	loadInitialLanguages(languageCollection)
	languagehandle := LanguageHandle{LanguageCollection: languageCollection}
	router.HandleFunc("/v1/languages", languagehandle.GetLanguages).Methods("GET")
	return router
}

func loadInitialLanguages(languageCollection mongo.LanguageCollection) {
	err := mongo.InitializeLanguages(languageCollection)
	if err != nil {
		logrus.Fatal(err)
	}
}
