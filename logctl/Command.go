package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strings"
)

func Command() {
	download := kingpin.Command("download", "下载命令.")
	download_log := download.Command("accesslog", "下载访问日志(当前仅支持访问日志).")
	download_log_type := download_log.Flag("type", "1:waiwang-lb, 2:neiwang-lb").
		Required().
		Int()
	download_log_date := download_log.Flag("date", "访问日志的日期,格式：20211225").
		Required().
		String()
	download_log_hour := download_log.Flag("hour", "访问日志的小时，24小时制，格式如：15，不指定默认为空，表示24个小时").
		String()
	download_log_path := download_log.Flag("path", "下载日志文件到目标目录.").
		Required().
		Short('p').
		String()

	stat := kingpin.Command("stat", "统计命令.")
	stat_pv := stat.Command("pv", "统计pv.")
	stat_pv_domainkey := stat_pv.Flag("domain", "域名关键字，多个已逗号分割（当前是外网日志，包含了多个服务）").
		Required().
		Short('d').
		String()
	stat_pv_filepath := stat_pv.Flag("path", "日志文件所在目录.").
		Required().
		Short('p').
		String()
	stat_pv_goroutinenum := stat_pv.Flag("num", "携程数.").
		Short('n').
		Int()
	stat_pv_csv := stat_pv.Flag("csv", "csv文件路径,结果文件.").
		Required().
		String()
	stat_pv_servicename := stat_pv.Flag("service", "服务名.").
		String()
	stat_pv_isgzip := stat_pv.Flag("isgzip", "是否是压缩文件.").
		Bool()
	stat_pv_type := stat_pv.Flag("type", "分析方式：1：程序方式；2：命令方式(不支持多线程).").
		Required().
		Int()

	kingpin.CommandLine.HelpFlag.Short('h')
	osParse := kingpin.Parse()
	cmds := strings.Split(osParse, " ")
	switch cmds[0] {

	case "download":
		switch cmds[1] {

		case "accesslog":
			params := AccessLogDownloadParams{
				Type: *download_log_type,
				Date: strings.Trim(*download_log_date, " "),
				Hour: strings.Trim(*download_log_hour, " "),
				Path: strings.Trim(*download_log_path, " "),
			}
			DoDownloadLog(params)
		}

	case "stat":

		switch cmds[1] {

		case "pv":
			params := AccessLogPvParams{
				LogDirPath:   *stat_pv_filepath,
				DomainKey:    strings.Trim(*stat_pv_domainkey, " "),
				GoroutineNum: *stat_pv_goroutinenum,
				IsGzip:       *stat_pv_isgzip,
				CsvName:      strings.Trim(*stat_pv_csv, " "),
				ServiceName:  strings.Trim(*stat_pv_servicename, " "),
			}

			if *stat_pv_type == 1 {
				if err := DoAnalyzePv(params); err != nil {
					os.Exit(1)
				}
			} else if *stat_pv_type == 2 {
				if err := DoAnalyzePVCommand(params); err != nil {
					os.Exit(1)
				}
			}

			return
		}
	}
}
