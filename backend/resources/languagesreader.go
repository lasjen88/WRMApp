package resources

import (
	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//ReadLanguageFromFile fetches the languages from the resource file
func ReadLanguageFromFile(path string) ([]models.Language, error) {
	var reader ResourceFileReader
	reader.SetReaderByPath(path)
	lines, err := reader.GetLines()
	if err != nil {
		return nil, err
	}
	var languages []models.Language
	for _, line := range lines {
		languages = append(languages, models.Language{LanguageName: line[0]})
	}
	return languages, nil
}
