package main

import "fmt"

func main() {
	//var r rune = 'A'
	// a 值是 48
	var a int = int('A')
	fmt.Println(a)

	i := 47
	// 转换成 rune 字符类型，但是打印出来发现依然是数字样式
	var r1 rune = rune(i)
	// 真正可以输出字符
	var str string = string(r1)
	fmt.Println(str)

	itoa64 := "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	fmt.Println(65 & 0x3f)
	fmt.Println(itoa64[65&0x3f])

}
