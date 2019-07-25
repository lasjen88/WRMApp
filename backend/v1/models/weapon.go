package models

//Weapon stores a weapon
type Weapon struct {
	WeaponName     string         `json:"weaponName"`
	ItemIdentifyer ItemIdentifyer `json:"itemidentifyer"`
	WeaponSkill    Skill          `json:"weaponSkill"`
	WeaponDamage   Damage         `json:"weaponDamage"`
	WeaponRange    int            `json:"weaponRange"`
	WeaponCost     int            `json:"weaponCost"`
	AmmoType       AmmoType       `json:"ammoType"`
}

//AmmoType stores the ammo types, with bundled sizes and costs
type AmmoType struct {
	AmmoName   string `json:"ammoName"`
	BundleSize int    `json:"bundleSize"`
	BundleCost int    `json:"bundleCost"`
}

//Bolt static ammo type used for bolts
var Bolt = AmmoType{
	AmmoName:   "Bolt",
	BundleSize: 10,
	BundleCost: 2,
}

//Arrow static ammo type used for arrows
var Arrow = AmmoType{
	AmmoName:   "Arrow",
	BundleSize: 10,
	BundleCost: 2,
}

//ThownDagger static ammo type used for daggers
var ThownDagger = AmmoType{
	AmmoName:   "Thown Dagger",
	BundleSize: 1,
	BundleCost: 2,
}

//DragonShot static ammo type used for dragon ammo
var DragonShot = AmmoType{
	AmmoName:   "Dragon Shot",
	BundleSize: 10,
	BundleCost: 4,
}

//ThownSpear static ammo type used for spears
var ThownSpear = AmmoType{
	AmmoName:   "Thown Spear",
	BundleSize: 1,
	BundleCost: 3,
}

//ThownStar static ammo type used for throwing stars
var ThownStar = AmmoType{
	AmmoName:   "Thown Star",
	BundleSize: 1,
	BundleCost: 2,
}

//NoAmmo static ammo type used to parse a non-ammo type
var NoAmmo = AmmoType{
	AmmoName:   "",
	BundleSize: 0,
	BundleCost: 0,
}

//WeaponSlot stores a weapon with meta data
type WeaponSlot struct {
	Weapon      Weapon `json:"weapon"`
	AttackBonus int    `json:"attackBonus"`
	Ammo        int    `json:"ammo"`
}
