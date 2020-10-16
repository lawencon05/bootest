package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"

	"lawencon.com/imamfarisi/models"
	"lawencon.com/imamfarisi/services"
)

var userService services.UserService = services.UserServiceImpl{}

func SetUser(c *echo.Echo) {
	c.POST("/user", createUser)
}

func createUser(c echo.Context) error {
	data := new(models.Users)

	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	userService.CreateUser(data)

	result, _ := json.Marshal(data)
	return c.String(http.StatusOK, string(result))
}
