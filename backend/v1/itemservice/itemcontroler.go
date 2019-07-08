package itemservice

import (
	"github.com/gorilla/mux"
)

//SetRouteHandles sets the item and spell handles on the router
func SetRouteHandles(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/items", GetItems).Methods("GET")
	router.HandleFunc("/v1/items", CreateItem).Methods("POST")

	router.HandleFunc("/v1/spells", GetSpells).Methods("GET")
	router.HandleFunc("/v1/spells", CreateSpell).Methods("POST")

	return router
}
