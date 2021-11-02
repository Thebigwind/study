
与 timer 结合，一般有两种玩法：实现超时控制，实现定期执行某个任务。

有时候，需要执行某项操作，但又不想它耗费太长时间，上一个定时器就可以搞定：

select {
case <-time.After(100 * time.Millisecond):
case <-s.stopc:
return false
}
等待 100 ms 后，如果 s.stopc 还没有读出数据或者被关闭，就直接结束。这是来自 etcd 源码里的一个例子，这样的写法随处可见。

定时执行某个任务，也比较简单：

func worker() {
ticker := time.Tick(1 * time.Second)
for {
select {
case <- ticker:
// 执行定时任务
fmt.Println("执行 1s 定时任务")
}
}