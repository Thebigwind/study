Timer实际上是一种单一事件的定时器，即经过指定的时间后触发一个事件，这个事件通过其本身提供的channel进行通知。之所以叫单一事件，是因为Timer只执行一次就结束，这也是Timer与Ticker的最重要的区别之一。

通过timer.NewTimer(d Duration)可以创建一个timer，参数即等待的时间，时间到来后立即触发一个事件。

type Timer struct { // Timer代表一次定时，时间到来后仅发生一个事件。
    C <-chan Time
    r runtimeTimer
}
Timer对外仅暴露一个channel，指定的时间到来时就往该channel中写入系统时间，也即一个事件。





