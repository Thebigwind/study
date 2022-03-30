
垃圾回收（GC）是一种自动管理内存的机制，垃圾回收器会去尝试回收程序不再使用的对象及其占用的内存

Stop The World（STW）

STW 代指在执行某个垃圾回收算法的某个阶段时，需要将整个应用程序暂停去处理 GC 相关的工作事项

    标记开始：在开始标记时，准备根对象的扫描，会打开写屏障（Write Barrier） 和 辅助GC（mutator assist），而回收器和应用程序是并发运行的，因此会暂停当前正在运行的所有 Goroutine。
    并发标记中：标记阶段，主要目的是标记堆内存中仍在使用的值。
    标记结束：在完成标记任务后，将重新扫描部分根对象，这时候会禁用写屏障（Write Barrier）和辅助GC（mutator assist），而标记阶段和应用程序是并发运行的，所以在标记阶段可能会有新的对象产生，因此在重新扫描时需要进行 STW


GC 频率

GOGC 变量设置初始垃圾收集器的目标百分比值

当新分配的数值与上一次收集后剩余的实时数值的比例达到设置的目标百分比时，就会触发 GC，默认值为 GOGC=100。
如果将其设置为 GOGC=off 可以完全禁用垃圾回收器



版本	  GC   算法	                         STW 时间
Go   1.0	STW（强依赖 tcmalloc）	     百ms到秒级别
Go   1.3	Mark STW, Sweep 并行	         百ms级别
Go   1.5	三色标记法, 并发标记清除。同时运行时从 C 和少量汇编，改为 Go 和少量汇编实现	10-50ms级别
Go   1.6	1.5 中一些与并发 GC 不协调的地方更改，集中式的 GC 协调协程，改为状态机实现	5ms级别
Go   1.7	GC 时由 mark 栈收缩改为并发，span 对象分配机制由 freelist 改为 bitmap 模式，SSA引入	ms级别
Go   1.8	混合写屏障（hybrid write barrier）, 消除 re-scanning stack	sub ms
Go   1.12	Mark Termination 流程优化	   sub ms, 但几乎减少一半


