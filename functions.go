package main

import (
	"./core/function"
	"./controller"
)

func InitializeControllers() {

	/*
	|--------------------------------------------------------------------------
	| Controller Function Registration
	|--------------------------------------------------------------------------
	|
	| Here is where you need to register all the controller functions
	| with a unique name. When You reference this in your router
	| with that unique name you go this method as a response. You can Use
	| Your Custom package here. You just need to import it on the import
	| section.
	*/

	functions.Register("Hello", controller.Hello)
}
