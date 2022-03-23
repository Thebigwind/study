堆（Heap）：一般来讲是人为手动进行管理，手动申请、分配、释放。一般所涉及的内存大小并不定，一般会存放较大的对象。另外其分配相对慢，涉及到的指令动作也相对多
栈（Stack）：由编译器进行管理，自动申请、分配、释放。一般不会太大，我们常见的函数参数（不同平台允许存放的数量不同），局部变量等等都会存放在栈上

什么是逃逸分析
在编译程序优化理论中，逃逸分析是一种确定指针动态范围的方法，简单来说就是分析在程序的哪些地方可以访问到该指针


逃逸分析就是确定一个变量要放堆上还是栈上，规则如下：

    1.是否有在其他地方（非局部）被引用。只要有可能被引用了，那么它一定分配到堆上。否则分配到栈上
    2.即使没有被外部引用，但对象过大，无法存放在栈区上。依然有可能分配到堆上
    

如果变量都分配到堆上了会出现什么事情？

    1.垃圾回收（GC）的压力不断增大
    2.申请、分配、回收内存的系统开销增大（相对于栈）
    3.动态分配产生一定量的内存碎片


怎么确定是否逃逸

第一，通过编译器命令，就可以看到详细的逃逸分析过程。而指令集 -gcflags 用于将标识参数传递给 Go 编译器，涉及如下：

    go build -gcflags '-m -l' main.go

第二，通过反编译命令查看

    go tool compile -S main.go




func main() {
    str := new(string)
    *str = "EDDYCJY"
}

    $ go build -gcflags '-m -l' main.go
    # command-line-arguments
    ./main.go:4:12: main new(string) does not escape

显然，该对象分配到栈上了。很核心的一点就是它有没有被作用域之外所引用，而这里作用域仍然保留在 main 中，因此它没有发生逃逸


func main() {
    str := new(string)
    *str = "EDDYCJY"
	fmt.Println(str)
}

    $ go build -gcflags '-m -l' main.go
    # command-line-arguments
    ./main.go:9:13: str escapes to heap
    ./main.go:6:12: new(string) escapes to heap
    ./main.go:9:13: main ... argument does not escape

str 变量逃到了堆上，也就是该对象在堆上分配.fmt.Println(), 当形参为 interface 类型时，在编译阶段编译器无法确定其具体的类型。因此会产生逃逸，最终分配到堆上




需要注意：

    静态分配到栈上，性能一定比动态分配到堆上好
    底层分配到堆，还是栈。实际上对你来说是透明的，不需要过度关心
    每个 Go 版本的逃逸分析都会有所不同（会改变，会优化）
    直接通过 go build -gcflags '-m -l' 就可以看到逃逸分析的过程和结果
    到处都用指针传递并不一定是最好的，要用对