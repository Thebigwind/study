package str

func Test1() bool {
	var v string
	if v == "" {
		return true
	}
	return false
}

func Test2() bool {
	var v string
	if len(v) == 0 {
		return true
	}
	return false
}

/*
无论是 len(v) == 0，又或是 v == "" 的判断，其编译出来的汇编代码都是完全一致的。可以明确 Go 编译器在这块做了明确的优化，大概率是直接比对了。
*/
