package main

import "fmt"

//copy复制为值复制，改变原切片的值不会影响新切片。
//而等号复制为指针复制，改变原切片或新切片都会对另一个产生影响。copy复制会比等号复制慢(很好理解)。

func main() {
	test1()
	fmt.Println("--------------")
	test2()
	fmt.Println("--------------")
	test3()
}

// 等号
func test1() {

	a := [3]int{0, 1, 2}
	s := a[1:2]
	fmt.Println(a, s) // [0 1 2] [1]
	s[0] = 11

	fmt.Println(a, s) // [0 11 2] [11]
	s = append(s, 12) // 底层还是a的容量
	s = append(s, 13) //再追加13之前s用的是a的部分空间，追加到13以后发生扩容搬家了
	fmt.Println(a, s) // [0 11 12] [11 12 13]
	s[0] = 21

	fmt.Println(a, s) // [0 11 12] [21 12 13]

}

//copy函数接收的参数类型为[]type形式，像上面等号复制里面的[3]int{0,1,2} 这种形式是不被认可的。
func test2() {

	a := []int{0, 1, 2}
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

func test3() {
	sl1 := make([]int, 3)
	sl1 = append(sl1, 11)
	sl1 = append(sl1, 22)
	sl1 = append(sl1, 33)
	fmt.Println(sl1) //[0 0 0 11 22 33]

	sl2 := make([]int, 3)
	sl2[0] = 11
	sl2[1] = 22
	sl2[2] = 33
	sl2[3] = 44 //越界   panic: runtime error: index out of range [3] with length 3
	fmt.Println(sl2)

}
