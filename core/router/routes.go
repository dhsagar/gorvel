package router

import (
	"./../app"
)

func Get(urlPattern string, handler string) {
	app.AddRouteContext(urlPattern, handler)
}
