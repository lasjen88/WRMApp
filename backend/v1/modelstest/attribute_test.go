package modelstest

import (
	"testing"

	"github.com/lasjen88/WRMApp/backend/v1/models"
)

func TestAttributeCapsWarriorValue(t *testing.T) {
	executeTest(t, "WARRIOR", models.Warrior)
}

func TestAttributeWarriorValue(t *testing.T) {
	executeTest(t, "Warrior", models.Warrior)
}

func TestAttributeLowerWarriorValue(t *testing.T) {
	executeTest(t, "warrior", models.Warrior)
}

func TestAttributeCapsRogueValue(t *testing.T) {
	executeTest(t, "ROGUE", models.Rogue)
}

func TestAttributeRogueValue(t *testing.T) {
	executeTest(t, "Rogue", models.Rogue)
}

func TestAttributeLowerRogueValue(t *testing.T) {
	executeTest(t, "rogue", models.Rogue)
}

func TestAttributeCapsMageValue(t *testing.T) {
	executeTest(t, "MAGE", models.Mage)
}

func TestAttributeMageValue(t *testing.T) {
	executeTest(t, "Mage", models.Mage)
}

func TestAttributeLowerMageValue(t *testing.T) {
	executeTest(t, "mage", models.Mage)
}

func executeTest(t *testing.T, attributeName string, expectedAttributeValue models.AttributeEnum) {
	attributeValue := models.Value(attributeName)
	if attributeValue != expectedAttributeValue {
		t.Errorf("Expected the value to be %d, but was %d", expectedAttributeValue, attributeValue)
	}
}
