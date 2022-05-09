package main

import "fmt"

func main() {

	s := [3]int{1, 2, 3}
	//a := s[:0]
	//b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Printf("%v,%v", len(c), cap(c))
}
