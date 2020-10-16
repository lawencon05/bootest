package controllers

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/labstack/echo"

// 	"lawencon.com/imamfarisi/models"
// )

// func SetAnswer(c *echo.Echo) {
// 	c.POST("/answer", createAnswer)
// }

// func createAnswer(c echo.Context) error {
// 	data := new(models.Users)

// 	if err := c.Bind(data); err != nil {
// 		return c.String(http.StatusBadRequest, "something wrong.. ")
// 	}

// 	result, _ := json.Marshal(data)
// 	return c.String(http.StatusOK, string(result))
// }
