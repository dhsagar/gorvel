package models

import "./../database"

func Register(model interface{}) {
	database.RegisterModel(model)
}
