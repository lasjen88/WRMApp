package mongo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lasjen88/WRMApp/backend/resources"
	"github.com/lasjen88/WRMApp/backend/v1/models"

	"github.com/sirupsen/logrus"
)

var raceFiles = [...]string{"./backend/resources/races.csv"}

//InitializRaces preloads the database with weapons from the resource files
func InitializRaces(raceCollection RaceCollection, talentCollection TalentCollection) error {
	racesIsEmpty, err := raceCollection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if !racesIsEmpty {
		logrus.Info("Races already loaded.")
		return nil
	}
	err = checkRequiredCollections(talentCollection)
	if err != nil {
		return err
	}
	for _, path := range raceFiles {
		err = loadRace(raceCollection, talentCollection, path)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadRace(raceCollection RaceCollection, talentCollection TalentCollection, path string) error {
	races, err := resources.ReadRaceFromFile(path)
	if err != nil {
		return err
	}
	logrus.Infof("Found %d races in %s.", len(races), path)
	talents, talentError := talentCollection.GetAllTalents()
	if talentError != nil {
		return talentError
	}
	for _, raceDto := range races {
		race, dtoError := getRaceFromDto(raceDto, talents)
		if dtoError != nil {
			return dtoError
		}
		err := raceCollection.PutRace(race)
		if err != nil {
			return err
		}
	}
	logrus.Info("Race load complete")
	return nil
}

func getRaceFromDto(raceDto resources.RaceDto, talents []models.Talent) (models.Race, error) {
	race := raceDto.Race
	for _, dtoTalent := range raceDto.Talents {
		talentName := strings.TrimSpace(dtoTalent)
		talentName = strings.ReplaceAll(talentName, "_", " ")
		talentFound := false
		if len(talentName) == 0 {
			talentFound = true
			break
		}
		for _, talent := range talents {
			knownTalentName := strings.ReplaceAll(talent.TalentName, "(", "")
			knownTalentName = strings.ReplaceAll(knownTalentName, ")", "")
			if strings.EqualFold(knownTalentName, talentName) {
				race.RaceTalents = append(race.RaceTalents, talent)
				talentFound = true
				break
			}
		}
		if !talentFound {
			var realTalentNames []string
			for _, talent := range talents {
				realTalentNames = append(realTalentNames, talent.TalentName)
			}
			errorMessage := fmt.Sprintf("RaceDto talent string '%s' does not match any talents in the database [%v]", talentName, strings.Join(realTalentNames, ", "))
			return race, errors.New(errorMessage)
		}
	}
	return race, nil
}

func checkRequiredCollections(talentCollection TalentCollection) error {
	talentsIsEmpty, talentDatabaseError := talentCollection.IsEmptyCollection()
	if talentDatabaseError != nil {
		return talentDatabaseError
	}
	if talentsIsEmpty {
		return errors.New("Cannot load races before talents have been loaded")
	}
	return nil
}
