package wrmrules

import "math/rand"

const (
	yardsPerMeter = 1.0936
)

//GetDamageFromFallenYards takes the fall in yards and output the damage taken from the fall
func GetDamageFromFallenYards(fallInYards int) int {
	return getDamageFromFallenYards(float64(fallInYards))
}

//GetDamageFromFallenMeters takes the fall in meters and output the damage taken from the fall
func GetDamageFromFallenMeters(fallInMeters int) int {
	return getDamageFromFallenYards(float64(fallInMeters) * yardsPerMeter)
}

func getDamageFromFallenYards(fallInYards float64) int {
	damage := 0
	for i := 1.0; i <= fallInYards/3.0; i++ {
		damage = damage + rand.Intn(6)
	}
	return damage
}
