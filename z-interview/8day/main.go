package main

import "fmt"

func hello() []string {
	return nil
}

func main() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

//not nil
//是将 hello() 赋值给变量 h，而不是函数的返回值，所以输出 not nil。

func GetValue() int {
	return 1
}

//cannot type switch on non-interface value i (type int)
//类型选择，类型选择的语法形如：i.(type)，其中 i 是接口，type 是固定关键字
/*
func test() {
	i := GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}

*/
