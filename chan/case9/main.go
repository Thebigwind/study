package mian

import (
	"fmt"
)

/*
将通道用做互斥锁（mutex）
上面的某个例子提到了容量为1的缓冲通道可以用做一次性二元信号量。
事实上，容量为1的缓冲通道也可以用做多次性二元信号量（即互斥锁）尽管这样的互斥锁效率不如sync标准库包中提供的互斥锁高效。

有两种方式将一个容量为1的缓冲通道用做互斥锁：

通过发送操作来加锁，通过接收操作来解锁；
通过接收操作来加锁，通过发送操作来解锁。
下面是一个通过发送操作来加锁的例子。
*/
func main() {
	mutex := make(chan struct{}, 1) // 容量必须为1

	counter := 0
	increase := func() {
		mutex <- struct{}{} // 加锁
		counter++
		<-mutex // 解锁
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println(counter) // 2000
}

/*
下面是一个通过接收操作来加锁的例子，其中只显示了相对于上例而修改了的部分。

...
func main() {
    mutex := make(chan struct{}, 1)
    mutex <- struct{}{} // 此行是必需的

    counter := 0
    increase := func() {
        <-mutex // 加锁
        counter++
        mutex <- struct{}{} // 解锁
    }
...
*/
