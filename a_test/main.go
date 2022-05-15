package main

import (
	"fmt"
)
func main() {
	a:=""
	b:=""

	n, _ := fmt.Scan(&a,&b)
	if n == 0 {
		return
	}
	indexRecord := make([]int,0)
	for _,av := range []rune(a){
		flag := false
		for bi,bv:= range []rune(b){
			if string(av) == string(bv){
				flag = true
				//记录当前下标位置，并继续向下遍历
				indexRecord = append(indexRecord,bi)
			}
		}
		if !flag{
			//不存在，返回-1
			fmt.Println(-1)
			return
		}
	}
	lengthA := len(a)
	if len(indexRecord)<lengthA{
		//不存在，返回-1
		fmt.Println(-1)
		return
	}
	//indexRecord排序，
	fmt.Println(indexRecord)
	indexRecord = qsort(indexRecord)
	fmt.Println(indexRecord)
	for i:= lengthA-1;i<len(indexRecord);i++{
		if indexRecord[i] - indexRecord[i-(lengthA-1)] == lengthA - 1{
			fmt.Println(indexRecord[i-(lengthA-1)])
			return
		}
	}
	fmt.Println(-1)
	return

	// 12355689

}

func qsort(a []int)[]int{
	if len(a)<2{
		return a
	}
	left,right := 0,len(a)-1
	for i := range a{
		if a[i] < a[right]{
			a[left],a[i] = a[i],a[left]
			left++
		}
	}
	a[left],a[right] = a[right],a[left]
	qsort(a[:left])
	qsort(a[left+1:])
	return a
}
