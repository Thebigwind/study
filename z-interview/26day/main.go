package main

import "fmt"

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

const (
	a = iota
	b = iota
)
const (
	name = "name"
	c    = iota
	d    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	/*
	   0
	   1
	   1
	   2

	   iota是golang语言的常量计数器,只能在常量的表达式中使用。

	   iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。

	   使用iota能简化定义，在定义枚举时很有用。

	*/

	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}

/*
s is nil
p is not nil

当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。
上面的代码，给变量 p 赋值之后，p 的动态值是 nil，但是动态类型却是 *Student，是一个 nil 指针，所以相等条件不成立
*/
