package wrmrules

import "math/rand"

//RollDSix returns a random number betweem 1 and 6.
func RollDSix() int {
	return getDieRoll(6)
}

//RollDFour returns a random number betweem 1 and 4.
func RollDFour() int {
	return getDieRoll(4)
}

//RollDTwenty returns a random number betweem 1 and 20.
func RollDTwenty() int {
	return getDieRoll(20)
}

func getDieRoll(die int) int {
	return rand.Intn(die-1) + 1
}
