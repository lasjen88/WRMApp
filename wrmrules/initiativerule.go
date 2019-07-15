package wrmrules

import (
	"errors"
	"sort"

	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//GetCharacterInitiatives provides the characters with initiatives
func GetCharacterInitiatives(characters []models.Character) (models.Initiatives, error) {
	initiatives := make(models.Initiatives, 0)
	for _, c := range characters {
		initiative, err := getInitiative(c)
		if err != nil {
			return nil, err
		}
		initiatives = append(initiatives, initiative)
	}
	sort.Sort(sort.Reverse(initiatives))
	return initiatives, nil
}

func getInitiative(character models.Character) (models.Initiative, error) {
	if character.CharacterName == "" {
		return models.Initiative{}, errors.New("Character must have a name to get initiative")
	}
	return models.Initiative{
		CharacterName:   character.CharacterName,
		InitiativeValue: character.Rogue + RollDSix(),
	}, nil
}
