package characterservice

import "github.com/gorilla/mux"

//SetRouteHandles sets the character handles on the router
func SetRouteHandles(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/characters", GetCharacters).Methods("GET")
	router.HandleFunc("/v1/characters", CreateCharacter).Methods("POST")
	router.HandleFunc("/v1/characters/{id}", GetCharacter).Methods("GET")
	router.HandleFunc("/v1/characters/{id}", UpdateCharacter).Methods("PUT")
	router.HandleFunc("/v1/characters/{id}", DeleteCharacter).Methods("DELETE")
	return router
}
