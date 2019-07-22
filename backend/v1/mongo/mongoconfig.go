package mongo

import (
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	DatabaseURL             = "localhost"
	DatabaseName            = "wrm"
	ItemCollectionName      = "equipment"
	SpellCollectionName     = "spell"
	CharacterCollectionName = "character"
	BackendPort             = "8000"
)

//GetSession creates and returns an mgo session to the MongoDb at the specified URL.
func GetSession(URL string) *mgo.Session {
	session, err := mgo.Dial(URL)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Infof("MongoDb connected at %v", URL)
	}
	return session
}
