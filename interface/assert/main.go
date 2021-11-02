package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func InterfaceToString(arg interface{}) string {
	switch arg := arg.(type) {
	case int64:
		return strconv.FormatInt(arg, 10)
	case string:
		return arg
	case bool:
		return strconv.FormatBool(arg)
	default:
		data, _ := json.Marshal(arg)
		return string(data)
	}
}

/*
t := i.(T)

	// T：表示的是具体的数据类型，i：表示的是接口变量i，t：表示的是转换之后的变量。

	这个语句实现的功能是：将接口变量i按照类型T转换成t，其中t中的值是i转换得来的，一旦转换不成功就会触发一个panic。

	除此之外还有另外一个写法：

	t, ok := i.(T)

	// 表达式里面多了一个ok，实现的功能是：将接口i按照类型T转换成t，如果类型匹配ok=true，如果类型不匹配，ok=false。 改语句不会触发panic。
*/

func main() {
	var a string = "abc"
	var i interface{}
	i = a
	res, ok := i.(int)
	if ok {
		fmt.Printf("int:%d", res)
	} else {
		fmt.Printf("not int")
	}

}
