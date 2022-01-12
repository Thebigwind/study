package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println(time.Now().Format(time.RFC3339))
	//timeNow := time.Now()
	//timeEnd := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 23, 59, 59, 0, time.Local)
	//fmt.Printf("month:%s\n",time.Now().Month())
	//fmt.Printf("month:%d\n",time.Now().Hour())
	//fmt.Printf("month:%s\n",time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("month:%s\n", time.Now().Format("0515010402"))
	//fmt.Printf("time:%s",timeEnd)
	fmt.Printf("\n")
	// get the location
	location, _ := time.LoadLocation("UTC") //"Europe/Rome"
	// this should give you time in location
	t := time.Now().In(location)
	fmt.Println(t)

	fmt.Printf("\n")
	fmt.Println(time.Now().ISOWeek())
	fmt.Println(time.Now().Unix())
}
