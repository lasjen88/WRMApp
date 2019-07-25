package itemservice

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/lasjen88/WRMApp/backend/router"
	"github.com/lasjen88/WRMApp/backend/v1/models"
	"github.com/lasjen88/WRMApp/backend/v1/mongo"
	"github.com/lasjen88/WRMApp/wrmrules"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

type spellDto struct {
	SpellName        string `json:"spellname"`
	SpellDescription string `json:"spelldescription"`
	Circle           int    `json:"circle"`
}

//SpellHandle Rest handle for spells
type SpellHandle struct {
	SpellCollection mongo.SpellCollection
}

//GetSpells provides all spells from the database, regardless of circle
func (handle *SpellHandle) GetSpells(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	spells, err := handle.SpellCollection.GetAllSpells()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Error while searching for the spells"))
		log.Error(err)
		return
	}
	logrus.Infof("Found %d spells", len(spells))
	json.NewEncoder(writer).Encode(spells)
}

//GetCircleSpells provides all spells from the specified circle
func (handle *SpellHandle) GetCircleSpells(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	params := mux.Vars(request)
	circle, err := strconv.Atoi(params["circle"])
	if err != nil {
		log.Error(err)
		return
	}
	spells, err := handle.SpellCollection.GetAllSpells()
	logrus.Infof("Found %d spells. Starting cicle filtering.", len(spells))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("400 - Unable to find spells in that circle"))
		log.Error(err)
		return
	}
	circledSpells := getSpellsFromCircle(spells, circle)
	logrus.Infof("Found %d spells at circle %d", len(circledSpells), circle)
	json.NewEncoder(writer).Encode(circledSpells)
}

//CreateSpell creates a spell
func (handle *SpellHandle) CreateSpell(writer http.ResponseWriter, request *http.Request) {
	writer = router.SetHtppWriterHeaders(writer)
	var spellDto spellDto
	err := json.NewDecoder(request.Body).Decode(&spellDto)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("400 - Unable to parse spell"))
		log.Error(err)
		return
	}
	err = handle.SpellCollection.PutSpell(createSpell(spellDto))
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("500 - Unable create spell"))
		log.Error(err)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Spell created"))
}

func getSpellsFromCircle(allSpells []models.Spell, circle int) []models.Spell {
	circledSpells := allSpells[:0]
	for _, spell := range allSpells {
		if isCircle(spell, circle) {
			circledSpells = append(circledSpells, spell)
		}
	}
	return circledSpells
}

func createSpell(dto spellDto) models.Spell {
	return models.Spell{
		SpellName:        dto.SpellName,
		DifficultyLevel:  wrmrules.GetDifficultyLevel(dto.Circle),
		ManaConsumption:  wrmrules.GetManaConsumption(dto.Circle),
		SpellCost:        wrmrules.GetCost(dto.Circle),
		SpellDescription: dto.SpellDescription,
	}
}

func isCircle(spell models.Spell, circle int) bool {
	return spell.SpellCost == wrmrules.GetCost(circle)
}
