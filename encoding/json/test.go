package main

import (
	"encoding/json"
	"fmt"
)

type Price struct {
	Date  string `json:"date"`
	Price string `json:"price"`
	Quota string `json:"quota"`
}

type Node struct {
	Data string
	Next *Node
}

type Node2 struct {
	Data string
}

func main() {
	/*
		str := `{&#34;date&#34;:&#34;2021-09-10&#34;,&#34;price&#34;:9900,&#34;quota&#34;:1}`
		p := Price{}
		err := json.Unmarshal([]byte(str),&p)
		if err != nil{
			fmt.Printf("err:%v",err)
		}
		fmt.Printf("price:%+v",p)
	*/
	priceData, err := json.Marshal(nil)
	if err != nil {
		fmt.Printf("errr:", err.Error())
	}

	fmt.Printf("data:%+v", string(priceData))
	fmt.Printf("\n-----\n")
	pArr := make([]Price, 0)
	priceData, err = json.Marshal(pArr)
	if err != nil {
		fmt.Printf("errr:", err.Error())
	}
	fmt.Printf("data:%+v", string(priceData))

	fmt.Printf("\n-----------------------------------------------\n")
	node := &Node{
		Data: "John",
	}
	node.Next = node

	_, err = json.Marshal(node)
	_, ok := err.(*json.UnsupportedValueError)

	fmt.Println("UnsupportedValueError ", ok)
	fmt.Printf("\n-----------------------------------------------\n")
	node2 := &Node2{
		Data: "John",
	}

	_, err = json.Marshal(node2)
	_, ok2 := err.(*json.UnsupportedValueError)
	fmt.Println("UnsupportedValueError ", ok2)

}
