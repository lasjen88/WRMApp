package models

//Race stores a race
type Race struct {
	RaceName        string     `json:"raceName"`
	RaceDescription string     `json:"raceDescription"`
	RaceLanguages   []Language `json:"raceLanguages"`
	RaceTalents     []Talent   `json:"raceTalents"`
}
