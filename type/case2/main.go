package main

import "fmt"

func main() {
	type MyInt int
	type IntPtr *int
	type MyIntPtr *MyInt

	var pi = new(int)  // pi的类型为*int
	var ip IntPtr = pi // 没问题，因为底层类型相同
	// 并且pi的类型为非定义类型。

	// var _ *MyInt = pi // 不能隐式转换
	var _ = (*MyInt)(pi) // 显式转换是没问题的

	// 类型*int的值不能被直接转换为类型MyIntPtr，
	// 但是可以间接地转换过去。
	/*
	   var _ MyIntPtr = pi  // 不能隐式转换
	   var _ = MyIntPtr(pi) // 也不能显式转换
	*/
	var _ MyIntPtr = (*MyInt)(pi)  // 间接隐式转换没问题
	var _ = MyIntPtr((*MyInt)(pi)) // 间接显式转换没问题

	// 类型IntPtr的值不能被直接转换为类型MyIntPtr，
	// 但是可以间接地转换过去。
	/*
	   var _ MyIntPtr = ip  // 不能隐式转换
	   var _ = MyIntPtr(ip) // 也不能显式转换
	*/
	// 间接隐式或者显式转换都是没问题的。
	var _ MyIntPtr = (*MyInt)((*int)(ip))  // ok
	var _ = MyIntPtr((*MyInt)((*int)(ip))) // ok

	fmt.Println("------------------------")

	var a int
	a = 4

	var ax MyInt
	ax = MyInt(a)
	fmt.Println(ax)

}
