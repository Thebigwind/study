package util

import (
	"encoding/json"
	"reflect"
	"regexp"
	"strconv"
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
