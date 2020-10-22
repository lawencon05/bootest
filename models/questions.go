package models

type Questions struct {
	BaseModels
	Question        string `json:"question" gorm:"column:qstn"`
	QuestionTypesId string
	QuestionTypes   QuestionTypes `gorm:"foreignKey:QuestionTypesId"`
}
