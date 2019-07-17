package resources

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//ReadItemsFromFile fetches the items from the resource file
func ReadItemsFromFile(path string) ([]models.Item, error) {
	csvFile, openFileError := os.Open(path)
	if openFileError != nil {
		dir, _ := os.Getwd()
		logrus.Warningf("Please confirm that the specified path is relative to the current working directory. Current directory: %v", dir)
		return nil, openFileError
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
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

func parseEquipment(line []string, items []models.Item) ([]models.Item, error) {
	cost, costError := strconv.Atoi(line[1])
	if costError != nil {
		return items, costError
	}
	items = append(items, models.Item{
		ItemIdentifyer:  models.ItemIdentifyer{line[0], line[2]},
		ItemName:        line[0],
		ItemDescription: "",
		ItemCost:        cost,
	},
	)
	return items, nil
}
