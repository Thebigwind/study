package main

import (
	"fmt"
	"time"
)

func main() {
	//test1()
	//test7()
	//test8()
	//test11()
	testss()
}

func test1() {
	ch1 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 20; i++ {
		a, ok := <-ch1
		if ok {
			fmt.Println(a)
		} else {
			fmt.Println(a)
			fmt.Println("already closed")
		}
	}
}

var ch1 = make(chan int, 5)

func test0() {
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			ch1 <- i
		}
	}()
	tick := time.Tick(1 * time.Second)
	//tick := time.NewTicker(1 * time.Second)
	//defer tick.Stop()
	for {
		select {
		case a := <-ch1:
			fmt.Println("get a:%d", a)
		case <-tick:
			fmt.Println("do no thing")
		}
	}
}

func test2() {
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			ch1 <- i
		}
	}()
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()
	for {
		select {
		case a := <-ch1:
			fmt.Println("get a:%d", a)
		case <-tick.C:
			fmt.Println("do no thing")
		}
	}
}

func test7() {
	ch1 := make(chan int, 10)
	fmt.Println(IsClosed(ch1))

	close(ch1)
	fmt.Println(IsClosed(ch1))
}

func IsClosed(c chan int) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}
func test8() {
	select {
	default:
	}
	fmt.Println("xxx")
}

func test11() {

	aa := []int{0, 1, 2, 66, 77}
	//a := []int{0, 1, 2}
	a := aa[:3]
	s := make([]int, 3)
	copy(s, a)
	fmt.Println(a, s) // [0 1 2] [0 1 2]
	s[0] = 11

	fmt.Println(a, s) // [0 1 2] [11 1 2]
	s = append(s, 12)
	s = append(s, 13)
	fmt.Println(a, s) // [0 1 2] [11 1 2 12 13]
	s[0] = 21

	fmt.Println(a, s) // [0 1 2] [21 1 2 12 13]

}

func testss() {
	a := 'a'
	fmt.Println(int(a))
	fmt.Println(string(int(a)))
	fmt.Println(string(a))
}
