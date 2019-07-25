package models

import "strings"

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

//Value returns the AttributeEnum of the input attributestring. Returns -1 if the value does not exist.
func Value(attribute string) AttributeEnum {
	if strings.EqualFold(attribute, Warrior.String()) {
		return Warrior
	}
	if strings.EqualFold(attribute, Rogue.String()) {
		return Rogue
	}
	if strings.EqualFold(attribute, Mage.String()) {
		return Mage
	}
	return -1
}
