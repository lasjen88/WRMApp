package wrmrules

const (
	spellCostPerCircle = 25
)

//GetDifficultyLevel provides a discrete difficulty level for circles in [1:4]
func GetDifficultyLevel(circle int) int {
	switch circle {
	case 1:
		return 5
	case 2:
		return 7
	case 3:
		return 9
	case 4:
		return 13
	default:
		return 0
	}
}

//GetManaConsumption provided a discrete mana consumption value for circles in [1:4] under the exponential function manacost = 0.5 e^(0.693147 * x)
func GetManaConsumption(circle int) int {
	switch circle {
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	case 4:
		return 8
	default:
		return 0
	}
}

//GetCost provides the cost of a spell under the function cost = circle * 25
func GetCost(circle int) int {
	return circle * spellCostPerCircle
}
