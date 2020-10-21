package models

type AnswerHdr struct {
	BaseModels
}

func (AnswerHdr) TableName() string {
	return "tb_r_hdr_answers"
}
