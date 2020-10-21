package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"

	"lawencon.com/imamfarisi/configs"
	"lawencon.com/imamfarisi/models"
	"lawencon.com/imamfarisi/services"
)

var userService services.UserService = services.UserServiceImpl{}

func SetUser(c *echo.Group, e *echo.Echo) {
	e.POST("/user", createUser)
	e.POST("/api/login", login)
	c.GET("/user/:id", getUserById)
	//c.GET("/user", getUserById)
}

func createUser(c echo.Context) error {
	data := new(models.Users)

	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := userService.CreateUser(data)
	if err == nil {
		return res(c, result)
	}

	return resErr(c, err)
}

func getUserById(c echo.Context) error {
	id := c.Param("id")
	//id := c.QueryParam("id")

	var result, err = userService.GetUserById(id)
	if err == nil {
		if result.Count != 0 {
			return res(c, result)
		}

		return resErr(c, errors.New("Data not found"))
	}

	return resErr(c, err)
}

func login(c echo.Context) error {
	data := new(models.Users)

	if err := c.Bind(data); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var result, err = userService.Login(data.Email, data.Pwd)
	if err == nil {
		v, _ := configs.CreateJwtToken(data.Email)
		result.Token = v
		return res(c, result)	
	}

	return resErr(c, err)
}
