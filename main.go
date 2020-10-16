package main

import (
	"github.com/labstack/echo"
	"lawencon.com/imamfarisi/controllers"

	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "it's works")
	})
	controllers.SetUser(e)

	e.Logger.Fatal(e.Start(":1234"))
}
