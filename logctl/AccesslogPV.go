package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type AccessLogPvParams struct {
	LogDirPath   string
	DomainKey    string
	GoroutineNum int
	IsGzip       bool
	CsvName      string
	ServiceName  string
}

//使用命令分析数据
func DoAnalyzePVCommand(params AccessLogPvParams) error {

	//
	//zcat neiwang-lb-20211217-*.log |grep "service-uic.xiaozhu.com\|service-uicproxy.xiaozhu.com" | awk '{print $6}' | sort | uniq -c | sort -n -k 1 -r -> service-uic-nei-17.csv &

	domainList := strings.Split(params.DomainKey, ",")
	domainKey := ""
	if len(domainList) == 1 {
		domainKey = params.DomainKey
	} else {
		domainKey = strings.Join(domainList, "\\|")
	}

	//csvpath
	csvPath := ""
	if params.CsvName == "" {
		csvPath = "default-" + strconv.FormatInt(time.Now().Unix(), 10) + "-.csv"
	} else {
		_, csvPath = GetPath(params.CsvName)
	}

	exc := ""
	if params.IsGzip {
		exc = "zcat" + " " + params.LogDirPath + "/* |grep " + domainKey + " |awk '{print $6}' | sort | uniq -c | sort -n -k 1 -r  -> " + csvPath + " &"
	} else {
		exc = "grep" + domainKey + " " + params.LogDirPath + "/* |awk '{print $6}' | sort | uniq -c | sort -n -k 1 -r  -> " + csvPath + " &"
	}

	if err, _ := ExecShell(exc); err != nil {
		return err
	}
	/*
	   2144504 "/user/checkUserSession"
	    119470 "/user/getTag"
	    119230 "/user/getKaTag"
	*/
	//结果文件需要二次处理，返回指定格式

	return nil
}

//使用程序分析数据
func DoAnalyzePv(params AccessLogPvParams) error {
	//服务的几个域名关键字
	domainList := strings.Split(params.DomainKey, ",")
	//根据日志目录获取日志文件列表
	filePahtList, err := GetPathList(params.LogDirPath)
	if err != nil {
		return err
	}
	if params.GoroutineNum <= 1 {
		//遍历文件列表，按文件统计
		for i, filePath := range filePahtList {
			fmt.Printf("i:%d,filepath:%s\n", i, filePath)
			if err := StatFilePV(domainList, filePath, params.IsGzip); err != nil {
				fmt.Printf("err:%s\n", err.Error())
				return err
			}
			fmt.Printf("当前结果:%+v\n", i, GlobalStatPvData.PvData)
		}
	} else {
		//起多个goroutine
		fileChan := make(chan string, len(filePahtList))
		//将文件路径发送到channel,然后多个goroutine从 fileChan 中取文件
		for _, filePath := range filePahtList {
			fileChan <- filePath
		}
		close(fileChan)

		wg := sync.WaitGroup{}
		wg.Add(params.GoroutineNum)
		for i := 0; i < params.GoroutineNum; i++ {
			go Consumer(&wg, fileChan, domainList, params.IsGzip)
		}
		wg.Wait()
	}

	//写csv文件
	csvPath := ""
	if params.CsvName == "" {
		csvPath = "default-" + strconv.FormatInt(time.Now().Unix(), 10) + "-.csv"
	}
	WriteCsv(csvPath, params.ServiceName)
	return nil
}

//PV统计
func StatFilePV(domainList []string, filePath string, isGzip bool) error {
	lineArr, err := LoadFile2Mem(filePath, isGzip)
	if err != nil {
		return err
	}
	fmt.Printf("总行数：%d\n", len(lineArr))
	for _, lineStr := range lineArr { //遍历
		lineCul := strings.Split(lineStr, "\"")
		if len(lineCul) < 2 {
			continue
		}
		for _, v := range domainList { //该行是否能匹配上指定的几个域名，满足一个即可
			if strings.Contains(lineCul[0], v) {
				//截取接口名称
				uri := GetUri(lineCul[1])
				//加入pvmap
				CaculatePv(uri, 1)
				//一旦匹配上，不需要继续遍历，跳过
				break
			}
		}
	}
	return nil
}

func GetUri(target string) string {
	//arrList := strings.Split(lineData,"\"")
	//if len(arrList)>2{
	//	target := arrList[1]
	//
	//}
	if !strings.HasPrefix(target, "/app") {
		return target
	}

	sli := strings.Split(target, "/")
	if len(sli) < 5 {
		return ""
	}

	uri := "/" + strings.Join(sli[5:], "/")
	//for _, v := range sli[5:] {
	//	uri = uri + v + "/"
	//}
	//uri = strings.TrimSuffix(uri, "/")
	return uri
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

func ConsumerDemo(id int, ch chan int, done chan bool) {
	for {
		value, ok := <-ch
		if ok {
			fmt.Printf("id: %d, recv: %d\n", id, value)
		} else {
			fmt.Printf("id: %d, closed\n", id)
			break
		}
	}
	done <- true
}

func Consumer(wg *sync.WaitGroup, fileChan chan string, domainList []string, isGzip bool) {
	defer wg.Done()
	for {
		//先判断file队列为空，则goroutine退出。(如果有生产者，且后续还往里放数据，则不是break,而是sleep. if else)
		filePath, ok := <-fileChan
		if !ok {
			break
		} else {
			//从file队列里取文件，然后计算pv。
			err := StatFilePV(domainList, filePath, isGzip)
			if err != nil {
				fmt.Printf("err:%s\n", err.Error())
				return
			}
		}
	}
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

//
//func CaculatePv2(uri string, num int64) {
//	GlobalStatPvData.lc.Lock()
//	v, exist := GlobalStatPvData.PvData[uri]
//	if exist {
//		GlobalStatPvData.PvData[uri] = v + num
//	} else {
//		GlobalStatPvData.PvData[uri] = num
//	}
//	GlobalStatPvData.lc.Unlock()
//}
