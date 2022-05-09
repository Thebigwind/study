package main

import "fmt"

var strChan = make(chan string, 1)

func main() {
	fmt.Println(len(strChan))
	//strChan <- 	"aaa"
	_, ok := <-strChan
	if ok {
		fmt.Printf("ok:%v\n", ok)

		//fmt.Printf("a:%v",a)
	}

}
