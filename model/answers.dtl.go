package model

type AnswerDtl struct {
	BaseModel
	AnswerHdrId string    `gorm:"type:varchar(50)"`
	AnswerHdr   AnswerHdr `gorm:"foreignKey:AnswerHdrId"`
	QuestionId  string    `json:"questionId" gorm:"column:qstn_id;type:varchar(50)"`
	Questions   Questions `gorm:"foreignKey:QuestionId"`
	Answers     string    `json:"answers"`
}

func (AnswerDtl) TableName() string {
	return "tb_r_dtl_answers"
}
