package main

import (
	"net/http"

	"github.com/globalsign/mgo"
	routerconfig "github.com/lasjen88/WRMApp/backend/router"
	"github.com/lasjen88/WRMApp/backend/v1/armorservice"
	"github.com/lasjen88/WRMApp/backend/v1/characterservice"
	"github.com/lasjen88/WRMApp/backend/v1/initiativeservice"
	"github.com/lasjen88/WRMApp/backend/v1/itemservice"
	"github.com/lasjen88/WRMApp/backend/v1/languageservice"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/lasjen88/WRMApp/backend/v1/raceservice"
	"github.com/lasjen88/WRMApp/backend/v1/skillservice"
	"github.com/lasjen88/WRMApp/backend/v1/talentservice"
	"github.com/lasjen88/WRMApp/backend/v1/weaponservice"
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
	router = skillservice.SetRouteHandles(router)
	router = talentservice.SetRouteHandles(router)
	router = languageservice.SetRouteHandles(router)
	router = armorservice.SetRouteHandles(router)
	router = weaponservice.SetRouteHandles(router)
	router = raceservice.SetRouteHandles(router)
	return router
}
