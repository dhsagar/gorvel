package main

import (
	"./core/app"
)

func main() {
	InitializeControllers()
	InitializeRoutes()
	app.Main()
}
