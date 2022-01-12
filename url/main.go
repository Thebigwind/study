package main

import (
	"fmt"
	"net/url"
)

func main() {
	var uri url.URL
	q := uri.Query()
	q.Add("name", "张三")
	q.Add("age", "20")
	q.Add("sex", "1")
	queryStr := q.Encode()
	fmt.Println(queryStr)
}

///127.0.0.1:8800/user/checkusersession?client=web&infoLevel=1&reqPath=%252Fwallet%252Frecord%253FcreatedAtStart%253D2021-11-01%2526userId%253D70000972252383%2526length%253D20%2526clientFrom%253Dweb%2526role%253Dxzfd%2526clientVersion%253D500%2526createdAtEnd%253D2022-01-05%2526walletId%253D132343779557690%2526xztoken%253DWyIwMDA2MDEyOTA1SG45MSIseyJ1c2VyaWQiOiI3MDAwMDk3MjI1MjM4MyIsImV4cGlyZSI6MTY0Mzk3MDU0MCwiYyI6IndlYiJ9LCI4MWE1OGVhNGJiYWE3MWRhMjEzYjRmODJiZjAzODUwMSJd%2526category%253D%2526sessId%253De864b98fafaf9858d543a2e1397b3590%2526needStatistics%253D1%2526offset%253D0%2526type%253D0&servName=test-wirelesspub-payment.xiaozhu.com&sessId=e864b98fafaf9858d543a2e1397b3590&userId=70000972252383
