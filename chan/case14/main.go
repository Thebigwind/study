package main

import "fmt"

/*
尝试发送和尝试接收
含有一个default分支和一个case分支的select代码块可以被用做一个尝试发送或者尝试接收操作，取决于case关键字后跟随的是一个发送操作还是一个接收操作。

如果case关键字后跟随的是一个发送操作，则此select代码块为一个尝试发送操作。
如果case分支的发送操作是阻塞的，则default分支将被执行，发送失败；否则发送成功，case分支得到执行。

如果case关键字后跟随的是一个接收操作，则此select代码块为一个尝试接收操作。
如果case分支的接收操作是阻塞的，则default分支将被执行，接收失败；否则接收成功，case分支得到执行。

尝试发送和尝试接收代码块永不阻塞。

标准编译器对尝试发送和尝试接收代码块做了特别的优化，使得它们的执行效率比多case分支的普通select代码块执行效率高得多。

下例演示了尝试发送和尝试接收代码块的工作原理。
*/

func main() {
	type Book struct{ id int }
	bookshelf := make(chan Book, 3)

	for i := 0; i < cap(bookshelf)*2; i++ {
		select {
		case bookshelf <- Book{id: i}:
			fmt.Println("成功将书放在书架上", i)
		default:
			fmt.Println("书架已经被占满了")
		}
	}

	for i := 0; i < cap(bookshelf)*2; i++ {
		select {
		case book := <-bookshelf:
			fmt.Println("成功从书架上取下一本书", book.id)
		default:
			fmt.Println("书架上已经没有书了")
		}
	}
}
