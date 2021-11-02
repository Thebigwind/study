package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		fmt.Printf("val:%v\n", &val)
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}

	fmt.Println("--------------")
	test2()
}

/*
for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，
所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.
val:0xc000018050
val:0xc000018050
val:0xc000018050
val:0xc000018050
2 -> 3
3 -> 3
0 -> 3
1 -> 3
--------------
value:0xc000018070
value:0xc000018078
value:0xc000018080
value:0xc000018088
0 ===> 0
1 ===> 1
2 ===> 2
3 ===> 3

*/

func test2() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		value := val
		m[key] = &value
		fmt.Printf("value:%v\n", &value)

	}

	for k, v := range m {
		fmt.Println(k, "===>", *v)
	}
}
