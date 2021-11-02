package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	batch := 20
	batch -= 1
	fmt.Println(batch)
	//return
	start := time.Now().UnixNano()
	CountHELLO("/Users/luxuefeng/countJDTXT.txt", "HELLO")
	end := time.Now().UnixNano()
	fmt.Printf("耗时：%d ms\n", (end-start)/1e6)
}

func CountHELLO(filePath string, keyWord string) (error, int) {
	s1, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return err, 0
	}

	count := bytes.Count(s1, []byte(keyWord))
	fmt.Printf("%s count:%d\n", keyWord, count)
	return nil, count
}
