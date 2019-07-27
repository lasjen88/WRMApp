package languageservice

import (
	"encoding/json"
	"net/http"

	"github.com/lasjen88/WRMApp/backend/router"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/sirupsen/logrus"
)

//LanguageHandle Rest handle for languages
type LanguageHandle struct {
	LanguageCollection mongo.LanguageCollection
}

//GetLanguages fetches all languages
func (l *LanguageHandle) GetLanguages(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	languages, err := l.LanguageCollection.GetAllLanguages()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the languages"))
		logrus.Error(err)
		return
	}
	logrus.Infof("Found %d languages", len(languages))
	json.NewEncoder(writer).Encode(languages)
}
