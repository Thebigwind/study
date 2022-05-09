package main

import (
	"fmt"
	"sync"
)

/*
事实上，上例中展示的多对单和单对多通知实现方式在实践中用的并不多。 在实践中，我们多使用sync.WaitGroup来实现多对单通知
，使用关闭一个通道的方式来实现单对多通知（详见下一个用例）。

通过关闭一个通道来实现群发通知
上一个用例中的单对多通知实现在实践中很少用，因为通过关闭一个通道的方式在来实现单对多通知的方式更简单。
我们已经知道，从一个已关闭的通道可以接收到无穷个值，我们可以利用这一特性来实现群发通知。

我们可以把上一个例子中的三个数据发送操作ready <- struct{}{}替换为一个通道关闭操作close(ready)来达到同样的单对多通知效果。

从一个已关闭的通道可以接收到无穷个值这一特性也将被用在很多其它在后面将要介绍的用例中。
实际上，这一特性被广泛地使用于标准库包中。比如，context标准库包使用了此特性来传达操作取消消息。
*/

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(index int) {
			defer func() {
				if e := recover(); e != nil {
					fmt.Printf("recover index:%v\n", index)
				}
			}()
			defer wg.Done()

			panic("crash")
			fmt.Printf("index:%v\n", index)
		}(i)

	}
	wg.Wait()
	fmt.Printf("finish")
}

/*

SCHED：调试信息输出标志字符串，代表本行是goroutine调度器的输出；
0ms：即从程序启动到输出这行日志的时间；
gomaxprocs: P的数量，本例有2个P, 因为默认的P的属性是和cpu核心数量默认一致，当然也可以通过GOMAXPROCS来设置；
idleprocs: 处于idle状态的P的数量；通过gomaxprocs和idleprocs的差值，我们就可知道执行go代码的P的数量；
threads: os threads/M的数量，包含scheduler使用的m数量，加上runtime自用的类似sysmon这样的thread的数量；
spinningthreads: 处于自旋状态的os thread数量；
idlethread: 处于idle状态的os thread的数量；
runqueue=0： Scheduler全局队列中G的数量；
[0 0]: 分别为2个P的local queue中的G的数量。

me@localhostt study % GODEBUG=schedtrace=1000 go run chan/case7/check.go
SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=6 spinningthreads=1 idlethreads=0 runqueue=0 [1 0 0 0 0 0 0 0]
# command-line-arguments
SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=5 spinningthreads=1 idlethreads=0 runqueue=0 [1 0 0 0 0 0 0 0]
# command-line-arguments
SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=5 spinningthreads=1 idlethreads=1 runqueue=0 [0 0 0 0 0 0 0 0]
SCHED 1008ms: gomaxprocs=8 idleprocs=8 threads=17 spinningthreads=0 idlethreads=10 runqueue=0 [0 0 0 0 0 0 0 0]
SCHED 0ms: gomaxprocs=8 idleprocs=5 threads=5 spinningthreads=1 idlethreads=1 runqueue=0 [1 0 0 0 0 0 0 0]
recover index:9
recover index:3
recover index:7
recover index:8
recover index:4
recover index:2
recover index:6
recover index:1
finishrecover index:5
recover index:0
me@localhostt study %

*/
