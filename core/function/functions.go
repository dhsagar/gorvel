package functions

import (
	"reflect"
	"fmt"
	"gopkg.in/macaron.v1"
)

var funcs = make(map[string]interface{})

func Register(name string, function func(*macaron.Context)) {
	registerFunction(name, function)
}

func registerFunction(name string, function func(*macaron.Context)) {
	if funcs[name] != nil {
		fmt.Println("Function name " + name + " already exists...")
		return
	}
	funcs[name] = function
}

func GetControllerFunctions(name string) (interface{}) {
	return funcs[name]
}

func Invoke(name string, params ... interface{}) (result []reflect.Value, err error) {
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
