package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//SpellCollection holds the data collection
type SpellCollection struct {
	DatabaseName, CollectionName string
	Session                      *mgo.Session
	spellCollection              *mgo.Collection
}

//IsEmptyCollection checks if the collection is currently empty
func (s *SpellCollection) IsEmptyCollection() (bool, error) {
	if s.spellCollection == nil {
		s.setupCollection()
	}
	var spells []models.Spell
	err := s.spellCollection.Find(nil).All(&spells)
	if err != nil {
		return true, err
	}
	return len(spells) == 0, nil
}

//GetAllSpells fetches all spells currently in the database
func (s *SpellCollection) GetAllSpells() ([]models.Spell, error) {
	if s.spellCollection == nil {
		s.setupCollection()
	}
	var spells []models.Spell
	err := s.spellCollection.Find(nil).All(&spells)
	if err != nil {
		return nil, err
	}
	return spells, nil
}

//GetCircleSpells fetches all spells from a circle
func (s *SpellCollection) GetCircleSpells(circle int) ([]models.Spell, error) {
	if s.spellCollection == nil {
		s.setupCollection()
	}
	var spells []models.Spell
	err := s.spellCollection.Find(bson.M{"spellCost": circle * 25}).All(&spells)
	if err != nil {
		return nil, err
	}
	return spells, nil
}

//PutSpells inserts a spell slice to the database
func (s *SpellCollection) PutSpells(newSpells []models.Spell) error {
	for _, newSpell := range newSpells {
		err := s.PutSpell(newSpell)
		if err != nil {
			return err
		}
	}
	return nil
}

//PutSpell inserts a spell to the database
func (s *SpellCollection) PutSpell(newSpell models.Spell) error {
	if s.spellCollection == nil {
		s.setupCollection()
	}
	err := s.spellCollection.Insert(newSpell)
	if err != nil {
		return err
	}
	return nil
}

func (s *SpellCollection) setupCollection() {
	s.spellCollection = s.Session.DB(s.DatabaseName).C(s.CollectionName)
}
