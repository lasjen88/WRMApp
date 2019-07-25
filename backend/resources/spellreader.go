package resources

import (
	"github.com/lasjen88/WRMApp/backend/v1/models"
	"github.com/lasjen88/WRMApp/wrmrules"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

//ReadSpellsFromFile fetches the spells from the resource file
func ReadSpellsFromFile(path string, circle int) ([]models.Spell, error) {
	var reader ResourceFileReader
	reader.SetReaderByPath(path)
	lines, err := reader.GetLines()
	if err != nil {
		return nil, err
	}
	var spells []models.Spell
	for _, line := range lines {
		var costError error
		spells, costError = parseSpell(line, circle, spells)
		if costError != nil {
			return nil, costError
		}
	}
	return spells, nil
}

func parseSpell(line []string, circle int, spells []models.Spell) ([]models.Spell, error) {
	uid, uidError := uuid.NewV4()
	if uidError != nil {
		logrus.Warning(uidError)
		return nil, uidError
	}
	spells = append(spells, models.Spell{
		SpellName:        line[0],
		DifficultyLevel:  wrmrules.GetDifficultyLevel(circle),
		ManaConsumption:  wrmrules.GetManaConsumption(circle),
		SpellCost:        wrmrules.GetCost(circle),
		SpellDescription: line[1],
		ItemIdentifyer: models.ItemIdentifyer{
			ID:     uid.String(),
			Source: line[2],
		},
	})
	return spells, nil
}
