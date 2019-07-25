package resources

import (
	"github.com/lasjen88/WRMApp/backend/v1/models"
)

//ReadTalentsFromFile fetches the talens from the resource file
func ReadTalentsFromFile(path string) ([]models.Talent, error) {
	var reader ResourceFileReader
	reader.SetReaderByPath(path)
	lines, err := reader.GetLines()
	if err != nil {
		return nil, err
	}
	var talents []models.Talent
	for _, line := range lines {
		talents = append(talents, models.Talent{
			TalentName:        line[0],
			TalentDescription: line[1],
		})
	}
	return talents, nil
}
