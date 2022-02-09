package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	//要压缩成一个zip的多个文件的路径
	files := []string{"/tmp/aa.json", "/tmp/bb.json", "/tmp/cc.json"}
	now := strconv.FormatInt(time.Now().UnixNano(), 10)
	//设置输出的zip的路径
	output := "/tmp/" + now + ".zip"
	if err := ZipFiles(output, files); err != nil {
		panic(err)
	}
}

func ZipFiles(filename string, files []string) error {
	fmt.Println("start zip file......")
	//创建输出文件目录
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()
	//创建空的zip档案，可以理解为打开zip文件，准备写入
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()
	// Add files to zip
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {
	//打开要压缩的文件
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()
	//获取文件的描述
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}
	//FileInfoHeader返回一个根据fi填写了部分字段的Header，可以理解成是将fileinfo转换成zip格式的文件信息
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = filename
	/*
	   预定义压缩算法。
	   archive/zip包中预定义的有两种压缩方式。一个是仅把文件写入到zip中。不做压缩。一种是压缩文件然后写入到zip中。默认的Store模式。就是只保存不压缩的模式。
	   Store   unit16 = 0  //仅存储文件
	   Deflate unit16 = 8  //压缩文件
	*/
	header.Method = zip.Deflate
	//创建压缩包头部信息
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	//将源复制到目标，将fileToZip 写入writer   是按默认的缓冲区32k循环操作的，不会将内容一次性全写入内存中,这样就能解决大文件的问题
	_, err = io.Copy(writer, fileToZip)
	return err
}
