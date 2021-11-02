package main

import "runtime"

/*
使当前协程永久阻塞
Go中的选择机制（select）是一个非常独特的特性。它给并发编程带来了很多新的模式和技巧。

我们可以用一个无分支的select流程控制代码块使当前协程永久处于阻塞状态。 这是select流程控制的最简单的应用。 事实上，上面很多例子中的for {time.Sleep(time.Second)}都可以换为select{}。

一般，select{}用在主协程中以防止程序退出。
*/

func DoSomething() {
	for {
		// 做点什么...

		runtime.Gosched() // 防止本协程霸占CPU不放
	}
}

func main() {
	go DoSomething()
	go DoSomething()
	select {}
}
