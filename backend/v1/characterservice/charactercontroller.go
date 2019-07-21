package characterservice

import "github.com/gorilla/mux"

//SetRouteHandles sets the character handles on the router
func SetRouteHandles(router *mux.Router, characterHandle CharacterHandle) *mux.Router {
	router.HandleFunc("/v1/characters", characterHandle.GetCharacters).Methods("GET")
	router.HandleFunc("/v1/characters", characterHandle.CreateCharacter).Methods("POST")
	router.HandleFunc("/v1/characters/{id}", characterHandle.GetCharacter).Methods("GET")
	//router.HandleFunc("/v1/characters/{id}", characterHandle.UpdateCharacter).Methods("PUT")
	router.HandleFunc("/v1/characters/{id}", characterHandle.DeleteCharacter).Methods("DELETE")
	return router
}
