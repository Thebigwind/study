package main

import (
	"fmt"
)

func main() {

	c := []int{10, 19, 24, 61, 5, 121, 9, 11, 34, 21, 22}
	fmt.Println(c)
	//fmt.Println(qsort(c))

	testa()
}

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivot := a[right]
	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < pivot { //如果小于分区点,需要交换位置，放在左边；否则，不需要交换，继续遍历
			a[i], a[left] = a[left], a[i] //把当前遍历的index位置数据(小于分区点)和left位置的数据交换
			fmt.Printf("----------------\n")
			fmt.Printf("left:%v，value：%v\n", left, a[left])
			fmt.Printf("index:%v,%v\n", i, a)

			left++
			//fmt.Printf("left:%v，value：%v\n",left,a[left])
		} else {
			fmt.Printf("----------------\n")
			fmt.Printf("left:%v，value：%v\n", left, a[left])
			fmt.Printf("index:%v,%v\n", i, a)
		}
	}

	// 一轮比较完后，将分区点放到分区点的合适位置上
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	//qsort(a[:left])
	//qsort(a[left+1:])

	return a
}

func testa() {
	a := make(chan int, 1)
	//a <- 1
	//close(a)

	fmt.Println(IsClosed(a))
}
func IsClosed(c chan int) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}
