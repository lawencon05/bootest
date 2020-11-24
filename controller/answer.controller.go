package controller

import (
	"github.com/labstack/echo/v4"

	"lawencon.com/bootest/config"
	"lawencon.com/bootest/model"
	"lawencon.com/bootest/service"
)

var answerService service.AnswerService = service.AnswerServiceImpl{}

func SetAnswer(eg *echo.Group) {
	eg.POST("/answer", createAnswer)
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
