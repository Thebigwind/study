package main

import "fmt"

/*
如果一个值v可以被显式地转换为类型T，则此转换可以使用语法形式(T)(v)来表示。
在大多数情况下，特别是T为一个类型名（即一个标识符）时，此形式可简化为T(v)。

当一个值x可以被隐式转换为一个类型T，这同时也意味着x可以被显式转换为类型T。


如果两个类型表示着同一个类型，则它们的值可以相互隐式转换为这两个类型中的任意一个。
比如，
类型byte和uint8的任何值可以转换为这两个类型中的任意一个。
类型rune和int32的任何值可以转换为这两个类型中的任意一个。
类型[]byte和[]uint8的任何值可以转换为这两个类型中的任意一个。



底层类型相关的类型转换规则
给定一个非接口值x和一个非接口类型T，并假设x的类型为Tx，
如果类型Tx和T的底层类型相同（忽略掉结构体字段标签），则x可以被显式转换为类型T。
如果类型Tx和T中至少有一个是非定义类型并且它们的底层类型相同（考虑结构体字段标签），则x可以被隐式转换为类型T。
如果类型Tx和T的底层类型不同，但是两者都是非定义的指针类型并且它们的基类型的底层类型相同（忽略掉结构体字段标签），则x可以（而且只能）被显式转换为类型T。


*/

func main() {
	fmt.Printf("---------------------------切片开始-----------------------------------------\n")

	// 类型[]int、IntSlice和MySlice共享底层类型：[]int。
	type IntSlice []int
	type MySlice []int

	var s = []int{}
	var is = IntSlice{}
	var ms = MySlice{}
	var x struct {
		n int `foo`
	}
	var y struct {
		n int `bar`
	}

	// 这两行隐式转换编译不通过。
	/*
	   is = ms
	   ms = is
	*/

	// 必须使用显式转换。
	is = IntSlice(ms)
	ms = MySlice(is)
	x = struct {
		n int `foo`
	}(y)
	y = struct {
		n int `bar`
	}(x)

	// 这些隐式转换是没问题的。
	s = is
	is = s
	s = ms
	ms = s

	fmt.Printf("---------------------------切片结束-----------------------------------------\n")

	fmt.Printf("---------------------------指针开始------------------------------------------\n")

	type MyInt int
	type IntPtr *int
	type MyIntPtr *MyInt

	var pi = new(int)  // pi的类型为*int
	var ip IntPtr = pi // 没问题，因为底层类型相同
	// 并且pi的类型为非定义类型。

	// var _ *MyInt = pi // 不能隐式转换
	var _ = (*MyInt)(pi) // 显式转换是没问题的

	// 类型*int的值不能被直接转换为类型MyIntPtr，
	// 但是可以间接地转换过去。
	/*
	   var _ MyIntPtr = pi  // 不能隐式转换
	   var _ = MyIntPtr(pi) // 也不能显式转换
	*/
	var _ MyIntPtr = (*MyInt)(pi)  // 间接隐式转换没问题
	var _ = MyIntPtr((*MyInt)(pi)) // 间接显式转换没问题

	// 类型IntPtr的值不能被直接转换为类型MyIntPtr，
	// 但是可以间接地转换过去。
	/*
	   var _ MyIntPtr = ip  // 不能隐式转换
	   var _ = MyIntPtr(ip) // 也不能显式转换
	*/
	// 间接隐式或者显式转换都是没问题的。
	var _ MyIntPtr = (*MyInt)((*int)(ip))  // ok
	var _ = MyIntPtr((*MyInt)((*int)(ip))) // ok

	fmt.Printf("----------------------------指针结束----------------------------------------\n")

	fmt.Printf("----------------------------通道开始-----------------------------------------\n")

	/*
		给定一个通道值x，假设它的类型Tx是一个双向通道类型，T也是一个通道类型（无论是双向的还是单向的）。
		如果Tx和T的元素类型相同并且它们中至少有一个为非定义类型，则x可以被隐式转换为类型T。
	*/

	type C chan string
	type C1 chan<- string
	type C2 <-chan string

	var ca C
	var cb chan string

	cb = ca // ok，因为底层类型相同
	ca = cb // ok，因为底层类型相同

	// 这4行都满足此第3条转换规则的条件。
	var _, _ chan<- string = ca, cb // ok
	var _, _ <-chan string = ca, cb // ok
	var _ C1 = cb                   // ok
	var _ C2 = cb                   // ok

	// 类型C的值不能直接转换为类型C1或C2。
	/*
	   var _ = C1(ca) // compile error
	   var _ = C2(ca) // compile error
	*/

	// 但是类型C的值可以间接转换为类型C1或C2。
	var _ = C1((chan<- string)(ca)) // ok
	var _ = C2((<-chan string)(ca)) // ok
	var _ C1 = (chan<- string)(ca)  // ok
	var _ C2 = (<-chan string)(ca)  // ok
}
