package main

import "fmt"

func main() {
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
