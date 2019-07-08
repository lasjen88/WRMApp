package wrmrulestest

import (
	"testing"

	"github.com/lasjen88/WRMApp/backend/v1/models"
	"github.com/lasjen88/WRMApp/wrmrules"
)

const (
	highestPossibleInitiative int = 12
)

func TestInitativeOrder(t *testing.T) {
	carl := models.Character{CharacterName: "Carl", Rogue: 6}
	bert := models.Character{CharacterName: "Bert", Rogue: 3}
	paul := models.Character{CharacterName: "Paul", Rogue: 0}
	characters := []models.Character{carl, bert, paul}
	expectedInitiativeCount := 3

	initiatives, _ := wrmrules.GetCharacterInitiatives(characters)

	if initiatives.Len() != expectedInitiativeCount {
		t.Errorf("Expected %d initiatives, but found %d.", expectedInitiativeCount, initiatives.Len())
	}

	lastInitiative := highestPossibleInitiative
	for i, init := range initiatives {
		if init.InitiativeValue > lastInitiative {
			t.Errorf("Something whent wrong in the initiative order at index %d: Previuse value: %d - This value: %d", i, lastInitiative, init.InitiativeValue)
		}
		lastInitiative = init.InitiativeValue
	}
}

func TestCharacterWithoutName(t *testing.T) {
	nameless := models.Character{Rogue: 6}
	characters := []models.Character{nameless}
	_, err := wrmrules.GetCharacterInitiatives(characters)
	if err == nil {
		t.Errorf("This should have throw an error")
	}
}
