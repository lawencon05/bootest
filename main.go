package main

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/controller"
	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/service"
)

func main() {
	e := echo.New()

	//init db and inject to dao and
	g := initDb()
	dao.SetDao(g)
	service.SetService(g)

	//set jwt
	jwtGroup := config.SetJwt(e)

	//set controllers
	controller.SetInit(e)
	controller.SetUser(jwtGroup, e)
	controller.SetAnswer(jwtGroup)
	controller.SetEmployeeProfile(jwtGroup)
	controller.SetCandidatProfile(jwtGroup)

	//start server
	e.Logger.Fatal(e.Start(":1234"))
}

func initDb() *gorm.DB {
	g, err := config.Conn()
	if err == nil {
		config.MigrateSchema(g)
		return g
	}
	panic(err)
}
