package mongo

import (
	"fmt"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

func PrintCollectionNames(DB *mgo.Database) {
	collections, err := DB.CollectionNames()

	if err == nil {
		log.Warnf("Collections not found for Database : %v", DB.Name)
	}

	for i, collection := range collections {
		fmt.Printf("%3v - %v \n", i+1, collection)
	}
}
