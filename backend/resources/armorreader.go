package resources

import (
	"strconv"

	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//ReadArmorFromFile fetches the armor from the resource file
func ReadArmorFromFile(path string) ([]models.Armor, error) {
	var reader ResourceFileReader
	reader.SetReaderByPath(path)
	lines, err := reader.GetLines()
	if err != nil {
		return nil, err
	}
	var armor []models.Armor
	for _, line := range lines {
		var valueError error
		armor, valueError = parseArmor(line, armor)
		if valueError != nil {
			return nil, valueError
		}
	}
	return armor, nil
}

func parseArmor(line []string, armor []models.Armor) ([]models.Armor, error) {
	defence, penalty, cost, err := getIntegerValues(line)
	if err != nil {
		return armor, err
	}
	armor = append(armor,
		models.Armor{
			ArmorName:      line[0],
			ArmorDefence:   defence,
			ArmorPenalty:   penalty,
			ArmorCost:      cost,
			ItemIdentifyer: models.ItemIdentifyer{Source: line[4]},
		})
	return armor, nil
}

func getIntegerValues(line []string) (int, int, int, error) {
	defence, err := strconv.Atoi(line[1])
	var penalty int
	penalty, err = strconv.Atoi(line[2])
	cost, costErr := strconv.Atoi(line[3])
	if costErr != nil && line[3] == "" {
		cost = 0
	} else if costErr != nil {
		err = costErr
	}
	return defence, penalty, cost, err
}
