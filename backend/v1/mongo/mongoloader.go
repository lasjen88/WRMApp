package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/lasjen88/WRMApp/backend/resources"
	"github.com/sirupsen/logrus"
)

var equipmentFiles = [...]string{"./backend/resources/equipment.csv"}
var spellFiles = [...]string{
	"./backend/resources/firstCircleSpells.csv",
	"./backend/resources/secondCircleSpells.csv",
	"./backend/resources/thirdCircleSpells.csv",
	"./backend/resources/fourthCircleSpells.csv",
}
var skillFiles = [...]string{"./backend/resources/skills.csv"}
var talentFiles = [...]string{"./backend/resources/talents.csv"}
var languageFiles = [...]string{"./backend/resources/languages.csv"}
var armorFiles = [...]string{"./backend/resources/armor.csv"}
var shieldFiles = [...]string{"./backend/resources/shields.csv"}
var raceFiles = [...]string{"./backend/resources/races.csv"}

//InitializeEquipment preloads the database with equipment from the resource files
func InitializeEquipment(mongoSession *mgo.Session) error {
	GetSession(DatabaseURL)
	collection := ItemCollection{DatabaseName: DatabaseName, CollectionName: ItemCollectionName, Session: mongoSession}
	isEmpty, err := collection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if isEmpty {
		for _, path := range equipmentFiles {
			err = loadEquipment(collection, path)
			if err != nil {
				return err
			}
		}
		return nil
	}
	logrus.Info("Equipment already loaded.")
	return nil
}

//InitializeSpells preloads the database with spells from the resource files
func InitializeSpells(mongoSession *mgo.Session) error {
	GetSession(DatabaseURL)
	collection := SpellCollection{DatabaseName: DatabaseName, CollectionName: SpellCollectionName, Session: mongoSession}
	isEmpty, err := collection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if isEmpty {
		for fileIndex, path := range spellFiles {
			err = loadSpells(collection, path, fileIndex+1)
			if err != nil {
				return err
			}
		}
		return nil
	}
	logrus.Info("Spells already loaded.")
	return nil
}

//InitializeSkills preloads the database with skills from the resource files
func InitializeSkills(collection SkillCollection) error {
	isEmpty, err := collection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if isEmpty {
		for _, path := range skillFiles {
			err = loadSkills(collection, path)
			if err != nil {
				return err
			}
		}
		return nil
	}
	logrus.Info("Skills already loaded.")
	return nil
}

//InitializeTalents preloads the database with talents from the resource files
func InitializeTalents(collection TalentCollection) error {
	isEmpty, err := collection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if isEmpty {
		for _, path := range talentFiles {
			err = loadTalents(collection, path)
			if err != nil {
				return err
			}
		}
		return nil
	}
	logrus.Info("Talents already loaded.")
	return nil
}

//InitializeLanguages preloads the database with languages from the resource files
func InitializeLanguages(collection LanguageCollection) error {
	isEmpty, err := collection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if isEmpty {
		for _, path := range languageFiles {
			err = loadLanguages(collection, path)
			if err != nil {
				return err
			}
		}
		return nil
	}
	logrus.Info("Languages already loaded.")
	return nil
}

//InitializeArmors preloads the database with armor from the resource files
func InitializeArmors(collection ArmorCollection) error {
	isEmpty, err := collection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if isEmpty {
		for _, path := range armorFiles {
			err = loadArmor(collection, path)
			if err != nil {
				return err
			}
		}
		return nil
	}
	logrus.Info("Armors already loaded.")
	return nil
}

//InitializeShields preloads the database with shields from the resource files
func InitializeShields(collection ShieldCollection) error {
	isEmpty, err := collection.IsEmptyCollection()
	if err != nil {
		return err
	}
	if isEmpty {
		for _, path := range shieldFiles {
			err = loadShield(collection, path)
			if err != nil {
				return err
			}
		}
		return nil
	}
	logrus.Info("Shields already loaded.")
	return nil
}

func loadShield(collection ShieldCollection, path string) error {
	shields, err := resources.ReadArmorFromFile(path)
	if err != nil {
		return err
	}
	logrus.Infof("Found %d shields in %s.", len(shields), path)
	for _, shield := range shields {
		databaseError := collection.PutShield(shield)
		if databaseError != nil {
			return databaseError
		}
	}
	logrus.Info("Shield load complete")
	return nil
}

func loadArmor(collection ArmorCollection, path string) error {
	armors, err := resources.ReadArmorFromFile(path)
	if err != nil {
		return err
	}
	logrus.Infof("Found %d armors in %s.", len(armors), path)
	for _, armor := range armors {
		databaseError := collection.PutArmor(armor)
		if databaseError != nil {
			return databaseError
		}
	}
	logrus.Info("Armor load complete")
	return nil
}

func loadLanguages(collection LanguageCollection, path string) error {
	languages, err := resources.ReadLanguageFromFile(path)
	if err != nil {
		return err
	}
	logrus.Infof("Found %d languages in %s.", len(languages), path)
	for _, language := range languages {
		databaseError := collection.PutLanguage(language)
		if databaseError != nil {
			return databaseError
		}
	}
	logrus.Info("Language load complete")
	return nil
}

func loadTalents(collection TalentCollection, path string) error {
	talents, err := resources.ReadTalentsFromFile(path)
	if err != nil {
		return err
	}
	logrus.Infof("Found %d tallents in %s.", len(talents), path)
	for _, talent := range talents {
		databaseError := collection.PutTalent(talent)
		if databaseError != nil {
			return databaseError
		}
	}
	logrus.Info("Talent load complete")
	return nil
}

func loadSkills(collection SkillCollection, path string) error {
	skills, err := resources.ReadSkillsFromFile(path)
	if err != nil {
		return err
	}
	logrus.Infof("Found %d skills in %s.", len(skills), path)
	for _, skill := range skills {
		databaseError := collection.PutSkill(skill)
		if databaseError != nil {
			return databaseError
		}
	}
	logrus.Info("Skill load complete")
	return nil
}

func loadSpells(collection SpellCollection, path string, circle int) error {
	spells, err := resources.ReadSpellsFromFile(path, circle)
	if err != nil {
		return err
	}
	logrus.Infof("Found %d spells in %s.", len(spells), path)
	collection.PutSpells(spells)
	logrus.Info("Spell load complete")
	return nil
}

func loadEquipment(collection ItemCollection, path string) error {
	items, err := resources.ReadItemsFromFile(path)
	if err != nil {
		return err
	}
	logrus.Infof("Found %d pices of equipment.", len(items))
	collection.PutItems(items)
	logrus.Info("Equipment load complete")
	return nil
}
