package main

import (
	"./core/app"
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
	InitializeControllers()
	app.InitWebContext()
	InitializeRoutes()
	InitializeModels()
	app.Run()
}
