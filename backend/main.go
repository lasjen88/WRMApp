package main

import (
	"net/http"

	"github.com/globalsign/mgo"
	routerconfig "github.com/lasjen88/WRMApp/backend/router"
	"github.com/lasjen88/WRMApp/backend/v1/characterservice"
	"github.com/lasjen88/WRMApp/backend/v1/initiativeservice"
	"github.com/lasjen88/WRMApp/backend/v1/itemservice"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	mongoSession := mongo.GetSession(mongo.DatabaseURL)
	router := mux.NewRouter()
	router = setRouteHandles(router, mongoSession)
	defer mongoSession.Close()
	logrus.Fatal(http.ListenAndServe(":"+routerconfig.BackendPort, router))
}

func setRouteHandles(router *mux.Router, mongoSession *mgo.Session) *mux.Router {
	router = itemservice.SetRouteHandles(router)
	router = characterservice.SetRouteHandles(router)
	router = initiativeservice.SetRouteHandles(router)
	return router
}
