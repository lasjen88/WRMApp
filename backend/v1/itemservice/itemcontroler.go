package itemservice

import (
	"github.com/gorilla/mux"
)

func setRouteHandles(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/items", GetItems).Methods("GET")
	router.HandleFunc("/v1/items", CreateItem).Methods("POST")
	return router
}
