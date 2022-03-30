package main

import (
	"fmt"
	"reflect"
)

type Blog struct {
	string
}

func main() {
	blog := Blog{"煎鱼"}
	typeof := reflect.TypeOf(blog)
	fmt.Println(typeof.String())
}

//reflect.TypeOf 成功解析出 blog 变量的类型是 main.Blog，

//TypeOf 方法中主要涉及三块操作，分别如下
//使用 unsafe.Pointer 方法获取任意类型且可寻址的指针值。
//利用 emptyInterface 类型进行强制的 interface 类型转换。
//调用 toType 方法转换为可供外部使用的 Type 类型
