package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
返回单向接收通道做为函数返回结果
在下面这个例子中，sumSquares函数调用的两个实参请求并发进行。 每个通道读取操作将阻塞到请求返回结果为止。
两个实参总共需要大约3秒钟（而不是6秒钟）准备完毕（以较慢的一个为准）
*/

func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(time.Second * 3) // 模拟一个工作负载
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSquares(<-a, <-b))
}
