几种导致当前协程永久阻塞的方法
无需引入任何包，我们可以使用下面几种方法使当前协程永久阻塞：
向一个永不会被接收数据的通道发送数据。
make(chan struct{}) <- struct{}{}
// 或者
make(chan<- struct{}) <- struct{}{}
从一个未被并且将来也不会被发送数据的（并且保证永不会被关闭的）通道读取数据。
<-make(chan struct{})
// 或者
<-make(<-chan struct{})
// 或者
for range make(<-chan struct{}) {}
从一个nil通道读取或者发送数据。
chan struct{}(nil) <- struct{}{}
// 或者
<-chan struct{}(nil)
// 或者
for range chan struct{}(nil) {}
使用一个不含任何分支的select流程控制代码块。
select{}