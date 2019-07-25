package resources

import (
	"strconv"

	"github.com/lasjen88/WRMApp/backend/v1/models"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

//ReadItemsFromFile fetches the items from the resource file
func ReadItemsFromFile(path string) ([]models.Item, error) {
	var reader ResourceFileReader
	reader.SetReaderByPath(path)
	lines, err := reader.GetLines()
	if err != nil {
		return nil, err
	}
	var items []models.Item
	for _, line := range lines {
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
