package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
)

/*
使用通道实现通知
通知可以被看作是特殊的请求/回应用例。在一个通知用例中，我们并不关心回应的值，我们只关心回应是否已发生。
所以我们常常使用空结构体类型struct{}来做为通道的元素类型，因为空结构体类型的尺寸为零，能够节省一些内存（虽然常常很少量）。

向一个通道发送一个值来实现单对单通知
我们已知道，如果一个通道中无值可接收，则此通道上的下一个接收操作将阻塞到另一个协程发送一个值到此通道为止。
所以一个协程可以向此通道发送一个值来通知另一个等待着从此通道接收数据的协程
*/

//通道done被用来做为一个信号通道来实现单对单通知
func main() {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // 也可以是缓冲的

	// 排序协程
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{} // 通知排序已完成
	}()

	// 并发地做一些其它事情...

	<-done // 等待通知
	fmt.Println(values[0], values[len(values)-1])
}
