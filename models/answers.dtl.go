package models

type AnswerDtl struct {
	BaseModels
	AnswerHdrId string
	AnswerHdr   AnswerHdr `gorm:"foreignKey:AnswerHdrId"`
	QuestionId  string    `json:"questionId" gorm:"column:qstn_id"`
	Answers     string    `json:"answers"`
}

func (AnswerDtl) TableName() string {
	return "tb_r_dtl_answers"
}
