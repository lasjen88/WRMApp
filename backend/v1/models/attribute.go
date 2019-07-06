package models

//AttributeEnum is the enumeration unit for the three attributes of warrior, rogue & mage
type AttributeEnum int

const (

	//Warrior is the warrior attribute.
	Warrior AttributeEnum = 0

	//Rogue is the rogue attribute.
	Rogue AttributeEnum = 1

	//Mage is the mage attribute.
	Mage AttributeEnum = 2
)

func (a AttributeEnum) String() string {
	values := [...]string{
		"Warrior",
		"Rogue",
		"Mage"}

	if a < Warrior || a > Mage {
		return ""
	}
	return values[a]
}
