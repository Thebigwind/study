package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(3 * time.Second)
			ch1 <- i
		}
	}()
	for {
		select {
		case a := <-ch1:
			fmt.Printf("get a:%d\n", a)
		case <-time.After(2 * time.Second):
			fmt.Printf("do no thing\n")
		}
	}

}

/*
为了避免太多的Timer值被创建，我们应该只使用（并复用）一个Timer值，像下面这样：
*/
func longRunning(messages <-chan string) {
	timer := time.NewTimer(time.Minute)
	defer timer.Stop()

	for {
		select {
		case <-timer.C: // 过期了
			return
		case msg := <-messages:
			fmt.Println(msg)

			// 此if代码块很重要。
			if !timer.Stop() {
				<-timer.C
			}
		}

		// 必须重置以复用。
		timer.Reset(time.Minute)
	}
}
