package main

import "fmt"

// 1.
func main() {
	s1 := make([]int, 5)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)

	fmt.Println("-----------------------")
	s2 := make([]int, 0)
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2)
}

/*
[0 0 0 0 0 1 2 3]
-----------------------
[1 2 3 4]

*/

/*
new() 与 make() 的区别

new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。

new(T)，返回一个指针，即类型为 *T的值,该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。
make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.

*/
