package core

var funcs map[string]interface{}

func RegisterFunction(name string, function func()) {
	if funcs[name] != nil {
		panic("function name already exixts")
	}
	funcs[name] = function
}

