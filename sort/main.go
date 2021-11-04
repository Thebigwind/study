package main

import (
	"fmt"
	"sort"
)

/*slice 排序示例*/
type Person struct {
	Age int
}

type PersonSlice []Person

func (s PersonSlice) Len() int           { return len(s) }
func (s PersonSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PersonSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

func main() {
	persons := PersonSlice{
		Person{
			Age: 1,
		},
		Person{
			Age: 5,
		},
		Person{
			Age: 2,
		},
	}
	sort.Sort(persons)
	fmt.Printf("after sort:%+v", persons)
}
