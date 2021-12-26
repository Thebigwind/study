package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	test5()
}

func test() {
	//str := `101.70.18.242 wirelesspub-general.xiaozhu.com 140.210.66.251 [17/Dec/2021:00:00:00 +0800] "/remind/getordervoice" "GET /remind/getordervoice?clientFrom=web&clientVersion=500&role=xzfd&gatets=1639670387&userId=26041036801&sessId=b77dbfcf812d3be59fb334355d3dc38b&xztoken=WyIyNzA4MTExOTI5OUpqNyIseyJ1c2VyaWQiOiIyNjA0MTAzNjgwMSIsImV4cGlyZSI6MTYzOTM1NDc2NywiYyI6IndlYiJ9LCJiYjFlNTBmZDg0ZDg0NWNjMzMwMjAwYWM1OTc2MGY2ZCJd&gatesign=ed521ab611081381ee22c4d92ce26733412da9da HTTP/2.0" 200 237 "-" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36" "0.172" 10.4.13.203:80 "0.172" - "-" "-"}`
	str2 := `115.195.85.248 securewireless.xiaozhu.com 140.210.66.251 [17/Dec/2021:00:00:00 +0800] "/app/xzfk/android/6.32.00/my/NavCount" "GET /app/xzfk/android/6.32.00/my/NavCount?anonymous_id=eb7742ec2fcb32f5&dispathChannel=vivo&gatets=1639670399&sessId=WyIxNDAxMDgwOTIxeHpMNSIseyJzc0lkIjoxMDgwODU1MjI3MjY5MTUsInNzVHlwZSI6Im1vYmlsZV9jb2RlIiwiZGF0YSI6IjYzMjE4NWJlMzkyMTVjNjAxYmI2OGRmZDhhYjM2NTY4IiwiZXhwaXJlIjoxNjQ1MDc0NTU0fSwiZmI4ZmJjMGJkZGRiMDM2YmE1NzI2ZDE5MzllM2M5NjYiXQ%3D%3D&uniqueId=b0d64e91b18b17b8f64a1f5ca23001ba1713cddda5feddcc04d7f8d262bcc3de&userId=14236839401&gatesign=0a8ceaa0f7d1e1ff320ebfcb8f8d59f89f091a8c HTTP/2.0" 200 342 "-" "XZTenant/6.32.00 (android; android10; Scale/2.00) XZApp/android(android; Phone; android10) XZAppChannel/vivo XZAppVersion/6.32.00 XZNetType/WIFI" "0.023" 10.4.12.131:80 "0.023" - "-" "-"`
	arr := strings.Split(str2, "\"")
	for _, v := range arr {
		fmt.Printf("v:%s\n", v)
	}
}

func test2() {
	pwd, _ := os.Getwd()
	//获取文件或目录相关信息
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(fileInfoList))
	for i := range fileInfoList {
		fmt.Println(fileInfoList[i].Name()) //打印当前文件或目录下的文件或目录名
	}
}

func test3() {
	a := "aaa"
	domainList := strings.Split(a, ",")
	domainKey := ""
	if len(domainList) == 1 {
		domainKey = a
	} else {
		domainKey = strings.Join(domainList, "\\|")
	}
	fmt.Printf("domainKey:%s\n", domainKey)
}

func test4() {
	err := os.MkdirAll("dir1/dir2/dir3", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("success:%s", "a")
	}
}

func test5() {
	aa := fmt.Sprintf("%.2d", 0)
	fmt.Println(aa)
}
