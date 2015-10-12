package response

import "github.com/Unknwon/macaron"

func Show(ctx *macaron.Context, templateName string, status ...int) {
	var s int
	if len(status) > 0 {
		s = status[0]
	} else {
		s = 200
	}
	ctx.HTML(s, templateName)
}

func Data(ctx *macaron.Context, name string, data interface{}) {
	ctx.Data[name] = data
}

func File(ctx *macaron.Context, file string) {
	ctx.ServeFile(file)
}

func Json(ctx *macaron.Context, value interface{}, status ...int) {
	var s int
	if len(status) > 0 {
		s = status[0]
	} else {
		s = 200
	}
	ctx.JSON(s, value)
}

func Xml(ctx *macaron.Context, value interface{}, status ...int) {
	var s int
	if len(status) > 0 {
		s = status[0]
	} else {
		s = 200
	}
	ctx.XML(s, value)
}

func Redirect(ctx *macaron.Context, location string, status ...int) {
	ctx.Redirect(location, status...)
}
