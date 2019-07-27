package mongo

import "github.com/lasjen88/WRMApp/backend/v1/models"

//WeaponCollection holds the data collection
type WeaponCollection struct {
	Wrmcollection WrmCollection
}

//GetAllWeapons fetches all weapons in the collection
func (w *WeaponCollection) GetAllWeapons() ([]models.Weapon, error) {
	if w.Wrmcollection.collection == nil {
		w.Wrmcollection.SetupCollection()
	}
	var weapons []models.Weapon
	err := w.Wrmcollection.collection.Find(nil).All(&weapons)
	if err != nil {
		return nil, err
	}
	return weapons, nil
}

//PutWeapon inserts a weapon to the collection
func (w *WeaponCollection) PutWeapon(newWeapon models.Weapon) error {
	if w.Wrmcollection.collection == nil {
		w.Wrmcollection.SetupCollection()
	}
	err := w.Wrmcollection.collection.Insert(newWeapon)
	if err != nil {
		return err
	}
	return nil
}

//IsEmptyCollection checks if the collection is currently empty
func (w *WeaponCollection) IsEmptyCollection() (bool, error) {
	weapons, err := w.GetAllWeapons()
	if err != nil {
		return true, err
	}
	return len(weapons) == 0, nil
}
