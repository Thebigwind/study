package main

import (
	"log"
	"strconv"
)

//二、后进先出
func main() {
	for i := 0; i < 6; i++ {
		defer log.Println("EDDYCJY" + strconv.Itoa(i) + ".")
	}

	log.Println("end.")
}

/*
$ go run main.go
2019/05/19 21:19:17 end.
2019/05/19 21:19:17 EDDYCJY5.
2019/05/19 21:19:17 EDDYCJY4.
2019/05/19 21:19:17 EDDYCJY3.
2019/05/19 21:19:17 EDDYCJY2.
2019/05/19 21:19:17 EDDYCJY1.
2019/05/19 21:19:17 EDDYCJY0.
*/
