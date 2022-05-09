package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("----------test1------")
	//
	//test1()
	//fmt.Println("----------test2------")
	//test2()
	//fmt.Println("----------test3------")
	//test3()
	//fmt.Println("----------test4------")
	//test4()
	//fmt.Println("----------test5------")
	//test5()
	//test0()
	test7()
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
			fmt.Println("already closed")
		}
	}
}
func test2() {
	ch1 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 20; i++ {
		a := <-ch1

		fmt.Println(a)

	}
}

func test3() {
	ch1 := make(chan int, 10)
	go func() {

		close(ch1)
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 20; i++ {
		a := <-ch1

		fmt.Println(a)

	}
}

func test4() {
	ch1 := make(chan int, 10)

	time.Sleep(1 * time.Second)
	for i := 0; i < 20; i++ {
		a := <-ch1

		fmt.Println(a)

	}
}
func test5() {
	ch1 := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			ch1 <- i
		}
	}()
	for {
		select {
		case a := <-ch1:
			fmt.Println("get a:%d", a)
		case <-time.After(1 * time.Second):
			fmt.Println("do no thing")
		}
	}
}

////
var ch1 = make(chan int, 5)

func test0() {
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(2 * time.Second)
			ch1 <- i
		}
	}()
	tick := time.Tick(1 * time.Second)
	for {
		select {
		case a := <-ch1:
			fmt.Println("get a:%d", a)
		case <-tick:
			fmt.Println("do no thing")
		}
	}
}

func test6() {
	ch1 := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch1 <- i
		}
		//close(ch1)
	}()
	for {
		select {
		case a := <-ch1:
			fmt.Println("get a:%d", a)
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

func test10() {
	fmt.Println("oo")
	select {}
	fmt.Println("xx")
}
