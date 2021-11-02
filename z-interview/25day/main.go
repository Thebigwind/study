package main

import "fmt"

type Myint int

func (i Myint) PrintInt() {
	fmt.Println(i)
}

func main() {
	var i Myint = 1
	i.PrintInt()

	fmt.Println("--------------------")

	//var peo People = Student{}  编译错误 Student does not implement People (Speak method has pointer receiver)，值类型 Student 没有实现接口的 Speak() 方法，而是指针类型 *Student 实现该方法。
	var peo People = &Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}
