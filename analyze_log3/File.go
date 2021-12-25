package main

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"strings"
)

//读取文件,加载到内存,并按行切分
//全部读入，然后按 \n 分割，存放到数组；再遍历数组
func LoadFile2Mem(filePath string, isGzip bool) ([]string, error) {
	s1, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	//当前先默认
	if isGzip {
		s2, err := ParseGzip(s1)
		if err != nil {
			return nil, err
		}
		return strings.Split(Bytes2str(s2), "\n"), nil
	} else {
		return strings.Split(Bytes2str(s1), "\n"), nil
	}

}

func ParseGzip(data []byte) ([]byte, error) {
	b := new(bytes.Buffer)
	binary.Write(b, binary.LittleEndian, data)
	r, err := gzip.NewReader(b)
	if err != nil {
		fmt.Printf("[ParseGzip] NewReader error: %v, maybe data is ungzip \n", err)
		return nil, err
	} else {
		defer r.Close()
		undatas, err := ioutil.ReadAll(r)
		if err != nil {
			fmt.Printf("[ParseGzip]  ioutil.ReadAll error: %v\n", err)
			return nil, err
		}
		return undatas, nil
	}
}
