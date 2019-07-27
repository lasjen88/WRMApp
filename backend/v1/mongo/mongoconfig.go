package mongo

import (
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

const (
	//DatabaseURL host name of the database server
	DatabaseURL = "localhost"
	//DatabaseName name of the mongo database
	DatabaseName = "wrm"
	//ItemCollectionName collection for equipment
	ItemCollectionName = "equipment"
	//SpellCollectionName collection for spells
	SpellCollectionName = "spell"
	//CharacterCollectionName collection for characters
	CharacterCollectionName = "character"
	//SkillCollectionName collection for skills
	SkillCollectionName = "skill"
	//TalentCollectionName collection for talents
	TalentCollectionName = "talent"
	//ArmorCollectionName collection for armors
	ArmorCollectionName = "armor"
	//LanguageCollectionName collection for languages
	LanguageCollectionName = "language"
	//RaceCollectionName collection for races
	RaceCollectionName = "race"
	//ShieldCollectionName collection for shields
	ShieldCollectionName = "shield"
	//WeaponCollectionName collection for weapons
	WeaponCollectionName = "weapon"
)

//GetSession creates and returns an mgo session to the MongoDb at the specified URL.
func GetSession(URL string) *mgo.Session {
	session, err := mgo.Dial(URL)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Infof("MongoDb connected at %v", URL)
	}
	return session
}
