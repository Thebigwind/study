package mian

import (
	"fmt"
	"time"
)

/*
定时通知（timer）
用通道实现一个一次性的定时通知器是很简单的。 下面是一个自定义实现：
*/

func AfterDuration(d time.Duration) <-chan struct{} {
	c := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		c <- struct{}{}
	}()
	return c
}

func main() {
	fmt.Println("Hi!")
	<-AfterDuration(time.Second)
	fmt.Println("Hello!")
	<-AfterDuration(time.Second)
	fmt.Println("Bye!")
}

/*
事实上，time标准库包中的After函数提供了和上例中AfterDuration同样的功能。 在实践中，我们应该尽量使用time.After函数以使代码看上去更干净。

注意，操作<-time.After(aDuration)将使当前协程进入阻塞状态，而一个time.Sleep(aDuration)函数调用不会如此。

<-time.After(aDuration)经常被使用在后面将要介绍的超时机制实现中。
*/
