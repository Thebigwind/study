package main

import "fmt"

func main() {
	type MyInt int

	var i int = 1

	var j MyInt = MyInt(i)

	fmt.Println("%T", j)
}

/*
多重赋值分为两个步骤，有先后顺序：

计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；

赋值；

所以本例，会先计算 s[i-1]，等号右边是两个表达式是常量，所以赋值运算等同于 i, s[0] = 2, “Z”。

*/
