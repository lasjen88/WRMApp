package resourcestest

import (
	"testing"

	"github.com/lasjen88/WRMApp/backend/resources"
)

func TestReadEquipmentFile(t *testing.T) {
	items, err := resources.ReadItemsFromFile("../resources/equipment.csv")
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(items) != 31 {
		t.Errorf("Expected 31 items, found %d", len(items))
	}
}
