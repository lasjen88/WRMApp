package characterservice

import (
	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
)

//SetRouteHandles sets the character handles on the router
func SetRouteHandles(router *mux.Router) *mux.Router {
	characterCollection := mongo.CharacterCollection{
		DatabaseName:   mongo.DatabaseName,
		CollectionName: mongo.CharacterCollectionName,
		Session:        mongo.GetSession(mongo.DatabaseURL),
	}
	characterhandle := CharacterHandle{CharacterCollection: characterCollection}
	router.HandleFunc("/v1/characters", characterhandle.GetCharacters).Methods("GET")
	router.HandleFunc("/v1/characters", characterhandle.CreateCharacter).Methods("POST")
	router.HandleFunc("/v1/characters/{id}", characterhandle.GetCharacter).Methods("GET")
	router.HandleFunc("/v1/characters/{id}", characterhandle.UpdateCharacter).Methods("PUT")
	router.HandleFunc("/v1/characters/{id}", characterhandle.DeleteCharacter).Methods("DELETE")
	return router
}
