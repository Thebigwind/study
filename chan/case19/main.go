package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	ch <- 18
	close(ch)
	x, ok := <-ch
	if ok {
		fmt.Println("received: ", x)
	}

	x, ok = <-ch
	if !ok {
		fmt.Println("channel closed, data invalid.")
	}

	test3()
}

func test() {
	ch := make(chan int, 3)
	ch <- 18
	//close(ch)
	x, ok := <-ch
	if ok {
		fmt.Println("received: ", x)
	}

	x, ok = <-ch
	if !ok {
		fmt.Println("channel closed, data invalid.")
	}
}

//Channel 在什么情况下必须 close?
// 对于有缓冲的channel,接收端不知道 channel 中数据量，在channel中没数据后，仍然从channel中取数据。
// 如果channel没 close,会：fatal error: all goroutines are asleep - deadlock!

func test2() {
	ch := make(chan int)
	ch <- 18

	x := <-ch

	fmt.Println("received: ", x)

	x = <-ch

	fmt.Println("channel closed, data invalid.")

}

/*
goroutine 1 [chan receive]:
main.test2()
        /Users/luxuefeng/go/study/chan/case19/read.md:49 +0x134

*/

//Channel 在什么情况下必须 close?
// 对于有缓冲的channel,接收端不知道 channel 中数据量，在channel中没数据后，仍然从channel中取数据。
// 如果channel没 close,会：fatal error: all goroutines are asleep - deadlock!

func test3() {
	ch := make(chan int)
	go func() {
		time.Sleep(4 * time.Second)
		ch <- 18
	}()

	x := <-ch

	fmt.Println("received: ", x)
}

//无缓冲channel，用于channel 通信
