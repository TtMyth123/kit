package funs

import (
	"fmt"
)

// Exists Checks if a function has been implemented
func Exists(name string) bool {
	if _, ok := a1boolmap[name]; ok {
		return true
	} else if _, ok := a1float64map[name]; ok {
		return true
	} else if _, ok := a2boolmap[name]; ok {
		return true
	} else if _, ok := a1int64map[name]; ok {
		return true
	} else if _, ok := a2int64map[name]; ok {
		return true
	} else if _, ok := a3boolmap[name]; ok {
		return true
	} else if _, ok := a3int64map[name]; ok {
		return true
	} else if _, ok := a2float64map[name]; ok {
		return true
	} else if _, ok := a3float64map[name]; ok {
		return true
	} else if _, ok := a2inter[name]; ok {
		return true
	} else if _, ok := a4inter[name]; ok {
		return true
	} else if _, ok := a1StringMap[name]; ok {
		return true
	} else if _, ok := a2StringMap[name]; ok {
		return true
	} else if _, ok := a3StringMap[name]; ok {
		return true
	} else if _, ok := a4StringMap[name]; ok {
		return true
	} else if _, ok := anStringMap[name]; ok {
		return true
	} else {
		return false
	}
}

// Call1 Invoke arity-1 functions
func Call1(name string, input interface{}) (ret interface{}, err error) {
	if fn, ok := a1boolmap[name]; ok {
		return fn(input), nil
	} else if fn, ok := a1float64map[name]; ok {
		return fn(input)
	} else if fn, ok := a1StringMap[name]; ok {
		return fn(input)
	} else if fn, ok := a1int64map[name]; ok {
		return fn(input)
	}

	err = fmt.Errorf("1函数未实现:%s", name)
	return
}

// Call2 Invoke arity-2 functions
func Call2(name string, input1 interface{}, input2 interface{}) (ret interface{}, err error) {
	if fn, ok := a2boolmap[name]; ok {
		return fn(input1, input2), nil
	} else if fn, ok := a2float64map[name]; ok {
		return fn(input1, input2)
	} else if fn, ok := a2int64map[name]; ok {
		return fn(input1, input2)
	} else if fn, ok := a2inter[name]; ok {
		return fn(input1, input2), nil
	} else if fn, ok := a2StringMap[name]; ok {
		return fn(input1, input2)
	}

	err = fmt.Errorf("2函数未实现:%s", name)
	return
}

// Call2 Invoke arity-2 functions
func Call3(name string, input1 interface{}, input2 interface{}, input3 interface{}) (ret interface{}, err error) {
	if fn, ok := a3boolmap[name]; ok {
		return fn(input1, input2, input3), nil
	} else if fn, ok := a3int64map[name]; ok {
		return fn(input1, input2, input3)
	} else if fn, ok := a3StringMap[name]; ok {
		return fn(input1, input2, input3)
	} else if fn, ok := a3float64map[name]; ok {
		return fn(input1, input2, input3)
	}

	err = fmt.Errorf("3函数未实现:%s", name)
	return
}

func Call4(name string, input1 interface{}, input2 interface{}, input3 interface{}, input4 interface{}) (ret interface{}, err error) {
	if fn, ok := a4inter[name]; ok {
		return fn(input1, input2, input3, input4), nil
	} else if fn, ok := a4StringMap[name]; ok {
		return fn(input1, input2, input3, input4)
	}

	err = fmt.Errorf("4函数未实现:%s", name)
	return
}
