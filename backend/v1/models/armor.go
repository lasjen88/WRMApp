package models

//Armor stores an armor item
type Armor struct {
	ArmorName    string `json:"armorName"`
	ArmorDefence int    `json:"armorDefence"`
	ArmorPenalty int    `json:"armorPenalty"`
	ArmorCost    int    `json:"armorCost"`
}
