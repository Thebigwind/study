package main

import (
	"encoding/binary"
	"sync"
)
import "crypto/rand"

/*
数据生成/搜集/加载
一个数据产生者可能通过以下途径生成数据：
加载一个文件、或者读取一个数据库、或者用爬虫抓取网页数据；
从一个软件或者硬件系统搜集各种数据；
产生一系列随机数；
等等。
这里，我们使用一个随机数产生器做为一个数据产生者的例子。
此数据产生者函数没有输入，只有输出。
*/

func RandomGenerator() <-chan uint64 {
	c := make(chan uint64)
	go func() {
		rnds := make([]byte, 8)
		for {
			_, err := rand.Read(rnds)
			if err != nil {
				close(c)
				break
			}
			c <- binary.BigEndian.Uint64(rnds)
		}
	}()
	return c
}

//数据聚合
//一个数据聚合模块的工作协程将多个数据流合为一个数据流。 假设数据类型为int64，下面这个函数将任意数量的数据流合为一个。
func Aggregator(inputs ...<-chan uint64) <-chan uint64 {
	out := make(chan uint64)
	for _, in := range inputs {
		go func(in <-chan uint64) {
			for {
				out <- <-in // <=> out <- (<-in)
			}
		}(in)
	}
	return out
}

func Aggregator2(inputs ...<-chan uint64) <-chan uint64 {
	output := make(chan uint64)
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func(int <-chan uint64) {
			defer wg.Done()
			// 如果通道in被关闭，此循环将最终结束。
			for x := range in {
				output <- x
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

//数据分流
//数据分流是数据聚合的逆过程。数据分流的实现很简单，但在实践中用的并不多。

func Divisor(input <-chan uint64, outputs ...chan<- uint64) {
	for _, out := range outputs {
		go func(o chan<- uint64) {
			for {
				o <- <-input // <=> o <- (<-input)
			}
		}(out)
	}
}
