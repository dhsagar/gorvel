package app
import (
	"gopkg.in/macaron.v1"
	"./../function"
	_ "./../../controller"
)

var m *macaron.Macaron

func Run() {
	m.Run(8080)
}

func InitWebContext() {
	m = macaron.Classic()
	rendererOptions := macaron.RenderOptions{
		Directory: "views",
		Delims: macaron.Delims{"{{", "}}"},
	}
	m.Use(macaron.Renderer(rendererOptions))
	m.Use(macaron.Static("resources"))
	m.NotFound(routerNotFoundHandler)
	m.InternalServerError(serverErrorHandler)
	m.Action(errorsHandler)
}

func AddRouteContext(method string, pattern string, handler string) {
	controller := functions.GetControllerFunctions(handler)
	cHandler := make([]macaron.Handler, 1)
	cHandler[0] = controller
	m.Handle(method, pattern, cHandler)
}

func routerNotFoundHandler(ctx *macaron.Context) {
	ctx.HTML(404, "404")
}

func serverErrorHandler(ctx *macaron.Context) {
	ctx.HTML(500, "500")
}

func errorsHandler(ctx *macaron.Context) {
	status := ctx.Resp.Status()
	if status == 404 {
		ctx.HTML(404, "404")
	} else if status >= 400 {
		ctx.HTML(500, "500")
	}
}
