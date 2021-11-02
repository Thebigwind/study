package main

//import "time"
/*
超时机制（timeout）
在一些请求/回应用例中，一个请求可能因为种种原因导致需要超出预期的时长才能得到回应，有时甚至永远得不到回应。 对于这样的情形，我们可以使用一个超时方案给请求者返回一个错误信息。 使用选择机制可以很轻松地实现这样的一个超时方案。

下面这个例子展示了如何实现一个支持超时设置的请求：
*/
/*
func requestWithTimeout(timeout time.Duration) (int, error) {
	c := make(chan int)
	go doRequest(c) // 可能需要超出预期的时长回应

	select {
	case data := <-c:
		return data, nil
	case <-time.After(timeout):
		return 0, errors.New("超时了！")
	}
}
*/
