package main

/*
 goroutien 泄露的场景

面试官为啥会问 Goroutine（协程）泄露这种奇特的问题呢？

可以猜测是：
Goroutine 实在是使用门槛实在是太低了，随手就一个就能起，出现了不少滥用的情况。例如：并发 map。
Goroutine 本身在 Go 语言的标准库、复合类型、底层源码中应用广泛。例如：HTTP Server 对每一个请求的处理就是一个协程去运行。

泄露的原因大多集中在：
Goroutine 内正在进行 channel/mutex 等读写操作，但由于逻辑问题，某些情况下会被一直阻塞。
Goroutine 内的业务逻辑进入死循环，资源一直无法释放。
Goroutine 内的业务逻辑进入长时间等待，有不断新增的 Goroutine 进入等待。
*/
