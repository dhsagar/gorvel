package functions

import "reflect"

var funcs = make(map[string]interface{})

func Register(name string, function func()) {
	registerFunction(name, function)
}

func registerFunction(name string, function func()) {
	if funcs[name] != nil {
		panic("functions name already exixts")
	}
	funcs[name] = function
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
