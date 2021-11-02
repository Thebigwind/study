//脉搏器（ticker）
//我们可以使用尝试发送操作来实现一个每隔一定时间发送一个信号的脉搏器。

package main

import "fmt"
import "time"

func Tick(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1) // 容量最好为1
	go func() {
		for {
			time.Sleep(d)
			select {
			case c <- struct{}{}:
			default:
			}
		}
	}()
	return c
}

func main() {
	t := time.Now()
	for range Tick(time.Second) {
		fmt.Println(time.Since(t))
	}
}

// 事实上，time标准库包中的Tick函数提供了同样的功能，但效率更高。 我们应该尽量使用标准库包中的实现。
