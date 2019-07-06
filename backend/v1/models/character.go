package models

//Character stores a character for wrm
type Character struct {
	ID              string       `json:"id"`
	PlayerName      string       `json:"playername"`
	Race            Race         `json:"race"`
	Appearance      string       `json:"appearance"`
	Warrior         int          `json:"warrior"`
	Rogue           int          `json:"rogue"`
	Mage            int          `json:"mage"`
	AdventuresTaken int          `json:"adventuresTaken"`
	CurrentHp       int          `json:"currentHp"`
	CurrentFate     int          `json:"currentFate"`
	CurrentMana     int          `json:"currentMana"`
	MaxHp           int          `json:"maxHp"`
	MaxMana         int          `json:"maxMana"`
	BaseDefence     int          `json:"baseDefence"`
	ArmorTotal      int          `json:"armorTotal"`
	ArmorPanalty    int          `json:"armorPenalty"`
	Wealth          int          `json:"wealth"`
	Background      string       `json:"background"`
	WeaponSlots     []WeaponSlot `json:"weaponSlots"`
	Armors          []Armor      `json:"armors"`
	Items           []Item       `json:"items"`
	Talents         []Talent     `json:"talents"`
	Skills          []Skill      `json:"skills"`
	Spells          []Spell      `json:"spells"`
	Langueages      []Language   `json:"Language"`
}
