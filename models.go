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
	| that will create table in the database. If You Do not register Your
	| Model you can use it as a struct table will not be created
	| automatically in Database. Your Models will be automatically modify it
	| self if you add an column. But Will Not delete existing to
	| prevent data loss.
	|--------------------------------------------------------------------------
	*/

	models.Register(model.Hello{})
}
