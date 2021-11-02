package main

/*
通过调用reflect.TypeOf函数，我们可以从一个任何非接口类型的值创建一个reflect.Type值。 此reflect.Type值表示着此非接口值的类型。
通过此值，我们可以得到很多此非接口类型的信息。

当然，我们也可以将一个接口值传递给一个reflect.TypeOf函数调用，但是此调用将返回一个表示着此接口值的动态类型的reflect.Type值。
实际上，reflect.TypeOf函数的唯一参数的类型为interface{}， reflect.TypeOf函数将总是返回一个表示着此唯一接口参数值的动态类型的reflect.Type值。
那如何得到一个表示着某个接口类型的reflect.Type值呢？ 我们必须通过下面将要介绍的一些间接途径来达到这一目的。

类型reflect.Type为一个接口类型，它指定了若干方法。 通过这些方法，我们能够观察到一个reflect.Type值所表示的Go类型的各种信息。
这些方法中的有些适用于所有种类的类型，有些只适用于一种或几种类型。 通过不合适的reflect.Type属主值调用某个方法将在运行时产生一个恐慌。
请阅读reflect代码库中各个方法的文档来获取如何正确地使用这些方法。


*/
import "fmt"
import "reflect"

func main() {
	type A = [16]int16
	var c <-chan map[A][]byte
	tc := reflect.TypeOf(c)
	fmt.Println(tc.Kind())    // chan
	fmt.Println(tc.ChanDir()) // <-chan
	tm := tc.Elem()
	ta, tb := tm.Key(), tm.Elem()
	fmt.Println(tm.Kind(), ta.Kind(), tb.Kind()) // map array slice
	tx, ty := ta.Elem(), tb.Elem()

	// byte是uint8类型的别名。
	fmt.Println(tx.Kind(), ty.Kind()) // int16 uint8
	fmt.Println(tx.Bits(), ty.Bits()) // 16 8
	fmt.Println(tx.ConvertibleTo(ty)) // true
	fmt.Println(tb.ConvertibleTo(ta)) // false

	// 切片类型和映射类型都是不可比较类型。
	fmt.Println(tb.Comparable()) // false
	fmt.Println(tm.Comparable()) // false
	fmt.Println(ta.Comparable()) // true
	fmt.Println(tc.Comparable()) // true

}

/*
reflect.Value类型和值
类似的，我们可以通过调用reflect.ValueOf函数，从一个非接口类型的值创建一个reflect.Value值。 此reflect.Value值代表着此非接口值。
和reflect.TypeOf函数类似，reflect.ValueOf函数也只有一个interface{}类型的参数。 当我们将一个接口值传递给一个reflect.ValueOf函数调用时，
此调用返回的是代表着此接口值的动态值的一个reflect.Value值。 我们必须通过间接的途径获得一个代表一个接口值的reflect.Value值。

被一个reflect.Value值代表着的值常称为此reflect.Value值的底层值（underlying value）。

reflect.Value类型有很多方法。 我们可以调用这些方法来观察和操纵一个reflect.Value属主值表示的Go值。
这些方法中的有些适用于所有种类类型的值，有些只适用于一种或几种类型的值。 通过不合适的reflect.Value属主值调用某个方法将在运行时产生一个恐慌。
请阅读reflect代码库中各个方法的文档来获取如何正确地使用这些方法。

一个reflect.Value值的CanSet方法将返回此reflect.Value值代表的Go值是否可以被修改（可以被赋值）。
如果一个Go值可以被修改，则我们可以调用对应的reflect.Value值的Set方法来修改此Go值。
注意：reflect.ValueOf函数直接返回的reflect.Value值都是不可修改的。


*/

func test() {
	n := 123
	p := &n
	vp := reflect.ValueOf(p)
	fmt.Println(vp.CanSet(), vp.CanAddr()) // false false
	vn := vp.Elem()                        // 取得vp的底层指针值引用的值的代表值
	fmt.Println(vn.CanSet(), vn.CanAddr()) // true true
	vn.Set(reflect.ValueOf(789))           // <=> vn.SetInt(789)
	fmt.Println(n)                         // 789
}

/*
reflect标准库包中也提供了一些对应着内置函数或者各种非反射功能的函数。
下面这个例子展示了如何利用这些函数将一个自定义泛型函数绑定到不同的类型的函数值上。

*/

func InvertSlice(args []reflect.Value) (result []reflect.Value) {
	inSlice, n := args[0], args[0].Len()
	outSlice := reflect.MakeSlice(inSlice.Type(), 0, n)
	for i := n - 1; i >= 0; i-- {
		element := inSlice.Index(i)
		outSlice = reflect.Append(outSlice, element)
	}
	return []reflect.Value{outSlice}
}

func Bind(p interface{}, f func([]reflect.Value) []reflect.Value) {
	// invert代表着一个函数值。
	invert := reflect.ValueOf(p).Elem()
	invert.Set(reflect.MakeFunc(invert.Type(), f))
}

func test2() {
	var invertInts func([]int) []int
	Bind(&invertInts, InvertSlice)
	fmt.Println(invertInts([]int{2, 3, 5})) // [5 3 2]

	var invertStrs func([]string) []string
	Bind(&invertStrs, InvertSlice)
	fmt.Println(invertStrs([]string{"Go", "C"})) // [C Go]
}
