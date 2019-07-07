package models

//Weapon stores a weapon
type Weapon struct {
	WeaponName   string   `json:"weaponName"`
	WeaponSkill  Skill    `json:"weaponSkill"`
	WeaponDamage Damage   `json:"weaponDamage"`
	WeaponRange  int      `json:"weaponRange"`
	WeaponCost   int      `json:"weaponCost"`
	AmmoType     AmmoType `json:"ammoType"`
}

//AmmoType stores the ammo types, with bundled sizes and costs
type AmmoType struct {
	AmmoName   string
	BundleSize int
	BundleCost int
}

//WeaponSlot stores a weapon with meta data
type WeaponSlot struct {
	Weapon      Weapon `json:"weapon"`
	AttackBonus int    `json:"attackBonus"`
	Ammo        int    `json:"ammo"`
}
