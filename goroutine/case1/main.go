package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 9999; i++ { //9999999
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("lxf")
	say("你好")
}
