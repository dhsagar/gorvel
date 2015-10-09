package view

import "github.com/Unknwon/macaron"

func Show(ctx *macaron.Context, name string) {
	ctx.HTML(200, name)
}

func Data(ctx *macaron.Context, name string, data interface{}) {
	ctx.Data[name] = data
}
