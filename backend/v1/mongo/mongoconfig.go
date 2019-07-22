package mongo

import (
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	//DatabaseURL host name of the database server
	DatabaseURL = "localhost"
	//DatabaseName name of the mongo database
	DatabaseName = "wrm"
	//ItemCollectionName collection for equipment
	ItemCollectionName = "equipment"
	//SpellCollectionName collection for spells
	SpellCollectionName = "spell"
	//CharacterCollectionName collection for characters
	CharacterCollectionName = "character"
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
