package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	NewStatPvMap()
	// /Users/luxuefeng/Downloads/api-user-csv
	// /Users/luxuefeng/Downloads/service-uic-csv
	// php-service-uic-nei-csv
	// php-center-nei-csv
	// go-center-nei-csv
	// api-user-nei-csv
	pathList, err := GetPathList("/Users/luxuefeng/Downloads/php-service-uic-nei-csv")
	if err != nil {
		return
	}

	for _, file := range pathList {
		if err := GetCsvData(file); err != nil {
			return
		}
	}

	WriteCsv("/Users/luxuefeng/go/study/ana4-nei/php-service-uic-nei-ok.csv", "service-uic")

}
func GetPathList(dirPath string) ([]string, error) {
	fileList := make([]string, 0)
	rd, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return nil, err
	}
	for i := range rd {
		filePath := dirPath + "/" + rd[i].Name()
		fileList = append(fileList, filePath)
	}

	//输出fileList
	for i, v := range fileList {
		fmt.Printf("index:%d,value:%s\n", i, v)
	}

	return fileList, nil
}

func GetCsvData(filepath string) error {
	//read file
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	rd := bufio.NewReaderSize(file, 2097152) //2097152
	for {

		lineText, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		if lineText == "" {
			continue
		}

		lineText = strings.Trim(lineText, "\n")
		mList := strings.Split(lineText, "\"")
		//fmt.Printf("mList:%+v",mList)
		//for k,v := range mList{
		//	fmt.Printf("key:%d,value:%s\n",k,v)
		//}
		//os.Exit(0)
		if len(mList) < 2 {
			fmt.Printf("err:")
			continue
		}
		uri := strings.Trim(mList[1], "\" ")
		if strings.HasPrefix(uri, "/app") {
			aList := strings.Split(uri, "/")
			a := "/"
			for _, v := range aList[5:] {
				a = a + v + "/"
			}

			uri = strings.TrimSuffix(a, "/")
		} else if strings.HasPrefix(uri, "/user/get/") {
			uri = "/user/get"
		}
		pv := strings.Trim(mList[0], " ")
		//fmt.Printf("pv:%v\n",pv)
		num, err := strconv.ParseInt(pv, 10, 64)
		if err != nil {
			fmt.Printf("err:", err.Error())
			continue
		}
		CaculatePv(uri, num)
	}
	return nil
}

func CaculatePv(uri string, num int64) {
	GlobalStatPvData.lc.Lock()
	v, exist := GlobalStatPvData.PvData[uri]
	if exist {
		GlobalStatPvData.PvData[uri] = v + num
	} else {
		GlobalStatPvData.PvData[uri] = num
	}
	GlobalStatPvData.lc.Unlock()
}

type PVMap struct {
	PvData map[string]int64
	lc     *sync.Mutex
}

var GlobalStatPvData *PVMap = nil

func GetStatPVMap() *PVMap {
	return GlobalStatPvData
}

//var StatPvData PVMap
func NewStatPvMap() *PVMap {
	pvdata := make(map[string]int64)
	pvmap := &PVMap{
		PvData: pvdata,
		lc:     new(sync.Mutex),
	}

	GlobalStatPvData = pvmap

	return pvmap
}

func WriteCsv(fileName string, serviceName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file is failed, err: ", err)
	}
	defer file.Close()
	// 写入UTF-8 BOM，防止中文乱码
	file.WriteString("\xEF\xBB\xBF")
	w := csv.NewWriter(file)
	w.Write([]string{"服务", "接口", "pv"})
	for k, v := range GlobalStatPvData.PvData {
		w.Write([]string{serviceName, k, strconv.FormatInt(v, 10)})
	}
	// 写文件需要flush，不然缓存满了，后面的就写不进去了，只会写一部分
	w.Flush()
}
