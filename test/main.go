package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

//=2409d6c435b16e09a63d4d9a1f609c48

func Str2Md5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}

func GetSessId4Html5(uid string, business string) string {
	prefix := "password#"
	if uid == "10828400" {
		return Str2Md5(prefix + uid + "salt4html5force")
	} else if business == "api" {
		//一天过期
		fmt.Printf("time:%s\n", time.Now().Format("2006-01-02"))
		return Str2Md5(prefix + uid + "salt4html5" + time.Now().Format("2006-01-02"))
	} else {
		return Str2Md5(prefix + uid + "salt4html5")
	}
}

func main() {
	fmt.Printf("result:%s\n", GetSessId4Html5("2070657982", "api"))

	fmt.Printf("xxx:%d\n", 15552000/3600/24)
	fmt.Printf("oooo:%d\n", 2592000/3600/24)
}
