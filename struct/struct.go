package main

/*
一个结构体类型的尺寸取决于它的各个字段的类型尺寸和这些字段的排列顺序。
为了程序执行性能，编译器需要保证某些类型的值在内存中存放时必须满足特定的内存地址对齐要求。
地址对齐可能会造成相邻的两个字段之间在内存中被插入填充一些多余的字节。
所以，一个结构体类型的尺寸必定不小于（常常会大于）此结构体类型的各个字段的类型尺寸之和。


为了防止在函数传参和通道操作中因为值复制代价太高而造成的性能损失，
我们应该避免使用大尺寸的结构体和数组类型做为参数类型和通道的元素类型，
应该在这些场合下使用基类型为这样的大尺寸类型的指针类型

要考虑到太多的指针将会增加垃圾回收的压力。所以到底应该使用大尺寸类型还是以大尺寸类型为基类型的指针类型做为参数类型或通道的元素类型取决于具体的应用场景

如果一个数组或者切片的元素类型是一个大尺寸类型，我们应该避免在for-range循环中使用双循环变量来遍历这样的数组或者切片类型的值中的元素。
因为，在遍历过程中，每个元素将被复制给第二个循环变量一次。
*/
