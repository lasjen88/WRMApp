package wrmrulestest

import (
	"testing"

	"github.com/lasjen88/WRMApp/wrmrules"
)

func TestFallingInMeters(t *testing.T) {
	fallInMeters := 6
	expectedMinimumDamage := 2
	expectedMaximumDamage := 12

	damage := wrmrules.GetDamageFromFallenMeters(fallInMeters)
	if damage < expectedMinimumDamage || damage > expectedMaximumDamage {
		t.Errorf("Expected damage between %d and %d, but was %d.", expectedMinimumDamage, expectedMaximumDamage, damage)
	}
}

func TestFallingInDivisableYards(t *testing.T) {
	fallInYards := 6
	expectedMinimumDamage := 2
	expectedMaximumDamage := 12

	damage := wrmrules.GetDamageFromFallenYards(fallInYards)
	if damage < expectedMinimumDamage || damage > expectedMaximumDamage {
		t.Errorf("Expected damage between %d and %d, but was %d.", expectedMinimumDamage, expectedMaximumDamage, damage)
	}
}

func TestFallingWithoutDamage(t *testing.T) {
	fallInYards := 2
	damage := wrmrules.GetDamageFromFallenYards(fallInYards)
	if damage != 0 {
		t.Errorf("Did not expect damage, but took %d.", damage)
	}
}
