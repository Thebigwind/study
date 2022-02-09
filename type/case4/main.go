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
	fmt.Println(e)

	//1.642386514e+09
	//1.642386307e+09
	u, e = strconv.ParseFloat("1.642388885e+09", 64)
	fmt.Println(int64(u))

	u, e = strconv.ParseFloat("1.642388883e+09", 64)
	fmt.Println(int64(u))
}
