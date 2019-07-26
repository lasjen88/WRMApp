package resources

import (
	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//ReadSkillsFromFile fetches the skills from the resource file
func ReadSkillsFromFile(path string) ([]models.Skill, error) {
	var reader ResourceFileReader
	reader.SetReaderByPath(path)
	lines, err := reader.GetLines()
	if err != nil {
		return nil, err
	}
	var skills []models.Skill
	for _, line := range lines {
		skills = append(skills, models.Skill{
			SkillName:        line[0],
			Attribute:        models.Value(line[1]),
			SkillDescription: line[2],
		})
	}
	return skills, nil
}
