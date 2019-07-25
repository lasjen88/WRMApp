package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//CharacterCollection holds the data collection
type CharacterCollection struct {
	DatabaseName, CollectionName string
	Session                      *mgo.Session
	charcterCollection           *mgo.Collection
}

//GetAllCharacters fetches all characters in the collection
func (c *CharacterCollection) GetAllCharacters() ([]models.Character, error) {
	if c.charcterCollection == nil {
		c.setupCollection()
	}
	var characters []models.Character
	err := c.charcterCollection.Find(nil).All(&characters)
	if err != nil {
		return nil, err
	}
	return characters, nil
}

//PutCharacter inserts a character to the collection
func (c *CharacterCollection) PutCharacter(newCharacter models.Character) error {
	if c.charcterCollection == nil {
		c.setupCollection()
	}
	err := c.charcterCollection.Insert(newCharacter)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCharacter delete an existing character from the collection
func (c *CharacterCollection) DeleteCharacter(characterid string) error {
	if c.charcterCollection == nil {
		c.setupCollection()
	}
	err := c.charcterCollection.Remove(bson.M{"id": characterid})
	if err != nil {
		return err
	}
	return nil
}

//UpdateCharacter updates an existing character from the collection, by replacing it with a new one. The id is persisted.
func (c *CharacterCollection) UpdateCharacter(characterid string, character models.Character) error {
	if c.charcterCollection == nil {
		c.setupCollection()
	}
	err := c.charcterCollection.Update(bson.M{"id": characterid}, character)
	if err != nil {
		return err
	}
	return nil
}

func (c *CharacterCollection) setupCollection() {
	c.charcterCollection = c.Session.DB(c.DatabaseName).C(c.CollectionName)
}
