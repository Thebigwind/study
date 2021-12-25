package main

import (
	"fmt"
	"time"
)

func main() {
	NewStatPvMap()
	start := time.Now().UnixNano()
	Command()
	end := time.Now().UnixNano()
	fmt.Printf("耗时：%d ms\n", (end-start)/1e6)
}
