package mongo

import (
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

func GetSession(URL string) *mgo.Session {
	session, err := mgo.Dial(URL)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Infof("MongoDb connected at %v", URL)
	}
	return session
}
