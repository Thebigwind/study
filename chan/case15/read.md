无阻塞地检查一个通道是否已经关闭
假设我们可以保证没有任何协程会向一个通道发送数据，则我们可以使用下面的代码来（并发安全地）检查此通道是否已经关闭，此检查不会阻塞当前协程。
func IsClosed(c chan T) bool {
select {
case <-c:
return true
default:
}
return false
}
此方法常用来查看某个期待中的通知是否已经来临。此通知将由另一个协程通过关闭一个通道来发送。



峰值限制（peak/burst limiting）
将通道用做计数信号量用例和通道尝试（发送或者接收）操作结合起来可用实现峰值限制。 峰值限制的目的是防止过大的并发请求数。

下面是对将通道用做计数信号量一节中的最后一个例子的简单修改，从而使得顾客不再等待而是离去或者寻找其它酒吧。

...
bar24x7 := make(Bar, 10) // 此酒吧只能同时招待10个顾客
for customerId := 0; ; customerId++ {
time.Sleep(time.Second)
consumer := Consumer{customerId}
select {
case bar24x7 <- consumer: // 试图进入此酒吧
go bar24x7.ServeConsumer(consumer)
default:
log.Print("顾客#", customerId, "不愿等待而离去")
}
}
...