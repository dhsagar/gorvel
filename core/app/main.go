package app
import (
	"github.com/Unknwon/macaron"
	"./../function"
	_ "./../../controller"
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
	m.NotFound()
}


func AddRouteContext(pattern string, handler string) {
	controller := functions.GetControllerFunctions(handler)
	m.Get(pattern, controller)
}
