package model

type PojoHelper struct {
	AnswerHdr AnswerHdr   `json:"answerHdr"`
	AnswerDtl []AnswerDtl `json:"answerDtl"`
}
