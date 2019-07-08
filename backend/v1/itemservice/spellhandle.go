package itemservice

import (
	"encoding/json"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/v1/models"
	"github.com/lasjen88/WRMApp/wrmrules"
	log "github.com/sirupsen/logrus"
)

type spellDto struct {
	SpellName        string `json:"spellname"`
	SpellDescription string `json:"spelldescription"`
	Circle           int    `json:"circle"`
}

func mockReadAllSpallsDatabase() []models.Spell {
	spells := make([]models.Spell, 0)
	spells = append(spells, models.Spell{SpellName: "Summon Rat", DifficultyLevel: 4, ManaConsumption: 1, SpellCost: 25, SpellDescription: "Summons a rat"})
	spells = append(spells, models.Spell{SpellName: "Summon Cat", DifficultyLevel: 6, ManaConsumption: 2, SpellCost: 50, SpellDescription: "Summons a cat"})
	spells = append(spells, models.Spell{SpellName: "Summon Horse", DifficultyLevel: 9, ManaConsumption: 4, SpellCost: 75, SpellDescription: "Summons a horse"})
	spells = append(spells, models.Spell{SpellName: "Summon Elephant", DifficultyLevel: 25, ManaConsumption: 8, SpellCost: 100, SpellDescription: "Summons an elephant"})
	return spells
}

//GetSpells provides all spells from the database, regardless of circle
func GetSpells(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(ContentTypeHeaderKey, ContentTypeHeaderValue)
	spells := mockReadAllSpallsDatabase()
	json.NewEncoder(writer).Encode(spells)
}

//CreateSpell creates a spell
func CreateSpell(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set(ContentTypeHeaderKey, ContentTypeHeaderValue)
	var spellDto spellDto
	err := json.NewDecoder(request.Body).Decode(&spellDto)
	if err != nil {
		log.Error(err)
		return
	}
	spell := models.Spell{
		SpellName:        spellDto.SpellName,
		DifficultyLevel:  wrmrules.GetDifficultyLevel(spellDto.Circle),
		ManaConsumption:  wrmrules.GetManaConsumption(spellDto.Circle),
		SpellCost:        wrmrules.GetCost(spellDto.Circle),
		SpellDescription: spellDto.SpellDescription,
	}
	json.NewEncoder(writer).Encode(spell)
}