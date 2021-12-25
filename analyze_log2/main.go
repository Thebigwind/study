package main

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

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

func main() {
	NewStatPvMap()
	start := time.Now().UnixNano()
	Command()
	end := time.Now().UnixNano()
	fmt.Printf("耗时：%d ms\n", (end-start)/1e6)
}
func Command() {
	stat := kingpin.Command("stat", "统计命令.")
	stat_pv := stat.Command("pv", "统计pv.")
	stat_pv_key := stat_pv.Flag("domain", "域名关键字，多个已逗号分割（当前是外网日志，包含了多个服务）").
		Short('d').
		String()
	stat_pv_filepath := stat_pv.Flag("path", "文件路径.").
		Short('p').
		String()
	stat_pv_goroutinenum := stat_pv.Flag("num", "携程数.").
		Short('n').
		Int()
	stat_pv_csv := stat_pv.Flag("csv", "csv文件路径,结果文件.").
		String()
	stat_pv_servicename := stat_pv.Flag("service", "服务名.").
		String()
	stat_pv_isgzip := stat_pv.Flag("isgzip", "是否是压缩文件.").
		Bool()
	kingpin.CommandLine.HelpFlag.Short('h')
	osParse := kingpin.Parse()
	cmds := strings.Split(osParse, " ")
	switch cmds[0] {

	case "stat":

		switch cmds[1] {

		case "pv":

			dirPath := *stat_pv_filepath
			domainKey := *stat_pv_key
			//服务的几个域名关键字
			domainList := strings.Split(domainKey, ",")
			//根据日志目录获取日志文件列表
			filePahtList, err := GetPathList(dirPath)
			if err != nil {
				return
			}
			if *stat_pv_goroutinenum <= 1 {
				//遍历文件列表，按文件统计
				for i, filePath := range filePahtList {
					fmt.Printf("i:%d,filepath:%s\n", i, filePath)
					if err := StatFilePV(domainList, filePath, *stat_pv_isgzip); err != nil {
						fmt.Printf("err:%s\n", err.Error())
						return
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
				wg.Add(*stat_pv_goroutinenum)
				for i := 0; i < *stat_pv_goroutinenum; i++ {
					go Consumer(&wg, fileChan, domainList, *stat_pv_isgzip)
				}
				wg.Wait()
			}
			//写csv文件
			csvPath := ""
			if *stat_pv_csv == "" {
				csvPath = "default-" + strconv.FormatInt(time.Now().Unix(), 10) + "-.csv"
			}
			WriteCsv(csvPath, *stat_pv_servicename)
			return

		}
	}
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

//获取日志文件路径列表
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
		return strings.Split(bytes2str(s2), "\n"), nil
	} else {
		return strings.Split(bytes2str(s1), "\n"), nil
	}

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
				CaculatePv(uri)
				//一旦匹配上，不需要继续遍历，跳过
				break
			}
		}
	}
	return nil
}

func ContainString(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

//Todo
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

	uri := "/"
	for _, v := range sli[5:] {
		uri = uri + v + "/"
	}
	uri = strings.TrimSuffix(uri, "/")
	return uri
}

func CaculatePv(uri string) {
	GlobalStatPvData.lc.Lock()
	v, exist := GlobalStatPvData.PvData[uri]
	if exist {
		GlobalStatPvData.PvData[uri] = v + 1
	} else {
		GlobalStatPvData.PvData[uri] = 1
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

func ExecShell(s string) (error, string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("exec_shell:\n", err.Error())
		return err, ""
	}
	//fmt.Printf("%s", out.String())
	return err, strings.Trim(out.String(), "\n")
}
