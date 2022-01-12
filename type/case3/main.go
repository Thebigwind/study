package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := fmt.Sprintf("%v", float64(74125748994051))
	fmt.Println(s)

	u, e := strconv.ParseFloat(s, 64)
	fmt.Println(int64(u))
	fmt.Println(u)
	fmt.Println(e)
}
