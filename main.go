package main

import (
	"errors"
	"fmt"
	valid "github.com/guanguans/id-validator"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var BizType interface{}

func main() {

	if time.Now().Format("20060102")[4:] > "0930" {
		fmt.Println("xxxx")
	}
	fmt.Println(time.Now().Format("20060102")[4:])
	fmt.Println(time.Now().AddDate(1, 0, 0))

	//BizType = "1"
	if BizType != nil && BizType.(string) != "2" {
		fmt.Printf("哈哈哈哈\n")
		fmt.Printf("xx:%v", BizType)
		os.Exit(0)
	}

	fmt.Println(time.Now().Format("20060102"))
	fmt.Println(VerifyMobileFormat("17910715315"))
	isMobile2("27710715315")
	fmt.Printf("chang：", len("陆飞\n "))
	a := strings.Trim("陆飞\n ,", "\n ,\t")
	fmt.Printf("chang：", len(a))
	fmt.Printf("chang：", a)

	age, _ := caclu("2003-09-30")
	fmt.Println("---")
	fmt.Println(age)

	idinfo, err := valid.GetInfo("21038119940829591X", false)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf("idinfo:%+v", idinfo)
	fmt.Println()
	fmt.Printf("idinfo sex:%+v", idinfo.Sex)

	Test2()
}
func isMobile(mobile string) {
	result, _ := regexp.MatchString(`^(1[3|4|5|7|8][0-9]\d{4,8})$`, mobile)
	if result {
		println(`正确的手机号`)
	} else {
		println(`错误的手机号`)
	}
}
func isMobile2(mobile string) {
	result, _ := regexp.MatchString(`^(1\d{2,11})$`, mobile)
	if result {
		println(`正确的手机号`)
	} else {
		println(`错误的手机号`)
	}
}

func caclu(birth string) (int, error) {
	birthday := strings.Split(birth, "-")

	if len(birthday) < 3 {
		return 0, errors.New("出生日期格式解析错误")
	}

	birYear, _ := strconv.Atoi(birthday[0])
	birMonth, _ := strconv.Atoi(birthday[1])
	day, _ := strconv.Atoi(birthday[2])

	age := time.Now().Year() - birYear

	if int(time.Now().Month()) < birMonth {
		age--
	}
	if time.Now().Day() < day {
		age--
	}
	//fmt.Println("month:",time.Now().Month())
	//fmt.Println("day:",time.Now().Day())
	return age, nil
}

func VerifyMobileFormat(mobileNum string) bool {

	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func Test2() {
	birthday := []string{"1990", "12", "12"}
	birYear, _ := strconv.Atoi(birthday[0])
	birMonth, _ := strconv.Atoi(birthday[1])
	day, _ := strconv.Atoi(birthday[2])

	age := time.Now().Year() - birYear
	fmt.Printf("age0:%d\n", age)
	if int(time.Now().Month()) < birMonth {
		age--
	}
	fmt.Printf("age1:%d\n", age)
	//fmt.Printf("age1:%d\n",age)
	if int(time.Now().Month()) == birMonth && time.Now().Day() < day {
		age--
	}
	fmt.Println("----------")
	fmt.Printf("age:%d", age)
}

func test() {
	var a string
	a = "abcd"
	fmt.Printf("aa:%+v", time.Now().Year())
	fmt.Println(a)
}
