package main

import (
	"fmt"
	"sync"
)

func main() {
	var s []string
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(9999)
	for i := 0; i < 9999; i++ {
		go func() {
			defer wg.Done()
			lock.Lock()
			s = append(s, "脑子进煎鱼了")
			lock.Unlock()
		}()
	}
	wg.Wait()

	fmt.Printf("进了 %d 只煎鱼", len(s))

	test2()
}

func test2() {
	var (
		slice1 = []int{}
		n      = 1000
		c      = make(chan int, n)
	)

	for i := 0; i < n; i++ {
		go func(a int) {
			c <- a
		}(i)
	}

	for i := 0; i < n; i++ { //在这里阻塞实现 等待n个协程执行完毕
		slice1 = append(slice1, <-c)
	}

	fmt.Printf("len of slice is:%v\n", len(slice1))
	//for _, v := range slice1 {
	//	fmt.Printf("value:%v\n", v)
	//}
}
