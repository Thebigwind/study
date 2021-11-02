package mian

import (
	"log"
	"math/rand"
	"time"
)

/*
将通道用做计数信号量（counting semaphore）
缓冲通道可以被用做计数信号量。 计数信号量可以被视为多主锁。如果一个缓冲通道的容量为N，那么它可以被看作是一个在任何时刻最多可有N个主人的锁。
上面提到的二元信号量是特殊的计数信号量，每个二元信号量在任一时刻最多只能有一个主人。

计数信号量经常被使用于限制最大并发数。

和将通道用做互斥锁一样，也有两种方式用来获取一个用做计数信号量的通道的一份所有权。

通过发送操作来获取所有权，通过接收操作来释放所有权；
通过接收操作来获取所有权，通过发送操作来释放所有权。
下面是一个通过接收操作来获取所有权的例子：
*/
type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {
	log.Print("顾客#", c, "进入酒吧")
	seat := <-bar // 需要一个位子来喝酒
	log.Print("++ customer#", c, " drinks at seat#", seat)
	log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
	bar <- seat // 释放座位，离开酒吧
}

func main() {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10) // 此酒吧有10个座位
	// 摆放10个座位。
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId) // 均不会阻塞
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		go bar24x7.ServeCustomer(customerId)
	}
	for {
		time.Sleep(time.Second)
	} // 睡眠不属于阻塞状态
}
