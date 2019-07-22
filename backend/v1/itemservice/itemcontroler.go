package itemservice

import (
	"github.com/gorilla/mux"
)

//SetRouteHandles sets the item and spell handles on the router
func SetRouteHandles(router *mux.Router, itemHandle ItemHandle, spellHandle SpellHandle) *mux.Router {
	router.HandleFunc("/v1/items", itemHandle.GetItems).Methods("GET")
	router.HandleFunc("/v1/items", itemHandle.CreateItem).Methods("POST")

	router.HandleFunc("/v1/spells", spellHandle.GetSpells).Methods("GET")
	router.HandleFunc("/v1/spells", spellHandle.CreateSpell).Methods("POST")
	router.HandleFunc("/v1/spells/{circle}", spellHandle.GetCircleSpells).Methods("GET")

	return router
}
