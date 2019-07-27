package mongo

import "github.com/lasjen88/WRMApp/backend/v1/models"

//ShieldCollection holds the data collection
type ShieldCollection struct {
	Wrmcollection WrmCollection
}

//GetAllShields fetches all armors in the collection
func (s *ShieldCollection) GetAllShields() ([]models.Armor, error) {
	if s.Wrmcollection.collection == nil {
		s.Wrmcollection.SetupCollection()
	}
	var shields []models.Armor
	err := s.Wrmcollection.collection.Find(nil).All(&shields)
	if err != nil {
		return nil, err
	}
	return shields, nil
}

//PutShield inserts a shield to the collection
func (s *ShieldCollection) PutShield(newShield models.Armor) error {
	if s.Wrmcollection.collection == nil {
		s.Wrmcollection.SetupCollection()
	}
	err := s.Wrmcollection.collection.Insert(newShield)
	if err != nil {
		return err
	}
	return nil
}

//IsEmptyCollection checks if the collection is currently empty
func (s *ShieldCollection) IsEmptyCollection() (bool, error) {
	shields, err := s.GetAllShields()
	if err != nil {
		return true, err
	}
	return len(shields) == 0, nil
}
