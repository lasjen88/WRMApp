package resources

import (
	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//ReadRaceFromFile fetches the races from the resource file
func ReadRaceFromFile(path string) ([]models.Race, error) {
	var reader ResourceFileReader
	reader.SetReaderByPath(path)
	lines, err := reader.GetLines()
	if err != nil {
		return nil, err
	}
	var races []models.Race
	for _, line := range lines {
		var valueError error
		races, valueError = parseRace(line, races)
		if valueError != nil {
			return races, valueError
		}
	}
	return races, nil
}

func parseRace(line []string, races []models.Race) ([]models.Race, error) {
	language := models.Language{LanguageName: line[3]}
	races = append(races, models.Race{RaceName: line[0],
		RaceDescription: line[2],
		RaceLanguages:   []models.Language{language},
	})
	return races, nil
}
