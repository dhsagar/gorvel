package app
import "github.com/Unknwon/macaron"

func Main() {
	m := macaron.Classic()
	rendererOptions := macaron.RenderOptions{
		// Directory to load templates. Default is "templates".
		Directory: "./../../view",
		// Extensions to parse template files from. Defaults are [".tmpl", ".html"].
		Extensions: []string{".tmpl", ".html"},
		Delims: macaron.Delims{"{{", "}}"},
	}
	m.Use(macaron.Renderer(rendererOptions))
	m.Use(macaron.Static("./../resource"))
	AddAllRoutes(m)
	m.Run(8080)
}
