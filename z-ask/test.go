package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	a := CompareVersion("", "6.33.23")
	fmt.Printf("a:%v", a)
	fmt.Printf("-----\n")
	fmt.Println(InterfaceToString(3232325125245113434))
}

func CompareVersion(version1, version2 string) bool {
	sli1 := strings.Split(version1, ".")
	sli2 := strings.Split(version2, ".")
	maxLen := int(math.Max(float64(len(sli1)), float64(len(sli2))))
	for {
		if len(sli1) == maxLen {
			break
		}
		sli1 = append(sli1, "0")
	}
	for {
		if len(sli2) == maxLen {
			break
		}
		sli2 = append(sli2, "0")
	}
	for i := 0; i < maxLen; i++ {
		num1, _ := strconv.Atoi(sli1[i])
		num2, _ := strconv.Atoi(sli2[i])
		if num1 > num2 {
			return true
		} else if num1 < num2 {
			return false
		}
	}
	return false
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
