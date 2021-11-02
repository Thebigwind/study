package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GetKey() (timestamp string, sign string, err error) {
	cmd := exec.Command("/bin/python", "/application/scripts/hezhong_host_status/aaa.py")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error:can not obtain stdout pipe for command:%s \n", err)
		return timestamp, sign, err
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err ", err)
		return timestamp, sign, err
	}
	var Key []string
	//读取所有输出
	bytes := bufio.NewReader(stdout)
	for {
		line, err := bytes.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Read Err:", err)
		}
		Key = append(Key, line)
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Wait", err.Error())
		return timestamp, sign, err
	}
	timestamp = Key[0]
	sign = Key[1]
	return timestamp, sign, nil
}
func SendDingMsg(msg string) {
	//请求地址模板
	//timestamp , sign , err := GetKey()
	//if err != nil {
	//	fmt.Println("Get Key Err:",err)
	//}
	//timestamp = strings.TrimRight(timestamp, "\n")
	//sign = strings.TrimRight(sign, "\n")
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	//timestamp := int64(1572870805748)
	sign := "SECd252d7f77a744391ef01a5431f3843376cbeb7e274d115d404ef29e685803d22"
	webHook := "https://oapi.dingtalk.com/robot/send?access_token=42fb167aa377daef9cc78d3c89b441c75b3735abb7e9cad29da1ed7f4386fda8&" + "timestamp=" + timestamp + "&sign=" + sign
	content := `{"msgtype": "text",
   "text": {"content": "` + msg + `"},
        "at": {
           "atMobiles": [
             "18301371817"
           ],
           "isAtAll": true
        }
  }`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	//设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-agent", "firefox")
	//发送请求
	resp, err := client.Do(req)
	//关闭请求
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		fmt.Println("handle error")
	}
}
func main() {
	SendDingMsg("[犯罪记录]校验结果通知xxxxxxxx")
}
