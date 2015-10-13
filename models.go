package main

import (
	"./core/models"
	"./model"
)

func InitializeModels() {

	/*
	|--------------------------------------------------------------------------
	| Models Registration For Create Table
	|--------------------------------------------------------------------------
	|
	| Here is where you need to register all the models of your app
	| that will create table in the database.
	|--------------------------------------------------------------------------
	*/

	models.Register(model.Hello{})
}
