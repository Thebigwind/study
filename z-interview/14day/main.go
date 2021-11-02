package main

import "fmt"

func main() {
	str := "hello"
	//str[0] = 'x'  //compilation error  常量，Go 语言中的字符串是只读的。
	fmt.Println(str)

	fmt.Println("-----------")
	//可变函数调用
	add(1, 2)
	add(1, 3, 7)
	add([]int{1, 3, 7}...)
}

func add(args ...int) int {

	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}
