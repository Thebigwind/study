package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var data = map[string]string{
		"1": "aa",
		"2": "bb",
		"3": "cc",
	}
	err := WriteJsonFile("/tmp/aa.json", data)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		return
	}
	fmt.Printf("success.\n")

}

func WriteJsonFile(path string, data interface{}) error {
	//var file *os.File
	file, err := os.Create(path)
	//file,err := os.OpenFile(path,os.O_TRUNC|os.O_CREATE,0666)
	if err != nil {
		return err
	}
	defer file.Close()
	// 写入UTF-8 BOM
	if _, err = file.WriteString("\xEF\xBB\xBF"); err != nil {
		return err
	}
	w := bufio.NewWriterSize(file, 50240)

	//result, err := json.MarshalIndent(data, "", "  ")
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = w.Write(result)
	if err != nil {
		return err
	}
	if err = w.Flush(); err != nil {
		return err
	}

	return nil
}
