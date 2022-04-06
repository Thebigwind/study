package main

import "fmt"

func main() {
	arr := []string{}
	for i, v := range arr {
		fmt.Printf("%v,%v", i, v)
	}
	fmt.Printf("xxxx")
}
