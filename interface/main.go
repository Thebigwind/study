package main

import "fmt"

func main() {
	// 3.最后，在 main 函数里调用 sayHello() 函数：
	qcrao := Student{age: 18}
	whatJob(&qcrao)

	growUp(&qcrao)
	fmt.Println(qcrao)

	stefno := Programmer{age: 25}
	whatJob(stefno)

	growUp(stefno)
	fmt.Println(stefno)
}

//1.先定义一个接口，和使用此接口作为参数的函数：
type Person interface {
	job()
	growUp()
}

func whatJob(p Person) {
	p.job()
}

func growUp(p Person) {
	p.growUp()
}

//2.再来定义两个结构体：
type Student struct {
	age int
}

func (p Student) job() {
	fmt.Println("I am a student.")
	return
}

func (p *Student) growUp() {
	p.age += 1
	return
}

type Programmer struct {
	age int
}

func (p Programmer) job() {
	fmt.Println("I am a programmer.")
	return
}

func (p Programmer) growUp() {
	// 程序员老得太快 ^_^
	p.age += 10
	return
}

// 3.最后，在 main 函数里调用 sayHello() 函数：
func test0() {

	qcrao := Student{age: 18}
	whatJob(&qcrao)

	growUp(&qcrao)
	fmt.Println(qcrao)

	stefno := Programmer{age: 25}
	whatJob(stefno)

	growUp(stefno)
	fmt.Println(stefno)
}

/*
代码里先定义了 1 个 Person 接口，包含两个函数：
job()
growUp()

然后，又定义了 2 个结构体，Student 和 Programmer，同时，类型 *Student、Programmer 实现了 Person 接口定义的两个函数。
注意，*Student 类型实现了接口， Student 类型却没有。

之后，我又定义了函数参数是 Person 接口的两个函数：
func whatJob(p Person)
func growUp(p Person)

main 函数里先生成 Student 和 Programmer 的对象，再将它们分别传入到函数 whatJob 和 growUp。
函数中，直接调用接口函数，实际执行的时候是看最终传入的实体类型是什么，调用的是实体类型实现的函数。
于是，不同对象针对同一消息就有多种表现，多态就实现了。


*/
