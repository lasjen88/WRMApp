package resources

import (
	"strings"

	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//RaceDto is a dto for the race resource
type RaceDto struct {
	Race    models.Race
	Talents []string
}

//ReadRaceFromFile fetches the races from the resource file
func ReadRaceFromFile(path string) ([]RaceDto, error) {
	var reader ResourceFileReader
	reader.SetReaderByPath(path)
	lines, err := reader.GetLines()
	if err != nil {
		return nil, err
	}
	var races []RaceDto
	for _, line := range lines {
		var valueError error
		races, valueError = parseRace(line, races)
		if valueError != nil {
			return races, valueError
		}
	}
	return races, nil
}

func parseRace(line []string, races []RaceDto) ([]RaceDto, error) {
	language := models.Language{LanguageName: line[3]}
	modelRace := models.Race{RaceName: line[0],
		RaceDescription: line[1],
		RaceLanguages:   []models.Language{language},
	}
	talentString := strings.Split(line[2], ",")
	races = append(races, RaceDto{
		Race:    modelRace,
		Talents: talentString,
	})
	return races, nil
}
