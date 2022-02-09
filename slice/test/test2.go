package main

import "fmt"

//import  OS   "os"
//import "strings"
//import "path/filepath"

type Stack []interface{}

func (s *Stack) f() {
	stack := *s

	fmt.Printf("%p %d %d\n", &s, len(stack), cap(stack))

	*s = stack[:len(stack)-1] //减少

	fmt.Printf("%p %d %d\n", &s, len(*s), cap(*s))

	fmt.Printf("%p %d %d\n", stack, len(stack), cap(stack))

	stack = append(stack, "e", "f") //增加

	fmt.Printf("%p %d %d\n", stack, len(stack), cap(stack))
}

func main() {
	var s Stack
	s = append(s, "a", "b", "c", "d")
	s.f()
	fmt.Println(s)
	fmt.Printf("%t\n", false)
}
