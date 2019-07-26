package mongo

import (
	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//TalentCollection holds the data collection
type TalentCollection struct {
	talentCollection WrmCollection
}

//GetAllTalents fetches all talents in the collection
func (t *TalentCollection) GetAllTalents() ([]models.Talent, error) {
	if t.talentCollection.collection == nil {
		t.talentCollection.SetupCollection()
	}
	var talents []models.Talent
	err := t.talentCollection.collection.Find(nil).All(&talents)
	if err != nil {
		return nil, err
	}
	return talents, nil
}

//PutTalent inserts a skill to the collection
func (t *TalentCollection) PutTalent(newTalent models.Talent) error {
	if t.talentCollection.collection == nil {
		t.talentCollection.SetupCollection()
	}
	err := t.talentCollection.collection.Insert(newTalent)
	if err != nil {
		return err
	}
	return nil
}

//IsEmptyCollection checks if the collection is currently empty
func (t *TalentCollection) IsEmptyCollection() (bool, error) {
	talents, err := t.GetAllTalents()
	if err != nil {
		return true, err
	}
	return len(talents) == 0, nil
}
