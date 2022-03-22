package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println("xxxxx")
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

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ { //
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Printf("result:%+v", arr)
}

func InsertSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[i] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}

	fmt.Printf("result:%+v", arr)

}

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1
	pivotIndex := rand.Int() % len(a) // Pick a pivot

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}
