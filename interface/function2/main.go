package main

import "fmt"

/*
	实现了接收者是值类型的方法，相当于自动实现了接收者是指针类型的方法；
    而实现了接收者是指针类型的方法，不会自动生成对应接收者是值类型的方法。

*/
type coder interface {
	code()
	debug()
}

type Gopher struct {
	language string
}

func (p Gopher) code() {
	fmt.Printf("I am coding %s language\n", p.language)
}

func (p *Gopher) debug() {
	fmt.Printf("I am debuging %s language\n", p.language)
}

func main() {
	var c coder = &Gopher{"Go"}
	c.code()
	c.debug()

	//var c2 coder = Gopher{"Go"}  编译错误
	//Gopher 没有实现 coder，因为 Gopher 类型并没有实现 debug 方法；
	// 表面上看， *Gopher 类型也没有实现 code 方法，但是因为 Gopher 类型实现了 code 方法，所以让 *Gopher 类型自动拥有了 code 方法。
	// 如果实现了接收者是值类型的方法，会隐含地也实现了接收者是指针类型的方法。

	// 如果方法的接收者是值类型，无论调用者是对象还是对象指针，修改的都是对象的副本，不影响调用者；
	// 如果方法的接收者是指针类型，则调用者修改的是指针指向的对象本身。对接收者的属性进行更改操作，从而影响接收者

	// 如果类型具备“原始的本质”，也就是说它的成员都是由 Go 语言里内置的原始类型，如字符串，整型值等，那就定义值接收者类型的方法。
	// 像内置的引用类型，如 slice，map，interface，channel，这些类型比较特殊，声明他们的时候，实际上是创建了一个 header，
	// 对于他们也是直接定义值接收者类型的方法。这样，调用函数时，是直接 copy 了这些类型的 header，而 header 本身就是为复制设计的。

	// 如果类型具备非原始的本质，不能被安全地复制，这种类型总是应该被共享，那就定义指针接收者的方法。
	// 比如 go 源码里的文件结构体（struct File）就不应该被复制，应该只有一份实体。
}
