package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"strconv"
	"strings"
	"sync"
	"time"
)

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
