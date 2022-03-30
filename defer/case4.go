package main

import "log"

//四、异常处理
func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Println("EDDYCJY.")
		}
	}()

	panic("end.")
}

//$ go run main.go
//2019/05/20 22:22:57 EDDYCJY.
