package models

//Initiative stores the random initiative with the character name
type Initiative struct {
	CharacterName   string `json:"charactername"`
	InitiativeValue int    `json:"initiative"`
}

//Initiatives stores a slice of initiatives
type Initiatives []Initiative

func (i Initiatives) Len() int {
	return len(i)
}

func (i Initiatives) Less(this, other int) bool {
	return i[this].InitiativeValue < i[other].InitiativeValue
}

func (i Initiatives) Swap(this, other int) {
	i[this], i[other] = i[other], i[this]
}
