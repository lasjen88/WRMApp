package mongo

import (
	"fmt"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

func PrintDBNames(session *mgo.Session) {
	DBNames, err := session.DatabaseNames()

	if err != nil {
		log.Warn(err)
	}

	for i, v := range DBNames {
		fmt.Printf("%3v - %v\n", i+1, v)
	}
}
