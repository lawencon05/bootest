package models

type AnswerDtl struct {
	BaseModels
	AnswerHdrId string
	AnswerHdr   AnswerHdr `gorm:"foreignKey:AnswerHdrId"`
	QstnId      string    `json:"questionId" `
	Answers     string    `json:"answers"`
}

func (AnswerDtl) TableName() string {
	return "tb_r_dtl_answers"
}
