package main

import (
	"fmt"
	"time"
)

func main() {
	//c,cnC := context.WithCancel(context.Background())
	cha := make(chan int, 1)
	go func() {
		time.Sleep(time.Second * 10)
		cha <- 1
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Printf("超时，取消子goroutine执行")
	case <-cha:
		fmt.Printf("已找到，取消子goroutine执行")
	}
	select {}
	fmt.Printf("xxxxxx")
}
