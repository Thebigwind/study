package main

import "fmt"

//1
func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

//5
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// 1
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

//0
func f4() int {
	r := 0
	defer func() {
		r++
	}()
	return r
}
func main() {
	fmt.Println(f1()) //1
	fmt.Println(f2()) //5
	fmt.Println(f3()) //1
	fmt.Println(f4()) //0
}
