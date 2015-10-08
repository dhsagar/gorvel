package main

import (
	"./router"
)

func AllRoutes() {

	/*
	|--------------------------------------------------------------------------
	| Application Routes
	|--------------------------------------------------------------------------
	|
	| Here is where you can register all of the routes for an application.
	| It's a breeze. Simply tell gorvel the URIs it should respond to
	| and give it the controller to call when that URI is requested.
	|
	*/

	router.Get("/index", "Hello")
}
