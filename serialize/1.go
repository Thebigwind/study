package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/techoner/gophp"
)

var m0 = map[string]string{"one": "1", "two": "2", "three": "3"}
var m1 = map[string]int{"one": 1, "two": 2, "three": 3}
var m2 = map[string]interface{}{
	"aa": 23,
	"bb": 45,
	"cc": 456,
	"dd": true,
	"ee": "adfasdfa",
}

type StreetGeoData struct {
	A string `json:"A"`
	B string `json:"B"`
	C string `json:"C"`
	D string `json:"D"`
	E string `json:"E"`
	F string `json:"F"`
	G string `json:"G"`
	H string `json:"H"`
}

var dd = StreetGeoData{
	A: "Cn",
	B: "sds",
}
var m3 = map[string]interface{}{
	"aa": 23,
	"bb": 45,
	"cc": 456,
	"dd": true,
	"ee": "adfasdfa",
	"ff": dd,
}

func main() {
	fmt.Println("-------------")

	test(m2)
	fmt.Println()
	fmt.Println("-------------")

	test(m3)

	fmt.Println("-------------")
	test2(m1)
}

func test(data interface{}) {
	jsonbyte, err := gophp.Serialize(data)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf(string(jsonbyte))
	}
}

func test2(data interface{}) {
	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)
	// Encoding the map
	err := e.Encode(data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("m2:%s\n", string(b.Bytes()))

	var decodedMap map[string]int
	d := gob.NewDecoder(b)

	// Decoding the serialized m2
	err = d.Decode(&decodedMap)
	if err != nil {
		panic(err)
	}

	// Ta da! It is a map!
	fmt.Printf("%#v\n", decodedMap)
}
