package controllers

import (
	"net/http"

	"github.com/labstack/echo"

	"lawencon.com/imamfarisi/models"
	"lawencon.com/imamfarisi/services"
)

var answerService services.AnswerService = services.AnswerServiceImpl{}

func SetAnswer(c *echo.Group) {
	c.POST("/answer", createAnswer)
}

func createAnswer(c echo.Context) error {
	data := new(models.PojoHelper)

	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var err = answerService.CreateAnswer(&data.AnswerHdr, &data.AnswerDtl)
	if err == nil {
		return res(c, data)
	}
	return resErr(c, err)
}
