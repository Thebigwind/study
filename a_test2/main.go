
package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main() {
	//[[3,5],[1,2],[6,8]]
	str:=""
	//record := make([]int,0)
	n, _ := fmt.Scan(&str)
	if n == 0 {
		return
	}
	str = str[1:len(str)-1]
	//fmt.Println(str)

	strArr := strings.Split(str, "],[")
	//fmt.Println(strArr)
	strArr[0] = strings.Trim(strArr[0], "[")
	strArr[len(strArr)-1] = strings.Trim(strArr[len(strArr)-1], "]")
	//fmt.Println(strArr)

	//m := make(map[int]bool)
	//predata1 := 0
	//predata2 := 0
	book1Max := 0
	book1Min := 0

	book2Max := 0
	book2Min := 0

	book3Max := 0
	book3Min := 0

	book4Max := 0
	book4Min := 0

	for i,v := range strArr{
		dataArr := strings.Split(v,",")
		data1,_ := strconv.Atoi(dataArr[0])
		data2,_ := strconv.Atoi(dataArr[1])

		minData := 0
		maxData := 0
		if data1>data2{
			maxData = data1
			minData = data2
			//m[data2] = true
		}else{
			maxData = data2
			minData = data1
			//m[data1] = true
		}
		switch i{
		case 0:
			book1Max = maxData
			book1Min = minData
		case 1:
			book2Max = maxData
			book2Min = minData
		case 2:
			book3Max = maxData
			book3Min = minData
		case 3:
			book4Max = maxData
			book4Min = minData
		}
	}

	//
	sum := 4
	if book1Min==book2Max || book1Min==book2Min || book1Max==book2Max || book1Max==book2Min{
		sum--
	}
	if book1Min==book3Max || book1Min==book3Min || book1Max==book3Max || book1Max==book3Min{
		sum--
	}
	if book1Min==book4Max || book1Min==book4Min || book1Max==book4Max || book1Max==book4Min{
		sum--
	}
	if !(book1Min==book2Max || book1Min==book2Min || book1Max==book2Max || book1Max==book2Min){
		if book2Min==book3Max || book2Min==book3Min || book2Max==book3Max || book2Max==book3Min{
			sum--
		}
	}
	if !(book1Min==book2Max || book1Min==book2Min || book1Max==book2Max || book1Max==book2Min) && !(book2Min==book3Max || book2Min==book3Min || book2Max==book3Max || book2Max==book3Min) {
		if book3Min == book4Max || book3Min == book4Min || book3Max == book4Max || book3Max == book4Min {
			sum--
		}
	}
		if sum<0{
			sum = 0
		}
		fmt.Println(sum)
	}

