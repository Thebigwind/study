package main

func Slice() {
	s := make([]int, 10000, 10000)

	for index, _ := range s {
		s[index] = index
	}
}

func main() {
	Slice()
}

/*
bogon:case2 luxuefeng$ go build -gcflags=-m
# study/mem-leak/case2
./check.go:11:6: can inline main
./check.go:4:11: make([]int, 1000, 1000) does not escape
bogon:case2 luxuefeng$
bogon:case2 luxuefeng$ go build -gcflags=-m
# study/mem-leak/case2
./check.go:11:6: can inline main
./check.go:4:11: make([]int, 10000, 10000) escapes to heap

*/
