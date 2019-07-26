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

func TestReadShieldsFile(t *testing.T) {
	shields, err := resources.ReadArmorFromFile("../resources/shields.csv")
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(shields) != 3 {
		t.Errorf("Expected 3 shields, found %d", len(shields))
	}
}

func TestReadLanguageFile(t *testing.T) {
	language, err := resources.ReadLanguageFromFile("../resources/languages.csv")
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(language) != 9 {
		t.Errorf("Expected 9 languages, found %d", len(language))
	}
}

func TestReadTalentFile(t *testing.T) {
	talents, err := resources.ReadTalentsFromFile("../resources/talents.csv")
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(talents) != 48 {
		t.Errorf("Expected 48 talents, found %d", len(talents))
	}
}

func TestJsjsjs(t *testing.T) {
	uhhh := "1d - 2"
	haha := uhhh[5:6]
	if haha != "2" {
		t.Errorf("Oh no... it was %s", haha)
	}
}

func TestReadSkillFile(t *testing.T) {
	skills, err := resources.ReadSkillsFromFile("../resources/skills.csv")
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(skills) != 21 {
		t.Errorf("Expected 21 skills, found %d", len(skills))
	}
	for _, skill := range skills {
		if skill.Attribute == -1 {
			t.Errorf("The attribute of %s was set to -1!", skill.SkillName)
		}
	}
}

func TestReadRaceFile(t *testing.T) {
	races, err := resources.ReadRaceFromFile("../resources/races.csv")
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(races) != 13 {
		t.Errorf("Expected 13 races, found %d", len(races))
	}
}

func TestReadWeaponFile(t *testing.T) {
	weapons, err := resources.ReadWeaponsFromFile("../resources/weapons.csv")
	if err != nil {
		t.Errorf("An error was thrown: %v", err)
	}
	if len(weapons) != 21 {
		t.Errorf("Expected 21 races, found %d", len(weapons))
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
