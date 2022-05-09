package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
	"unsafe"
)

func getFileSize(filename string) (int64, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func main() {
	start := time.Now().UnixNano()

	keyWord := "HELLO"
	filePath := "/Users/me/Downloads/countJDTXT.txt"

	filesize, err := getFileSize(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("size:%d\n",filesize)

	numGroutime := runtime.NumCPU()
	//fmt.Println(numGroutime)//8
	segSize := filesize / int64(numGroutime)
	numChan := make(chan int, numGroutime)

	f, err := os.OpenFile(filePath, os.O_RDONLY, 0660)
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	defer f.Close()

	var wg = sync.WaitGroup{}
	wg.Add(numGroutime)
	for i := 0; i < numGroutime; i++ {
		go func(i int) {
			defer wg.Done()
			pos := segSize * int64(i)
			readBuf := make([]byte, segSize+int64(len(keyWord)-1))
			readString, err := f.ReadAt(readBuf, pos)
			if err != nil && err != io.EOF {
				return
			}
			err, countNum := CountHELLO(readBuf[:readString], keyWord)
			if err != nil {
				return
			}
			numChan <- countNum
		}(i)
	}
	//fmt.Println(runtime.NumGoroutine())
	wg.Wait()

	totalNum := 0
	for i := 0; i < numGroutime; i++ {
		totalNum += <-numChan
	}
	fmt.Println(totalNum)

	end := time.Now().UnixNano()
	fmt.Printf("耗时：%d ms\n", (end-start)/1e6)
}

//统计关键字
func CountHELLO(s []byte, keyWord string) (error, int) {

	//count := strings.Count(string(s1),keyWord)
	//count := strings.Count(bytes2str(s1),keyWord)
	count := bytes.Count(s, str2bytes(keyWord))
	//fmt.Printf("%s count:%d\n", keyWord, count)
	return nil, count
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
