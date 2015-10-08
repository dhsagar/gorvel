package router

import (
	"fmt"
	"./../function"
)

func Get(urlPattern string, handler string) {
	fmt.Println(urlPattern, handler)
	functions.Invoke(handler)
}
