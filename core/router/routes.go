package router

import (
	"./../app"
)

func Get(urlPattern string, handler string) {
	app.AddRouteContext("GET", urlPattern, handler)
}

func Post(urlPattern string, handler string) {
	app.AddRouteContext("POST", urlPattern, handler)
}

func Put(urlPattern string, handler string) {
	app.AddRouteContext("PUT", urlPattern, handler)
}

func Patch(urlPattern string, handler string) {
	app.AddRouteContext("PATCH", urlPattern, handler)
}

func Delete(urlPattern string, handler string) {
	app.AddRouteContext("DELETE", urlPattern, handler)
}

func Options(urlPattern string, handler string) {
	app.AddRouteContext("OPTIONS", urlPattern, handler)
}

func Head(urlPattern string, handler string) {
	app.AddRouteContext("HEAD", urlPattern, handler)
}

func Any(urlPattern string, handler string) {
	app.AddRouteContext("*", urlPattern, handler)
}

