package main

import "fmt"

// 1.先定义一个接口，和使用此接口作为参数的函数：
type IGreeting interface {
	sayHello()
}

func sayHello(i IGreeting) {
	i.sayHello()
}

// 2.再来定义两个结构体：
type Go struct{}

func (g Go) sayHello() {
	fmt.Println("Hi, I am GO!")
}

type PHP struct{}

func (p PHP) sayHello() {
	fmt.Println("Hi, I am PHP!")
}

// 3.最后，在 test 函数里调用 sayHello() 函数：
func test() {
	golang := Go{}
	php := PHP{}

	sayHello(golang)
	sayHello(php)
}
