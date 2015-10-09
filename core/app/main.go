package app
import (
	"github.com/Unknwon/macaron"
	"./../function"
	_ "./../../controller"
	"fmt"
)

var m *macaron.Macaron

func Main() {
	m.Run(8080)
}


func InitWebContext() {
	m = macaron.Classic()
	rendererOptions := macaron.RenderOptions{
		// Directory to load templates. Default is "templates".
		Directory: "view",
		Delims: macaron.Delims{"{{", "}}"},
	}
	m.Use(macaron.Renderer(rendererOptions))
	m.Use(macaron.Static("resource"))
	m.NotFound(routerNotFoundHandler)
	m.Action(errorHandler)
}


func AddRouteContext(pattern string, handler string) {
	controller := functions.GetControllerFunctions(handler)
	m.Get(pattern, controller)
}

func routerNotFoundHandler(ctx *macaron.Context) {
	ctx.HTML(404, "404")
}

func errorHandler(ctx *macaron.Context) {
	status := ctx.Resp.Status()
	if  status == 400 {

	} else if status == 503 {
		fmt.Println("After Request Not Found")
		ctx.HTML(503, "503")
	}
}
