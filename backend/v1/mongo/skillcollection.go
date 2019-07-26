package mongo

import (
	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//SkillCollection holds the data collection
type SkillCollection struct {
	Wrmcollection WrmCollection
}

//GetAllSkills fetches all skills in the collection
func (s *SkillCollection) GetAllSkills() ([]models.Skill, error) {
	if s.Wrmcollection.collection == nil {
		s.Wrmcollection.SetupCollection()
	}
	var skills []models.Skill
	err := s.Wrmcollection.collection.Find(nil).All(&skills)
	if err != nil {
		return nil, err
	}
	return skills, nil
}

//PutSkill inserts a skill to the collection
func (s *SkillCollection) PutSkill(newSkill models.Skill) error {
	if s.Wrmcollection.collection == nil {
		s.Wrmcollection.SetupCollection()
	}
	err := s.Wrmcollection.collection.Insert(newSkill)
	if err != nil {
		return err
	}
	return nil
}

//IsEmptyCollection checks if the collection is currently empty
func (s *SkillCollection) IsEmptyCollection() (bool, error) {
	skills, err := s.GetAllSkills()
	if err != nil {
		return true, err
	}
	return len(skills) == 0, nil
}
