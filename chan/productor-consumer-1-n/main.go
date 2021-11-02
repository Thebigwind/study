package main

import (
	"fmt"
	"time"
)

//5 个工作协程在不断地从工作队列里取任务，生产方只管往 channel 发送任务即可，解耦生产方和消费方。
func main() {
	taskCh := make(chan int, 100)
	go worker(taskCh)

	// 塞任务
	for i := 0; i < 10; i++ {
		taskCh <- i
	}

	// 等待 1 小时
	select {
	case <-time.After(time.Hour):
	}
}

func worker(taskCh <-chan int) {
	const N = 5
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <-taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
