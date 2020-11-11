package model

type AnswerHdr struct {
	BaseModel
}

func (AnswerHdr) TableName() string {
	return "tb_r_hdr_answers"
}
