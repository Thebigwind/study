package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
将单向发送通道类型用做函数实参
和上例一样，在下面这个例子中，sumSquares函数调用的两个实参的请求也是并发进行的。
和上例不同的是longTimeRequest函数接收一个单向发送通道类型参数而不是返回一个单向接收通道结果。
*/
func longTimeRequest(r chan<- int32) {
	time.Sleep(time.Second * 3) // 模拟一个工作负载
	r <- rand.Int31n(100)
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ra, rb := make(chan int32), make(chan int32)
	go longTimeRequest(ra)
	go longTimeRequest(rb)

	fmt.Println(sumSquares(<-ra, <-rb))
}
