package wrmrulestest

import (
	"testing"

	"github.com/lasjen88/WRMApp/wrmrules"
)

func TestDifficultyLevelCircleOne(t *testing.T) {
	circle := 1
	expectedDifficultyLevel := 5
	difficultyLevel := wrmrules.GetDifficultyLevel(circle)
	if difficultyLevel != expectedDifficultyLevel {
		t.Errorf("Expected difficulty level %d, but found level %d.", expectedDifficultyLevel, difficultyLevel)
	}
}

func TestDifficultyLevelCircleTwo(t *testing.T) {
	circle := 2
	expectedDifficultyLevel := 7
	difficultyLevel := wrmrules.GetDifficultyLevel(circle)
	if difficultyLevel != expectedDifficultyLevel {
		t.Errorf("Expected difficulty level %d, but found level %d.", expectedDifficultyLevel, difficultyLevel)
	}
}

func TestDifficultyLevelCircleThree(t *testing.T) {
	circle := 3
	expectedDifficultyLevel := 9
	difficultyLevel := wrmrules.GetDifficultyLevel(circle)
	if difficultyLevel != expectedDifficultyLevel {
		t.Errorf("Expected difficulty level %d, but found level %d.", expectedDifficultyLevel, difficultyLevel)
	}
}

func TestDifficultyLevelCircleFour(t *testing.T) {
	circle := 4
	expectedDifficultyLevel := 13
	difficultyLevel := wrmrules.GetDifficultyLevel(circle)
	if difficultyLevel != expectedDifficultyLevel {
		t.Errorf("Expected difficulty level %d, but found level %d.", expectedDifficultyLevel, difficultyLevel)
	}
}

func TestDifficultyLevelCircleFive(t *testing.T) {
	circle := 5
	expectedDifficultyLevel := 0
	difficultyLevel := wrmrules.GetDifficultyLevel(circle)
	if difficultyLevel != expectedDifficultyLevel {
		t.Errorf("Expected difficulty level %d, but found level %d.", expectedDifficultyLevel, difficultyLevel)
	}
}
func TestManaConsumptionCircleOne(t *testing.T) {
	circle := 1
	expectedManaConsumption := 1
	manaConsumption := wrmrules.GetManaConsumption(circle)
	if manaConsumption != expectedManaConsumption {
		t.Errorf("Expected mana consumption %d, but found consumption %d.", expectedManaConsumption, manaConsumption)
	}
}

func TestManaConsumptionCircleTwo(t *testing.T) {
	circle := 2
	expectedManaConsumption := 2
	manaConsumption := wrmrules.GetManaConsumption(circle)
	if manaConsumption != expectedManaConsumption {
		t.Errorf("Expected mana consumption %d, but found consumption %d.", expectedManaConsumption, manaConsumption)
	}
}

func TestManaConsumptionCircleThree(t *testing.T) {
	circle := 3
	expectedManaConsumption := 4
	manaConsumption := wrmrules.GetManaConsumption(circle)
	if manaConsumption != expectedManaConsumption {
		t.Errorf("Expected mana consumption %d, but found consumption %d.", expectedManaConsumption, manaConsumption)
	}
}

func TestManaConsumptionCircleFour(t *testing.T) {
	circle := 4
	expectedManaConsumption := 8
	manaConsumption := wrmrules.GetManaConsumption(circle)
	if manaConsumption != expectedManaConsumption {
		t.Errorf("Expected mana consumption %d, but found consumption %d.", expectedManaConsumption, manaConsumption)
	}
}

func TestManaConsumptionCircleFive(t *testing.T) {
	circle := 5
	expectedManaConsumption := 0
	manaConsumption := wrmrules.GetManaConsumption(circle)
	if manaConsumption != expectedManaConsumption {
		t.Errorf("Expected mana consumption %d, but found consumption %d.", expectedManaConsumption, manaConsumption)
	}
}

func TestSpellCostCircleZero(t *testing.T) {
	runSpellCostTestForCircle(0, t)
}

func TestSpellCostCircleThree(t *testing.T) {
	runSpellCostTestForCircle(3, t)
}

func TestSpellCostCircleFive(t *testing.T) {
	runSpellCostTestForCircle(5, t)
}

func runSpellCostTestForCircle(circle int, t *testing.T) {
	expectedSpellCost := circle * 25
	spellCost := wrmrules.GetCost(circle)
	if spellCost != expectedSpellCost {
		t.Errorf("Expected spell cost %d, but found cost %d.", expectedSpellCost, spellCost)
	}
}
