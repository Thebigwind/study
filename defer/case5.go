package main

import "fmt"

func main() {
	test1()
	test2()
	test3()
	fmt.Println("------")
	println(f1())
	println(f2())
	println(f3())
}

func test1() {
	var whatever [6]struct{}
	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}
}

/*
5
5
5
5
5
5
*/

func test2() {
	var whatever [6]struct{}
	for i := range whatever {
		a := i
		defer func() {
			fmt.Println(a)
		}()
	}
}

/*
5
4
3
2
1
0
*/

func test3() {
	var whatever [6]struct{}
	for i := range whatever {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
}

/*
5
4
3
2
1
0
*/

/*
其根本原因是闭包所导致的，有两点原因：

在 for 循环结束后，局部变量 i 的值已经是 5 了，并且 defer 的闭包是直接引用变量的 i。
结合defer 关键字的特性，可得知会在 main 方法主体结束后再执行。

*/

//1
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

//5
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

//1
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
