package main

import (
	"fmt"
	"os"
	"time"
)

//对话（或称乒乓）
//两个协程可以通过一个通道进行对话，整个过程宛如打乒乓球一样。 下面是一个这样的例子，它将打印出一系列斐波那契（Fibonacci）数。

type Ball uint64

func Play(playerName string, table chan Ball) {
	var lastValue Ball = 1
	for {
		ball := <-table // 接球
		fmt.Println(playerName, ball)
		ball += lastValue
		if ball < lastValue { // 溢出结束
			os.Exit(0)
		}
		lastValue = ball
		table <- ball // 回球
		time.Sleep(time.Second)
	}
}

func main() {
	table := make(chan Ball)
	go func() {
		table <- 1 // （裁判）发球
	}()
	go Play("A:", table)
	Play("B:", table)
}
