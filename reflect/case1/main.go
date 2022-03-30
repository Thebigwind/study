package main

import (
	"fmt"
	"reflect"
)

func main() {
	//变量类型为 interface{}
	rv := []interface{}{"hi", 42, func() {}}
	for _, v := range rv {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}
}
