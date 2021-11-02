package main

import "fmt"

func f(n int) (r int) {
	defer func() {
		r += n
		recover()
	}()

	var f func()

	defer f()
	f = func() {
		r += 2
	}
	return n + 1
}

func main() {
	fmt.Println(f(3)) //7

	test()
	test2()
	test3()
}

/*
第一步执行r = n +1，接着执行第二个 defer，由于此时 f() 未定义，引发异常，随即执行第一个 defer，异常被 recover()，程序正常执行，最后 return。
*/

func test() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a { //range 表达式是副本参与循环，就是说例子中参与循环的是 a 的副本，而不是真正的 a
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)
	//r =  [1 2 3 4 5]
	//a =  [1 12 13 4 5]
}

/*
就这个例子来说，假设 b 是 a 的副本，则 range 循环代码是这样的
for i, v := range b {
    if i == 0 {
        a[1] = 12
        a[2] = 13
    }
    r[i] = v
}
因此无论 a 被如何修改，其副本 b 依旧保持原值，并且参与循环的是 b，因此 v 从 b 中取出的仍旧是 a 的原值，而非修改后的值。
*/

func test2() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range &a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r) //r =  [1 12 13 4 5]
	fmt.Println("a = ", a) //a =  [1 12 13 4 5]
}

/*
使用 *[5]int 作为 range 表达式，其副本依旧是一个指向原数组 a 的指针，
因此后续所有循环中均是 &a 指向的原数组亲自参与的，
因此 v 能从 &a 指向的原数组中取出 a 修改后的值。

*/

func test3() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r) //r =  [1 12 13 4 5]
	fmt.Println("a = ", a) //a =  [1 12 13 4 5]

}

/*
切片在 go 的内部结构有一个指向底层数组的指针，当 range 表达式发生复制时，副本的指针依旧指向原底层数组，
所以对切片的修改都会反应到底层数组上，所以通过 v 可以获得修改后的数组元素
*/
