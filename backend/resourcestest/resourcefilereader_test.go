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

func TestReadArmorFile(t *testing.T) {
	armor, err := resources.ReadArmorFromFile("../resources/armor.csv")
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(armor) != 9 {
		t.Errorf("Expected 9 armors, found %d", len(armor))
	}
}

func TestReadFirstSpellFile(t *testing.T) {
	items, err := resources.ReadSpellsFromFile("../resources/firstCircleSpells.csv", 1)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(items) != 24 {
		t.Errorf("Expected 31 items, found %d", len(items))
	}
}

func TestReadSecondSpellFile(t *testing.T) {
	items, err := resources.ReadSpellsFromFile("../resources/secondCircleSpells.csv", 2)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(items) != 7 {
		t.Errorf("Expected 31 items, found %d", len(items))
	}
}

func TestReadThirdSpellFile(t *testing.T) {
	items, err := resources.ReadSpellsFromFile("../resources/thirdCircleSpells.csv", 3)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(items) != 5 {
		t.Errorf("Expected 31 items, found %d", len(items))
	}
}

func TestReadFourthSpellFile(t *testing.T) {
	items, err := resources.ReadSpellsFromFile("../resources/fourthCircleSpells.csv", 4)
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(items) != 5 {
		t.Errorf("Expected 31 items, found %d", len(items))
	}
}
