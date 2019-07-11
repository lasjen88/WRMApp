package models

//Item stores an item
type Item struct {
	ItemIdentifyer  ItemIdentifyer `json:"itemidentifyer"`
	ItemName        string         `json:"itemName"`
	ItemDescription string         `json:"itemDescription"`
	ItemCost        int            `json:"itemCost"`
}
