package controller

import (
	"github.com/labstack/echo"

	"lawencon.com/bootest/model"
	"lawencon.com/bootest/service"
)

var answerService service.AnswerService = service.AnswerServiceImpl{}

func SetAnswer(c *echo.Group) {
	c.POST("/answer", createAnswer)
}

func createAnswer(c echo.Context) (e error) {
	defer catchError(&e)
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
