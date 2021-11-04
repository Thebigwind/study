package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func main() {
	//url := "http://prerelease-service-uic.xiaozhu.com/backGroundCheck/identityCheckToBatch"

	url := "http://127.0.0.1:8800/backgroundcheck/identitychecktobatch"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("nameAndCard", "[{\"card\":\"110225196403026127\",\"name\":\"\\u5d14\\u79c0\\u829d\"}]")
	_ = writer.WriteField("businessScene", "submitOrder")
	_ = writer.WriteField("userId", "154874068097")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "userInfo=%7B%22xztoken%22%3A%22WyIwMjA0MTA0NjA5a3lkUSIseyJ1c2VyaWQiOiIyMDcwNjU3OTgyIiwiZXhwaXJlIjoxNjM2MzYxMTYyLCJjIjoid2ViIn0sIjgxZDQ4ODE2N2U2ODA4OTliZWI0YzM3NjQzNDFjY2MyIl0%3D%22%2C%22sessId%22%3A%225fab6aece6cd6c688edc1d3a1b779d91%22%2C%22imToken%22%3A%22f56d32742731a66d5a6a050d40a283db%22%2C%22isPersonalLandlord%22%3Atrue%2C%22isOrgLandlord%22%3Afalse%2C%22isLandlord%22%3Atrue%2C%22createTime%22%3A%222016-01-11+02%3A21%3A20%22%2C%22isNew%22%3Afalse%2C%22userName%22%3A%22%5Cu4e00%5Cu4e8c%5Cu4e09%5Cu56db%5Cu4e94%5Cu516d%5Cu4e03%5Cu516b%5Cu4e5d%5Cu5341JQK%22%2C%22realName%22%3A%22%5Cu4efb%5Cu540d%5Cu626c%22%2C%22mobile%22%3A%2218515630238%22%2C%22mobileNation%22%3A0%2C%22nationCode%22%3A%2286%22%2C%22nationName%22%3A%22CN%22%2C%22email%22%3A%22myphper%40foxmail.com%22%2C%22isActiveEmail%22%3Atrue%2C%22zhimaScore%22%3A0%2C%22xiaoBaiScore%22%3A%220.0%22%2C%22userId%22%3A%222070657982%22%2C%22nickName%22%3A%22%5Cu4e00%5Cu4e8c%5Cu4e09%5Cu56db%5Cu4e94%5Cu516d%5Cu4e03%5Cu516b%5Cu4e5d%5Cu5341JQK%22%2C%22isDefaultHeadImg%22%3Afalse%2C%22realVery%22%3Atrue%2C%22busiTravelUserIdentify%22%3A%22%22%2C%22xzCredit%22%3Atrue%2C%22headImageId%22%3A%22114242694938948%22%2C%22headImgurl%22%3A%22https%3A%5C%2F%5C%2Fimage.xiaozhustatic1.com%5C%2F00%2C100%2C100%2C1%2C80%2C1%5C%2Fs%2C5%2C1sIXf%2C132%2C132%2C2%2Cb5d93091.jpg%22%2C%22bindLandlordStr%22%3A%22xzdz007%22%2C%22showRegSuccessPage%22%3Afalse%2C%22kaTag%22%3Afalse%7D")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
