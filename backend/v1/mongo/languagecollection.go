package mongo

import "github.com/lasjen88/WRMApp/backend/v1/models"

//LanguageCollection holds the data collection
type LanguageCollection struct {
	Wrmcollection WrmCollection
}

//GetAllLanguages fetches all languages in the collection
func (l *LanguageCollection) GetAllLanguages() ([]models.Language, error) {
	if l.Wrmcollection.collection == nil {
		l.Wrmcollection.SetupCollection()
	}
	var languages []models.Language
	err := l.Wrmcollection.collection.Find(nil).All(&languages)
	if err != nil {
		return nil, err
	}
	return languages, nil
}

//PutLanguage inserts a language to the collection
func (l *LanguageCollection) PutLanguage(newLanguage models.Language) error {
	if l.Wrmcollection.collection == nil {
		l.Wrmcollection.SetupCollection()
	}
	err := l.Wrmcollection.collection.Insert(newLanguage)
	if err != nil {
		return err
	}
	return nil
}

//IsEmptyCollection checks if the collection is currently empty
func (l *LanguageCollection) IsEmptyCollection() (bool, error) {
	languages, err := l.GetAllLanguages()
	if err != nil {
		return true, err
	}
	return len(languages) == 0, nil
}
