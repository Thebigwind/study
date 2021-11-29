package main

import "fmt"

func main() {
	code := "001RMnFa1CHicC0X3HHa116Z8v1RMnF3"
	wxInfo, err := LoginWXSmall(code)
	if err != nil {
		fmt.Printf("LoginWXSmall err:%s",err.Error())
		//return
	}
	fmt.Printf("wxInfo:%v",wxInfo)
	sessionKey := "YUuU7qjh5cJ0pFZKPqdzvg=="

	encryptData := "RHJlWWdiSjVFTmVjZFZsdGRYUVNRd290UHdpUWtJem9LdlNHc2tpaUdSY3NRMmNnZmY3ZGR3U3FuUUEvU3p6d3QvNmsxbkF5bWpmZG9Zd29lMk5vaHNIUGZnelkyTHUxWWNNSjRwQzhnTFlVTlZmNVcxakptcVM5NjdDQlFFV1FQbDR6bnpmNlVpME11UmVua05OTDE2dHF5Qnk4TkVQa1VkYmgxbGgwengrQS8vYTJwL0R1UytBMmJ6eXZzQ3MwUE5jMmRNcDBxbTlvTWVzdVhGaDVlZz09"
	//encryptData := "6Ebfb0UsrK2DwXP2VaKnvcpW4m/d1AeNclsjYg+JNZylJLwbk0VvmkugidjVVE+Kc0siY2qNZhxRU21iMRi7HudWDbhdDd8nHUHzljjowPokd4uF7KL1eSQPNXb21LanVZLXjZQw6RfRN39tIn/bnQe2/2Yn1jxRPNf1+tmQZ/2pboJDgdCY620BYo9yD5eI5Ot9Opveov+aslSoxFttKw=="
	iv := "nG9z27AERfB8jKuH8NNLhg==" // bGpacFpyc1NIQ2EraTZkbDZlQW9WZz09
	m,err := DecryptWXOpenData(sessionKey,encryptData, iv)
	if err != nil{
		fmt.Printf("DecryptWXOpenData err:%s",err.Error())
		return
	}
	fmt.Printf("m:%+",m)
}
