package main

import (
	"lawencon.com/bootest/config"
	"lawencon.com/bootest/controller"
)

func main() {
	//set echo framework
	e := config.SetWeb()

	//set db
	g := config.SetDb()

	//set jwt group
	eg := config.SetJwt(e)

	//set controllers
	controller.SetInit(e, eg, g)

	//start server
	e.Logger.Fatal(e.Start(":1234"))
}