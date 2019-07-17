package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/lasjen88/WRMApp/backend/resources"
	"github.com/sirupsen/logrus"
)

//InitializeEquipment hahahaha
func InitializeEquipment(mongoSession *mgo.Session) error {
	GetSession("localhost")
	collection := ItemCollection{DatabaseName: "data", CollectionName: "equipment", Session: mongoSession}
	isEmpty, err := collection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if isEmpty {
		return loadEquipment(collection)
	}
	return nil
}

func loadEquipment(collection ItemCollection) error {
	items, err := resources.ReadItemsFromFile("./backend/resources/equipment.csv")
	if err != nil {
		return err
	}
	logrus.Infof("Found %d pices of equipment.", len(items))
	collection.PutItems(items)
	logrus.Info("Load complete")
	return nil
}
