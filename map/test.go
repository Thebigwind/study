package main

import (
	"fmt"
	"sync"
)

var counter = struct {
	sync.RWMutex
	m map[string]int
}{m: make(map[string]int)}

//这条语句声明了一个变量，它是一个匿名结构（struct）体，包含一个原生和一个嵌入读写锁 sync.RWMutex。

func main() {
	counter.RLock()
	n := counter.m["煎鱼"]
	counter.RUnlock()
	fmt.Println("煎鱼:", n)
}
