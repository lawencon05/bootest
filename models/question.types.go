package models

type QuestionTypes struct {
	BaseModels
	QuestionName string `json:"name" gorm:"column:qstn_name"`
	QuestionCode string `json:"code" gorm:"unique;column:code"`
}

func (QuestionTypes) TableName() string {
	return "tb_m_qstn_types"
}
