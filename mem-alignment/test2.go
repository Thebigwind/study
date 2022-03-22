package main

import (
	"fmt"
	"unsafe"
)

type Part1 struct {
	a bool  //1  初始地址，偏移量为 0。占用了第 1 位
	b int32 //4  根据规则 1，其偏移量必须为 4 的整数倍。确定偏移量为 4，因此 2-4 位为 Padding。而当前数值从第 5 位开始填充，到第 8 位。如下：axxx|bbbb
	c int8  //1 根据规则 1，其偏移量必须为 1 的整数倍。当前偏移量为 8。不需要额外对齐，填充 1 个字节到第 9 位。如下：axxx|bbbb|c…
	d int64 //8 其偏移量必须为 8 的整数倍。确定偏移量为 16，因此 9-16 位为 Padding。而当前数值从第 17 位开始写入，到第 24 位。如下：axxx|bbbb|cxxx|xxxx|dddd|dddd
	e byte  //1 根据规则 1，其偏移量必须为 1 的整数倍。当前偏移量为 24。不需要额外对齐，填充 1 个字节到第 25 位。如下：axxx|bbbb|cxxx|xxxx|dddd|dddd|e…
}

/*
在每个成员变量进行对齐后，根据规则 2，整个结构体本身也要进行字节对齐，因为可发现它可能并不是 2^n，不是偶数倍。显然不符合对齐的规则

根据规则 2，可得出对齐值为 8。现在的偏移量为 25，不是 8 的整倍数。因此确定偏移量为 32。对结构体进行对齐

结果
Part1 内存布局：axxx|bbbb|cxxx|xxxx|dddd|dddd|exxx|xxxx
*/

type Part2 struct {
	e byte  //1 初始地址，偏移量为 0。占用了第 1 位
	c int8  //1 根据规则 1，其偏移量必须为 1 的整数倍。当前偏移量为 2。不需要额外对齐
	a bool  //1 根据规则 1，其偏移量必须为 1 的整数倍。当前偏移量为 3。不需要额外对齐
	b int32 //4 根据规则 1，其偏移量必须为 4 的整数倍。确定偏移量为 4，因此第 3 位为 Padding。而当前数值从第 4 位开始填充，到第 8 位。如下：ecax|bbbb
	d int64 //8 根据规则 1，其偏移量必须为 8 的整数倍。当前偏移量为 8。不需要额外对齐，从 9-16 位填充 8 个字节。如下：ecax|bbbb|dddd|dddd
}

/*
整体对齐
符合规则 2，不需要额外对齐

结果
Part2 内存布局：ecax|bbbb|dddd|dddd
*/

func main() {

	part1 := Part1{}
	part2 := Part2{}

	fmt.Printf("part1 size: %d, align: %d\n", unsafe.Sizeof(part1), unsafe.Alignof(part1))
	fmt.Printf("part2 size: %d, align: %d\n", unsafe.Sizeof(part2), unsafe.Alignof(part2))
	/*
		part1 size: 32, align: 8
		part2 size: 16, align: 8
	*/
}

/*
默认系数
在不同平台上的编译器都有自己默认的 “对齐系数”，可通过预编译命令 #pragma pack(n) 进行变更，n 就是代指 “对齐系数”。一般来讲，我们常用的平台的系数如下：

32 位：4
64 位：8
*/
