package main

import (
	"io/ioutil"
	"log"
)

func ReadFile(filePath string) []byte {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("Read error")
	}
	return content
}
