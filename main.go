package main

import (
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"lawencon.com/imamfarisi/configs"
	"lawencon.com/imamfarisi/controllers"
	"lawencon.com/imamfarisi/dao"
	"lawencon.com/imamfarisi/services"
)

func main() {
	e := echo.New()

	//get connection to db and set to dao
	g := newConn()
	dao.SetDao(g)
	services.SetService(g)

	//set jwt
	jwtGroup := configs.SetJwt(e)

	//set controllers
	controllers.SetInit(e)
	controllers.SetUser(jwtGroup, e)
	controllers.SetAnswer(jwtGroup)
	controllers.SetEmployeeProfile(jwtGroup)
	controllers.SetCandidatProfile(jwtGroup)

	//start server
	e.Logger.Fatal(e.Start(":1234"))
}

func newConn() *gorm.DB {
	g, err := configs.Conn()
	if err == nil {
		configs.MigrateSchema(g)
		return g
	}
	panic(err)
}
