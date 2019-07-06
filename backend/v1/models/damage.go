package models

import "fmt"

type modifierOperator int

const (
	//Addition modifies the damage by adding additional damage to the d6 throw
	Addition modifierOperator = 0

	//Substraction modifies the damage by subtracting damage from the d6 throw
	Substraction modifierOperator = 1

	//Divisor modifies the damage by dividing the damage of the d6 throw
	Divisor modifierOperator = 2
)

//Damage stores a damage structure
type Damage struct {
	DSixDiceCount    int              `json:"dSixDiceCount"`
	ModifierOperator modifierOperator `json:"addition"`
	ModifierValue    int              `json:"divisor"`
}

//String provides a readable string representation of the damage calculation
func (d Damage) String() string {
	modifierString := ""
	switch d.ModifierOperator {
	case Addition:
		modifierString = combineModifierOperatorAndValue("+", d)
	case Substraction:
		modifierString = combineModifierOperatorAndValue("-", d)
	case Divisor:
		modifierString = combineModifierOperatorAndValue("/", d)
	default:
	}
	return fmt.Sprintf("%dd%s", d.DSixDiceCount, modifierString)
}

func combineModifierOperatorAndValue(operator string, d Damage) string {
	return fmt.Sprintf(" %s %d", operator, d.ModifierValue)
}
