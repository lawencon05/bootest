package controller

import (
	"github.com/labstack/echo/v4"

	"lawencon.com/bootest/config"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
	"lawencon.com/bootest/service"
)

var answerService service.AnswerService

func setAnswer() {
	eg.POST("/answer", createAnswer)

	answerService = service.AnswerServiceImpl{
		DB: db,
		AnswerDao:    dao.AnswerDaoImpl{
			DB: db,
		},
		AnswerDtlDao: dao.AnswerDtlDaoImpl{
			DB: db,
		},
	}
}

func createAnswer(c echo.Context) (e error) {
	defer config.CatchError(&e)
	data := new(model.PojoHelper)

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	var err = answerService.CreateAnswer(&data.AnswerHdr, &data.AnswerDtl)
	if err == nil {
		return res(c, data)
	}
	return resErr(c, err)
}
