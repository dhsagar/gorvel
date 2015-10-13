package controller

import (
	"gopkg.in/macaron.v1"
	"./../core/response"
)

func Hello(ctx *macaron.Context) {
	type JSONTest struct {
		Name string
		Value int
	}

	newJSON := JSONTest{"Sadlil", 5}
	response.Json(ctx, newJSON)
}
