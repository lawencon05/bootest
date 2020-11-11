package controller

import (
	"errors"

	"github.com/labstack/echo"

	"lawencon.com/bootest/model"
	"lawencon.com/bootest/service"
)

var userService service.UserService = service.UserServiceImpl{}

func SetUser(c *echo.Group, e *echo.Echo) {
	e.POST("/user", createUser)
	e.POST("/api/login", login)
	c.GET("/user/:id", getUserById)
	//c.GET("/user", getUserById)
}

func createUser(c echo.Context) (e error) {
	defer catchError(&e)
	data := new(model.Users)

	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}

	result, err := userService.CreateUser(data)
	if err == nil {
		return res(c, result)
	}

	return resErr(c, err)
}

func getUserById(c echo.Context) (e error) {
	defer catchError(&e)
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

func login(c echo.Context) (e error) {
	defer catchError(&e)
	data := new(model.Users)
	if err := c.Bind(data); err != nil {
		return resErr(c, err)
	}
	var result, err = userService.Login(data.Email, data.Pwd)
	if err == nil {
		return res(c, result)
	}
	return resErr(c, err)
}
