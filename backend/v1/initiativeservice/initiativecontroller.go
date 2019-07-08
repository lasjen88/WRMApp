package initiativeservice

import "github.com/gorilla/mux"

func setRouteHandles(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/initiative", GetInitiative).Methods("GET")
	return router
}
