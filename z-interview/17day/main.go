package main

import "fmt"

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func main() {
	fmt.Println(increaseA()) //0
	fmt.Println(increaseB()) //1
}

//increaseA() 的返回参数是匿名，increaseB() 是具名。关于 defer 与返回值的知识点
