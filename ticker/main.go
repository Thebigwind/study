package ticker

import (
	"log"
	"time"
)

//1. 简单定时任务
// 每隔1s记录一次日志：
// TickerDemo 用于演示ticker基础用法
func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Ticker tick.")
	}
}

//for range ticker.C会持续从管道中获取事件，收到事件后打印一行日志，
//如果管道中没有数据会阻塞等待事件，由于ticker会周期性的向管道中写入事件，所以上述程序会周期性的打印日志

// 2. 定时聚合任务
// TickerLaunch用于演示ticker聚合任务用法
func TickerLaunch() {
	ticker := time.NewTicker(5 * time.Minute)
	maxPassenger := 30 // 每车最大装载人数
	passengers := make([]string, 0, maxPassenger)

	for {
		passenger := GetNewPassenger() // 获取一个新乘客
		if passenger != "" {
			passengers = append(passengers, passenger)
		} else {
			time.Sleep(1 * time.Second)
		}

		select {
		case <-ticker.C: // 时间到，发车
			Launch(passengers)
			passengers = []string{}
		default:
			if len(passengers) >= maxPassenger { // 时间没到，车已座满，发车
				Launch(passengers)
				passengers = []string{}
			}
		}
	}
}

func GetNewPassenger() string {
	return ""
}

func Launch(passengers []string) {

}

// 3.
