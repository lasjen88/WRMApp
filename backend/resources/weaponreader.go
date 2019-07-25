package resources

import (
	"errors"
	"strconv"
	"strings"

	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//WeaponDto is a dto for the weapon resource
type WeaponDto struct {
	weapon      models.Weapon
	skillString string
}

//ReadWeaponsFromFile fetches the weapons from the resource file
func ReadWeaponsFromFile(path string) ([]WeaponDto, error) {
	var reader ResourceFileReader
	reader.SetReaderByPath(path)
	lines, err := reader.GetLines()
	if err != nil {
		return nil, err
	}
	var weapons []WeaponDto
	for _, line := range lines {
		var valueError error
		weapons, valueError = parseWeapon(line, weapons)
		if valueError != nil {
			return nil, valueError
		}
	}
	return weapons, nil
}

func parseWeapon(line []string, weapons []WeaponDto) ([]WeaponDto, error) {
	var weapon WeaponDto
	var err error
	weapon.weapon, err = createModelWeapon(line)
	if err != nil {
		return nil, err
	}
	weapon.skillString = strings.TrimSpace(line[1])
	weapons = append(weapons, weapon)
	return weapons, nil
}

func createModelWeapon(line []string) (models.Weapon, error) {
	var modelWeapon models.Weapon
	modelWeapon.WeaponName = line[0]
	damage, damageError := parseDamage(strings.TrimSpace(line[2]))
	if damageError != nil {
		return modelWeapon, damageError
	}
	weaponRange, rangeError := parseRange(strings.TrimSpace(line[3]))
	if rangeError != nil {
		return modelWeapon, rangeError
	}
	cost, costErr := strconv.Atoi(line[4])
	if costErr != nil {
		return modelWeapon, costErr
	}
	ammo, ammoError := parseAmmo(strings.TrimSpace(line[5]))
	if ammoError != nil {
		return modelWeapon, ammoError
	}
	modelWeapon.ItemIdentifyer = models.ItemIdentifyer{Source: line[6]}
	modelWeapon.WeaponDamage = damage
	if weaponRange > 0 {
		modelWeapon.WeaponRange = weaponRange
	}
	modelWeapon.WeaponCost = cost
	if ammo != models.NoAmmo {
		modelWeapon.AmmoType = ammo
	}
	return modelWeapon, nil
}

func parseDamage(damageString string) (models.Damage, error) {
	damage := models.Damage{
		DSixDiceCount:    0,
		ModifierOperator: models.None,
		ModifierValue:    0,
	}
	dSix, dSixError := strconv.Atoi(damageString[:1])
	if dSixError != nil {
		return damage, dSixError
	}
	damage.DSixDiceCount = dSix
	if len(damageString) > 3 {
		if damageString[3:4] == "-" {
			damage.ModifierOperator = models.Substraction
		} else if damageString[3:4] == "+" {
			damage.ModifierOperator = models.Addition
		} else if damageString[3:4] == "/" {
			damage.ModifierOperator = models.Divisor
		} else {
			return damage, errors.New("Bad modifier")
		}
		var modiferValueError error
		damage.ModifierValue, modiferValueError = strconv.Atoi(damageString[5:6])
		if modiferValueError != nil {
			return damage, modiferValueError
		}
	}
	return damage, nil
}

func parseRange(rangeString string) (int, error) {
	if rangeString == "" {
		return 0, nil
	}
	return strconv.Atoi(rangeString)
}

func parseAmmo(ammoString string) (models.AmmoType, error) {
	if ammoString == "" {
		return models.NoAmmo, nil
	}
	return models.ThownStar, nil
}
