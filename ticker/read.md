Ticker是周期性定时器，即周期性的触发一个事件，通过Ticker本身提供的管道将事件传递出去。

Ticker的数据结构与Timer完全一致：

type Ticker struct {
    C <-chan Time
    r runtimeTimer
}

Ticker对外仅暴露一个channel，指定的时间到来时就往该channel中写入系统时间，也即一个事件。
在创建Ticker时会指定一个时间，作为事件触发的周期。这也是Ticker与Timer的最主要的区别。


创建一个Ticker后，紧跟着使用defer语句关闭Ticker总是好的习惯。
因为，有可能别人无意间拷贝了你的部分代码，而忽略了关闭Ticker的动作。

使用time.NewTicker()来创建一个定时器；
使用Stop()来停止一个定时器；
定时器使用完毕要释放，否则会产生资源泄露；

