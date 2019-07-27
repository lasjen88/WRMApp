package mongo

import "github.com/lasjen88/WRMApp/backend/v1/models"

//RaceCollection holds the data collection
type RaceCollection struct {
	Wrmcollection WrmCollection
}

//GetAllRaces fetches all races in the collection
func (r *RaceCollection) GetAllRaces() ([]models.Race, error) {
	if r.Wrmcollection.collection == nil {
		r.Wrmcollection.SetupCollection()
	}
	var races []models.Race
	err := r.Wrmcollection.collection.Find(nil).All(&races)
	if err != nil {
		return nil, err
	}
	return races, nil
}

//PutRace inserts a race to the collection
func (r *RaceCollection) PutRace(newRace models.Race) error {
	if r.Wrmcollection.collection == nil {
		r.Wrmcollection.SetupCollection()
	}
	err := r.Wrmcollection.collection.Insert(newRace)
	if err != nil {
		return err
	}
	return nil
}

//IsEmptyCollection checks if the collection is currently empty
func (r *RaceCollection) IsEmptyCollection() (bool, error) {
	races, err := r.GetAllRaces()
	if err != nil {
		return true, err
	}
	return len(races) == 0, nil
}
