package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//ItemCollection holds the data collection
type ItemCollection struct {
	DatabaseName, CollectionName string
	Session                      *mgo.Session
	itemCollection               *mgo.Collection
}

//IsEmptyCollection checks if the collection is currently empty
func (i *ItemCollection) IsEmptyCollection() (bool, error) {
	if i.itemCollection == nil {
		i.setupCollection()
	}
	var items []models.Item
	err := i.itemCollection.Find(nil).All(&items)
	if err != nil {
		return true, err
	}
	return len(items) == 0, nil
}

//GetAllItems fetches all items currently in the database
func (i *ItemCollection) GetAllItems() ([]models.Item, error) {
	if i.itemCollection == nil {
		i.setupCollection()
	}
	var items []models.Item
	err := i.itemCollection.Find(nil).All(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

//PutItems inserts an item slice to the database
func (i *ItemCollection) PutItems(newItems []models.Item) error {
	for _, newItem := range newItems {
		err := i.PutItem(newItem)
		if err != nil {
			return err
		}
	}
	return nil
}

//PutItem inserts an item to the database
func (i *ItemCollection) PutItem(newItem models.Item) error {
	if i.itemCollection == nil {
		i.setupCollection()
	}
	err := i.itemCollection.Insert(newItem)
	if err != nil {
		return err
	}
	return nil
}

func (i *ItemCollection) setupCollection() {
	i.itemCollection = i.Session.DB(i.DatabaseName).C(i.CollectionName)
}
