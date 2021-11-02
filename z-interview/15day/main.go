package main

import "fmt"

func main() {
	i := 65
	fmt.Println(string(i)) //A。UTF-8 编码中，十进制数字 65 对应的符号是 A。
	fmt.Println("----------")
	test()
}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func test() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA()) //13
	fmt.Println(b.ShowB()) //23
}

//接口。一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23。
