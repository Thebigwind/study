package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main(){
	//读取输入
	reader := bufio.NewReader(os.Stdin)
	str,_ := reader.ReadString('\n')
	//var str string
	//fmt.Scanln(&str)
	str = strings.Replace(str,"\n", " ", -1)

	//读取输入
	var str2 string
	fmt.Scanln(&str2)
	str2 = strings.Replace(str2,"\n"," ",-1)

	fmt.Printf("str:%s\n",str)
	fmt.Printf("str2:%s\n",str2)
	//统计
	num := 0
	for _,v := range str{
		if strings.ToLower(string(v)) == strings.ToLower(str2) {
			num++
		}
	}
	fmt.Println(num)

}