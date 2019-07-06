package models

//Item stores an item
type Item struct {
	ItemName        string `json:"itemName"`
	ItemDescription string `json:"itemDescription"`
	ItemCost        int    `json:"itemCost"`
}
