package _defer

import "log"

//三、运行时间点
func main() {
	func() {
		defer log.Println("defer.EDDYCJY.")
	}()

	log.Println("main.EDDYCJY.")
}

/*
$ go run main.go
2019/05/22 23:30:27 defer.EDDYCJY.
2019/05/22 23:30:27 main.EDDYCJY.
*/
