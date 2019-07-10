package wrmrules

import (
	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//CalculateHealthPoints gets the initial maximum health of a character
func CalculateHealthPoints(character models.Character) int {
	return character.Warrior + 6
}

//CalculateFate gets the initial fate of a character
func CalculateFate(character models.Character) int {
	return character.Rogue
}

//CalculateMana gets the initial maximum mana of a character
func CalculateMana(character models.Character) int {
	return character.Mage * 2
}

//CalculateBaseDefence gets the initial base defence of a character
func CalculateBaseDefence(character models.Character) int {
	return ((character.Warrior + character.Rogue) / 2) + 4
}
