package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/lasjen88/WRMApp/backend/resources"
	"github.com/sirupsen/logrus"
)

var equipmentFiles = [...]string{"./backend/resources/equipment.csv"}
var spellFiles = [...]string{
	"./backend/resources/firstCircleSpells.csv",
	"./backend/resources/secondCircleSpells.csv",
	"./backend/resources/thirdCircleSpells.csv",
	"./backend/resources/fourthCircleSpells.csv",
}

//InitializeEquipment preloads the database with equipment from the resource files
func InitializeEquipment(mongoSession *mgo.Session) error {
	GetSession("localhost")
	collection := ItemCollection{DatabaseName: "wrm", CollectionName: "equipment", Session: mongoSession}
	isEmpty, err := collection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if isEmpty {
		for _, path := range equipmentFiles {
			err = loadEquipment(collection, path)
			if err != nil {
				return err
			}
		}
		return nil
	}
	logrus.Info("Equipment already loaded.")
	return nil
}

//InitializeSpells preloads the database with spells from the resource files
func InitializeSpells(mongoSession *mgo.Session) error {
	GetSession("localhost")
	collection := SpellCollection{DatabaseName: "wrm", CollectionName: "spell", Session: mongoSession}
	isEmpty, err := collection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if isEmpty {
		for fileIndex, path := range spellFiles {
			err = loadSpells(collection, path, fileIndex+1)
			if err != nil {
				return err
			}
		}
		return nil
	}
	logrus.Info("Spells already loaded.")
	return nil
}

func loadSpells(collection SpellCollection, path string, circle int) error {
	spells, err := resources.ReadSpellsFromFile(path, circle)
	if err != nil {
		return err
	}
	logrus.Infof("Found %d spells in %s.", len(spells), path)
	collection.PutSpells(spells)
	logrus.Info("Spell load complete")
	return nil
}

func loadEquipment(collection ItemCollection, path string) error {
	items, err := resources.ReadItemsFromFile(path)
	if err != nil {
		return err
	}
	logrus.Infof("Found %d pices of equipment.", len(items))
	collection.PutItems(items)
	logrus.Info("Equipment load complete")
	return nil
}
