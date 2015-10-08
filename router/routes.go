package router

import (
	"fmt"
	"reflect"
	"./../controller"
)

var funcs = map[string]interface{}{"foo": controller.Hello, "bar": controller.HelloTwo}

func Get(urlPattern string, handler string) {
	fmt.Println(urlPattern, handler)
	Call("foo")
}


func Call(name string, params ... interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(funcs[name])
	if len(params) != f.Type().NumIn() {
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
