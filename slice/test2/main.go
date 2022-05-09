package main

import "fmt"

func main() {
	var s []string
	for i := 0; i < 9999; i++ {
		go func() {
			s = append(s, "脑子进煎鱼了")
		}()
	}

	fmt.Printf("进了 %d 只煎鱼", len(s))
}
