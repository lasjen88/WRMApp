package mongo

import "github.com/lasjen88/WRMApp/backend/v1/models"

//ArmorCollection holds the data collection
type ArmorCollection struct {
	Wrmcollection WrmCollection
}

//GetAllArmor fetches all armors in the collection
func (a *ArmorCollection) GetAllArmor() ([]models.Armor, error) {
	if a.Wrmcollection.collection == nil {
		a.Wrmcollection.SetupCollection()
	}
	var armors []models.Armor
	err := a.Wrmcollection.collection.Find(nil).All(&armors)
	if err != nil {
		return nil, err
	}
	return armors, nil
}

//PutArmor inserts an armor to the collection
func (a *ArmorCollection) PutArmor(newArmor models.Armor) error {
	if a.Wrmcollection.collection == nil {
		a.Wrmcollection.SetupCollection()
	}
	err := a.Wrmcollection.collection.Insert(newArmor)
	if err != nil {
		return err
	}
	return nil
}

//IsEmptyCollection checks if the collection is currently empty
func (a *ArmorCollection) IsEmptyCollection() (bool, error) {
	armors, err := a.GetAllArmor()
	if err != nil {
		return true, err
	}
	return len(armors) == 0, nil
}
