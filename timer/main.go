package timer

import (
	"log"
	"time"
)

//从一个管道中读取数据，在管道中没有数据时，我们不想让程序永远阻塞在管道中，而是设定一个超时时间，在此时间段中如果管道中还是没有数据到来，则判定为超时。
//WaitChannel作用就是检测指定的管道中是否有数据到来，通过select语句轮询conn和timer.C两个管道，timer会在1s后向timer.C写入数据，
//如果1s内conn还没有数据，则会判断为超时。
func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)

	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C: // 超时
		println("WaitChannel timeout!")
		return false
	}
}

//延迟执行某个方法
func DelayFunction() {
	timer := time.NewTimer(5 * time.Second)

	select {
	case <-timer.C:
		log.Println("Delayed 5s, start to do something.")
	}
}
