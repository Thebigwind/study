package main

import (
	"fmt"
	"strings"
)

func Test1() bool {
	var v string
	if v == "" {
		return true
	}
	return false
}

func Test2() bool {
	var v string
	if len(v) == 0 {
		return true
	}
	return false
}

/*
无论是 len(v) == 0，又或是 v == "" 的判断，其编译出来的汇编代码都是完全一致的。可以明确 Go 编译器在这块做了明确的优化，大概率是直接比对了。
*/

func main() {
	Test3()
}
func Test3() {

	str := "Go爱好者"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))

	aa := []rune(str)
	for i := len(aa) - 1; i >= 0; i-- {
		fmt.Printf("%q", aa[i])
	}
	fmt.Printf("\n")
	mid := len(aa) / 2
	for i := 0; i < mid; i++ {
		aa[i], aa[len(aa)-1-i] = aa[len(aa)-1-i], aa[i]
	}
	fmt.Printf("%q\n", aa)

}
func test3() {
	//采用strings.Join函数
	//str := strings.Join(testSlice, "")

	//strings.Builder的WriteString函数
	var testBuilder strings.Builder
	for i := 0; i < 10; i++ {
		testBuilder.WriteString("test")
	}
	fmt.Println(testBuilder.String())

	//
	reslut := fmt.Sprintf("%s%s%s", "aa", "bb", "cc")
	fmt.Println(reslut)

}
