package mongo

import (
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

func Use(session *mgo.Session, DBName string) *mgo.Database {
	DB := session.DB(DBName)

	if DB == nil {
		log.Errorf("Database %v not found", DBName)
	}

	log.Infof("Connected to Database: %v", DBName)

	return DB
}
