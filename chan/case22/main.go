package main

import "fmt"

func main() {
	strMap := make(map[string]string)

	test(strMap)
	fmt.Printf("%+v\n", strMap)

	strSli := make([]string, 0)

	test2(strSli)
	fmt.Printf("%+v", strSli)
}
func test(m map[string]string) {
	m["1"] = "111"
	m["2"] = "222"
}

func test2(sli []string) {
	sli = append(sli, "aaa")
	sli = append(sli, "bbb")
}
