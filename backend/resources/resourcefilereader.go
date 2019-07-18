package resources

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/lasjen88/WRMApp/wrmrules"

	uuid "github.com/satori/go.uuid"

	"github.com/sirupsen/logrus"

	"github.com/lasjen88/WRMApp/backend/v1/models"
)

const (
	delimiter = ';'
)

//ReadItemsFromFile fetches the items from the resource file
func ReadItemsFromFile(path string) ([]models.Item, error) {
	csvFile := openCsvFile(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = delimiter
	var items []models.Item
	for {
		line, readError := reader.Read()
		if readError == io.EOF {
			break
		} else if readError != nil {
			return nil, readError
		}
		var costError error
		items, costError = parseEquipment(line, items)
		if costError != nil {
			return nil, costError
		}
	}
	return items, nil
}

//ReadSpellsFromFile fetches the spells from the resource file
func ReadSpellsFromFile(path string, circle int) ([]models.Spell, error) {
	csvFile := openCsvFile(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = delimiter
	var spells []models.Spell
	for {
		line, readError := reader.Read()
		if readError == io.EOF {
			break
		} else if readError != nil {
			return nil, readError
		}
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

func parseEquipment(line []string, items []models.Item) ([]models.Item, error) {
	cost, costError := strconv.Atoi(line[1])
	if costError != nil {
		return items, costError
	}
	uid, uidError := uuid.NewV4()
	if uidError != nil {
		logrus.Warning(uidError)
		return nil, uidError
	}
	items = append(items, models.Item{
		ItemIdentifyer: models.ItemIdentifyer{
			ID:     uid.String(),
			Source: line[2],
		},
		ItemName:        line[0],
		ItemDescription: "",
		ItemCost:        cost,
	},
	)
	return items, nil
}

func openCsvFile(path string) *os.File {
	csvFile, openFileError := os.Open(path)
	if openFileError != nil {
		dir, _ := os.Getwd()
		logrus.Fatalf("Please confirm that the specified path is relative to the current working directory. Current directory: %v", dir)
	}
	return csvFile
}
