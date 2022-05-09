package main

import (
	"fmt"
	"reflect"
)

func main() {
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil) // nil, true
	fmt.Println(in, in == nil)     //nil, true

	in = data
	fmt.Println(in, in == nil) //nil,false
	fmt.Println(in, IsNil(in)) //nil,true
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
