package main

import "fmt"

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Printf("Fibonacci: %d\n", f())
	}
}

//Fibonacci()函数中原本属于局部变量的a和b由于闭包的引用，不得不将二者放到堆上，以致产生逃逸

/*
bogon:case4 luxuefeng$ go build -gcflags=-m
# study/mem-leak/case4
./check.go:7:9: can inline Fibonacci.func1
./check.go:17:13: inlining call to fmt.Printf
./check.go:6:2: moved to heap: a
./check.go:6:5: moved to heap: b
./check.go:7:9: func literal escapes to heap
./check.go:17:34: f() escapes to heap
./check.go:17:13: []interface {} literal does not escape
<autogenerated>:1: .this does not escape

*/
