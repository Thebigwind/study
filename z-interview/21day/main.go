package main

import "fmt"

func main() {
	var a []int  //  声明的是 nil 切片,声明不会分配内存，优先选择
	b := []int{} //声明的是长度和容量都为 0 的空切片
	fmt.Println(a)
	fmt.Println(b)
}
