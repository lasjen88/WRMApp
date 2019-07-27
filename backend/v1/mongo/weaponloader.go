package mongo

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lasjen88/WRMApp/backend/resources"
	"github.com/lasjen88/WRMApp/backend/v1/models"
	"github.com/sirupsen/logrus"
)

var weaponFiles = [...]string{"./backend/resources/weapons.csv"}

//InitializeWeapons preloads the database with weapons from the resource files
func InitializeWeapons(weaponCollection WeaponCollection, skillCollection SkillCollection) error {
	weaponsIsEmpty, err := weaponCollection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if !weaponsIsEmpty {
		logrus.Info("Weapons already loaded.")
		return nil
	}
	skillsIsEmpty, DatabaseError := skillCollection.IsEmptyCollection()
	if DatabaseError != nil {
		return DatabaseError
	}
	if skillsIsEmpty {
		return errors.New("Cannot load weapons before skills have been loaded")
	}
	for _, path := range weaponFiles {
		err = loadWeapon(weaponCollection, skillCollection, path)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadWeapon(collection WeaponCollection, skillCollection SkillCollection, path string) error {
	weapons, err := resources.ReadWeaponsFromFile(path)
	if err != nil {
		return err
	}
	logrus.Infof("Found %d weapons in %s.", len(weapons), path)
	skills, skillError := skillCollection.GetAllSkills()
	if skillError != nil {
		return skillError
	}
	for _, weaponDto := range weapons {
		weapon, dtoError := getWeaponFromDto(weaponDto, skills)
		if dtoError != nil {
			return dtoError
		}
		databaseError := collection.PutWeapon(weapon)
		if databaseError != nil {
			return databaseError
		}
	}
	logrus.Info("Weapon load complete")
	return nil
}

func getWeaponFromDto(weaponDto resources.WeaponDto, skills []models.Skill) (models.Weapon, error) {
	weapon := weaponDto.Weapon
	for _, skill := range skills {
		if strings.EqualFold(weaponDto.SkillString, skill.SkillName) {
			weapon.WeaponSkill = skill
			return weapon, nil
		}
	}
	errorMessage := fmt.Sprintf("WeaponDto skill string '%s' does not match any skills in the database", weaponDto.SkillString)
	return weapon, errors.New(errorMessage)
}
