package main

import (
	"./core/app"
	"./core/config"
)

/*
	|--------------------------------------------------------------------------|
	| Main -- Application Starts                                               |
	|--------------------------------------------------------------------------|
	|                                                                          |
	| Main Application Starts from Here. There is no need to Change or         |
	| Modify this File.                                                        |
	|--------------------------------------------------------------------------|
*/
func main() {
	config.ReadAllConfig()
	InitializeControllers()
	app.InitWebContext()
	InitializeRoutes()
	InitializeModels()
	app.Run()
}
