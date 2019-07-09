package initiativeservice

import "github.com/gorilla/mux"

//SetRouteHandles sets the initiative handles on the router
func SetRouteHandles(router *mux.Router) *mux.Router {
	router.HandleFunc("/v1/initiative", GetInitiative).Methods("GET")
	return router
}
