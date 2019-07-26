package skillservice

import (
	"encoding/json"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/router"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//SkillHandle Rest handle for skills
type SkillHandle struct {
	SkillCollection mongo.SkillCollection
}

//GetSkills fetches all skills
func (s *SkillHandle) GetSkills(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	skills, err := s.SkillCollection.GetAllSkills()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the skills"))
		logrus.Error(err)
		return
	}
	logrus.Infof("Found %d skills", len(skills))
	json.NewEncoder(writer).Encode(skills)
}
