package _defer

import "log"

//一、延迟调用
func main() {
	defer log.Println("EDDYCJY.")

	log.Println("end.")
}

/*
$ go run main.go
2019/05/19 21:15:02 end.
2019/05/19 21:15:02 EDDYCJY.
*/
