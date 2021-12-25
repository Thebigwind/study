package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func GetPath(path string) (error, string) {
	fullPath := ""
	var err error
	if strings.HasPrefix(path, "/") { //绝对路经
		return nil, path
	} else if path == "" { //当前路经
		fullPath, _ = os.Getwd()
	} else { //
		err, fullPath = JoinPath(path)
	}

	return err, fullPath
}

func JoinPath(val string) (error, string) {
	var err error
	fullPath := ""
	//relative path
	envPath, _ := os.Getwd()
	if envPath != "/" {
		envPath = strings.TrimRight(envPath, "/")
		if !strings.HasPrefix(val, "..") {
			if strings.HasPrefix(val, ".") {
				valList := strings.Split(val, "/")
				valNew := strings.Join(valList[1:len(valList)], "/")
				fullPath = envPath + "/" + valNew
			} else {
				fullPath = envPath + "/" + val
			}
		} else {
			dirList := strings.Split(envPath, "/")
			valList := strings.Split(val, "/")

			valAheadLen := 0
			valBack := ""
			newPath := ""
			newPathLen := 0

			//calcu valAhead.lentgh
			for i := 0; i < len(valList); i++ {
				if valList[i] == ".." {
					valAheadLen += 1
				} else {
					i = len(valList) //exit for
				}
			}
			//join valBack
			valBack = strings.Join(valList[valAheadLen:len(valList)], "/")

			if len(dirList) >= valAheadLen {
				//join newPath
				newPathLen = len(dirList) - valAheadLen
				newPath = strings.Join(dirList[0:newPathLen], "/")
				fullPath = newPath + "/" + valBack

				if fullPath != "/" && strings.HasSuffix(fullPath, "/") {
					fullPath = strings.TrimRight(fullPath, "/")
				}
			} else {
				err = errors.New("Already at /")
				return err, ""
			}
		}
	} else { //if pwd in "/"
		if strings.HasPrefix(val, "..") {
			err = errors.New("Already at /")
			return err, ""
		} else {
			fullPath = "/" + val
		}
	}
	return err, fullPath
}
