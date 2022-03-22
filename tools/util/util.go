package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

func ContainInt(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func ContainInt64(arr []int64, target int64) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func ContainString(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func MapStrToStr(arr []string, fn func(s string) string) []string {
	var newArray = []string{}
	for _, it := range arr {
		newArray = append(newArray, fn(it))
	}
	return newArray
}

func IfEmpty(src string, dst string) string {
	if dst == "" {
		return src
	}
	return dst
}

func DeleteRepeat(list []string) []string {
	mapdata := make(map[string]interface{})
	if len(list) <= 0 {
		return nil
	}
	for _, v := range list {
		mapdata[v] = "true"
	}
	var datas []string
	for k, _ := range mapdata {
		if k == "" {
			continue
		}
		datas = append(datas, k)
	}
	return datas
}

func RemoveSliceMap(a []interface{}) (ret []interface{}) {
	n := len(a)
	for i := 0; i < n; i++ {
		state := false
		for j := i + 1; j < n; j++ {
			if j > 0 && reflect.DeepEqual(a[i], a[j]) {
				state = true
				break
			}
		}
		if !state {
			ret = append(ret, a[i])
		}
	}
	return
}

func InterfaceToString(arg interface{}) string {
	switch arg := arg.(type) {
	case int64:
		return strconv.FormatInt(arg, 10)
	case string:
		return arg
	case bool:
		return strconv.FormatBool(arg)
	default:
		data, _ := json.Marshal(arg)
		return string(data)
	}
}

//只需要共享底层 Data 和 Len 就可以实现 zero-copy。原理上是利用指针的强转，代
func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}
func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func isMobile(mobile string) {
	result, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	if result {
		println(`正确的手机号`)
	} else {
		println(`错误的手机号`)
	}
}

func Str1() {
	str := fmt.Sprintf("%s %s %s", "format", "string", "by fmt.Sprintf")
	fmt.Printf(str)
}

func Str2() {
	str := fmt.Sprintf("%s %s %s", "format", "string", "by fmt.Sprintf")
	fmt.Printf(str)
}

//golang的 strings 包为字符串的拼接提供了一个方法func Join(a []string, sep string) string, Join的内部实现比fmt.Sprintf要简单的多,思路就是: Join会先根据字符串数组的内容，计算出一个拼接之后的长度，然后申请对应大小的内存，一个一个字符串填入.代码如下:
// Join 将传如的字符串连接成一个字符串
func Join(a []string, sep string) string {
	// 如果字符串数量少,直接使用运算符拼接
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		return a[0] + sep + a[1]
	case 3:
		return a[0] + sep + a[1] + sep + a[2]
	}

	// 计算最终字符串的字符大小
	// 首先计算连接符sep的大小
	n := len(sep) * (len(a) - 1)

	// 计算被连接的字符串的字符数
	for _, value := range a {
		n += len(value)
	}

	// 知道了总的字符数量,创建对应大小的数组
	b := make([]byte, n)
	bp := copy(b, a[0])

	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}

	return string(b)
}

/*
bytes包中的Buffer提供了一个方法 func (b *Buffer) WriteString(s string) (n int, err error)

WriteString将s的内容追加到缓冲区，并根据需要增加缓冲区。返回值n为s的长度;err总是nil。

如果缓冲区太大，WriteString将会因为ErrTooLarge而陷入恐慌。
*/
func StringBuffer() {
	// 声明一个Buffer
	var buf bytes.Buffer
	buf.WriteString("good ")
	buf.WriteString("boy!")
	fmt.Println(buf.String()) // good boy!
}

/**
strings.Builder 内部通过 slice 来保存和管理内容。slice 内部则是通过一个指针指向实际保存内容的数组。

strings.Builder 同样也提供了 Grow() 来支持预定义容量。

当我们可以预定义我们需要使用的容量时，strings.Builder 就能避免扩容而创建新的 slice 了。strings.Builder是非线程安全，性能上和 bytes.Buffer 相差无几。
*/
func StringBuilder() {
	// 声明一个Buffer
	var buf strings.Builder
	buf.WriteString("good ")
	buf.WriteString("boy!")
	fmt.Println(buf.String()) // good boy!
}
