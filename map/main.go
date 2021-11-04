package main

import (
	"fmt"
	"math"
)

/*
Go 语言中读取 map 有两种语法：带 comma 和 不带 comma。当要查询的 key 不在 map 里，带 comma 的用法会返回一个 bool 型变量提示 key 是否在 map 中；
而不带 comma 的语句则会返回一个 key 类型的零值。如果 key 是 int 型就会返回 0，如果 key 是 string 类型，就会返回空字符串。
*/
func main() {
	//ageMap := make(map[string]int)
	//ageMap["qcrao"] = 18
	//
	//// 不带 comma 用法
	//age1 := ageMap["stefno"]
	//fmt.Println(age1)
	//
	//// 带 comma 用法
	//age2, ok := ageMap["stefno"]
	//fmt.Println(age2, ok)

	//test()

	test2()
}

func test() {
	m := make(map[float64]int)
	m[1.4] = 1
	m[2.4] = 2
	m[math.NaN()] = 3
	m[math.NaN()] = 3

	for k, v := range m {
		fmt.Printf("[%v, %d] ", k, v)
	}

	fmt.Printf("\nk: %v, v: %d\n", math.NaN(), m[math.NaN()])
	fmt.Printf("k: %v, v: %d\n", 2.400000000001, m[2.400000000001])
	fmt.Printf("k: %v, v: %d\n", 2.4000000000000000000000001, m[2.4000000000000000000000001])

	fmt.Println(math.NaN() == math.NaN())
}

func test2() {
	aa := make(map[int][]string)

	value, ok := aa[4]
	if ok {
		fmt.Println("xx")
		fmt.Println(value)
	} else {
		fmt.Println("oo")
		fmt.Println(value)
	}

	//aaV := make([]string,0)
	aa[4] = append(aa[4], "3")

	if ok {
		fmt.Println("xxxx")
		fmt.Println(aa[4])
	} else {
		fmt.Println("oooo")
		fmt.Println(aa[4])
	}
}
