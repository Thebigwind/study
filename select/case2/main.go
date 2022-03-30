package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Minute)
	defer timer.Stop()

	ch := make(chan int, 10)
	go func() {
		in := 1
		for {
			in++
			ch <- in
		}
	}()

	for {
		select {
		case _ = <-ch:
			// do something...
			continue
		case <-timer.C:
			fmt.Printf("现在是：%d，我脑子进煎鱼了！", time.Now().Unix())
		}
	}
}
