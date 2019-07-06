package mongo

import (
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

//GetSession creates and returns an mgo session to the MongoDb at the specified URL.
func GetSession(URL string) *mgo.Session {
	session, err := mgo.Dial(URL)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("MongoDb connected")
	}
	return session
}
