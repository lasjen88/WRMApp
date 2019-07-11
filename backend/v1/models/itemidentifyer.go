package models

//ItemIdentifyer stores the origin of an item by providing it with a unique id and the rule set it came from
type ItemIdentifyer struct {
	ID     string `json:"id"`
	Source string `json:"source"`
}
