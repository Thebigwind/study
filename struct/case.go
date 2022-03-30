package main

import "fmt"

type People struct{}

func main() {
	test1()
	test2()
}

func test1() {
	a := &People{}      //分配到栈上
	b := &People{}      //分配到栈上
	fmt.Println(a == b) //分配到栈上。在 Go 编译器的代码优化阶段，会对其进行优化，直接返回 false
}

func test2() {
	a := &People{}
	b := &People{}
	fmt.Printf("%p\n", a) //调用了 fmt.Println 方法，该方法内部是涉及到大量的反射相关方法的调用，会造成逃逸行为，也就是分配到堆上。
	fmt.Printf("%p\n", b) //调用了 fmt.Println 方法，该方法内部是涉及到大量的反射相关方法的调用，会造成逃逸行为，也就是分配到堆上。
	fmt.Println(a == b)   //true  就是空（0字节）的在进行了逃逸分析后，往堆分配的都会指向 zerobase 这一个地址
}
