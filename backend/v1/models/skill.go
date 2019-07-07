package models

//Skill stores a skill
type Skill struct {
	SkillName        string        `json:"skillName"`
	Attribute        AttributeEnum `json:"attribute"`
	SkillDescription string        `json:"skillDescription"`
}
