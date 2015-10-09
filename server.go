package main

import (
	"./core/app"
)

func main() {
	InitializeControllers()
	app.InitWebContext()
	InitializeRoutes()
	app.Main()
}
