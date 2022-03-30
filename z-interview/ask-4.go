package main

import "log"

type MyErr struct {
	Msg string
}

func main() {
	var e error
	e = GetErr()
	log.Println(e)
	log.Println(e == nil) //false
}

func GetErr() *MyErr {
	return nil
}

func (m *MyErr) Error() string {
	return "脑子进煎鱼了"
}
