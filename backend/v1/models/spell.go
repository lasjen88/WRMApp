package models

//Spell stores a spell
type Spell struct {
	SpellName        string         `json:"spellName"`
	DifficultyLevel  int            `json:"difficultyLevel"`
	ManaConsumption  int            `json:"manaConsumption"`
	SpellCost        int            `json:"spellCost"`
	SpellDescription string         `json:"spellDescription"`
	ItemIdentifyer   ItemIdentifyer `json:"itemidentifyer"`
}
