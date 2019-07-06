package mongo

import (
	"fmt"

	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

func PrintCollections(DB *mgo.Database) {
	collections, err := DB.CollectionNames()

	if err != nil {
		log.Warnf("Collections not found for Database : %v", DB.Name)
		return
	}

	for i, collection := range collections {
		fmt.Printf("%3v - %v \n", i+1, collection)
		printDocuments(DB, collection)
	}
}

func printDocuments(DB *mgo.Database, collection string) {
	coll := DB.C(collection)

	var result []interface{}
	coll.Find(nil).All(&result)
	for i, document := range result {
		fmt.Printf("\tDocuments %3v -\t%v\n", i+1, document)
	}

}
