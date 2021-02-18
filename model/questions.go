package model

type Questions struct {
	BaseModel
	Question        string        `json:"question" gorm:"column:qstn"`
	QuestionTypesId string        `gorm:"type:varchar(50)"`
	QuestionTypes   QuestionTypes `gorm:"foreignKey:QuestionTypesId"`
}

func (Questions) TableName() string {
	return "tb_m_qstn"
}
