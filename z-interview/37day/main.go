package main

import (
	"fmt"
)

func main() {
	ch := make(chan *int, 5)

	//sender
	input := []int{1, 2, 3, 4, 5}

	go func() {
		for _, v := range input {
			ch <- &v
		}
		close(ch)
	}()
	//receiver
	for v := range ch {
		fmt.Println(*v)
		fmt.Println(v)
	}

}

/*
5
0xc000018050
5
0xc000018050
5
0xc000018050
5
0xc000018050
5
0xc000018050
*/

/*
解决方案：

引入一个中间变量，每次迭代都重新声明一个变量 temp ，赋值后再将其地址发送给 ch ：
for _, v := range input {
  temp := v
  ch <- &temp
}

抑或直接引用数据的内存（推荐，无需开辟新的内存空间）：
for k, _ := range input {
  c <- &input[k]
}
再次运行，就可看到预期的效果。以上方案是用于讨论 range 语句带来的问题，当然，平时还是尽量避免使用指针类型的通道。
*/
