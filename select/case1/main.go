package main

import (
	"fmt"
	"time"
)

func main() {
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
		//不断地在调用 time.After，从而导致计时器 time.NerTimer 的不断创建和内存申请。 for在循环时，就会调用都 select 语句，因此在每次进行 select 时，都会重新初始化一个全新的计时器（Timer）。在 3 分钟后才会被触发去执行某些事，但重点在于计时器激活后，却又发现和 select 之间没有引用关系了，因此很合理的也就被 GC 给清理掉了。被抛弃的 time.After 的定时任务还是在时间堆中等待触发，在定时任务未到期之前，是不会被 GC 清除的。
		case <-time.After(3 * time.Minute):
			fmt.Printf("现在是：%d，我脑子进煎鱼了！", time.Now().Unix())
		}
	}
}
