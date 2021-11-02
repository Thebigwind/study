package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// or you can use []byte(`...`) and convert to Buffer later on
	body := "<ValidateRQ>\n  <AuthenticationToken>\n    <Username>556471e1</Username>\n    <Password>0915350284be0668</Password>\n    <CreateToken>22251178182015010620150107497867981843210904377</CreateToken>\n  </AuthenticationToken>\n  <TaoBaoHotelId>13577578181</TaoBaoHotelId>\n  <HotelId>74551593009156</HotelId>\n  <TaoBaoRoomTypeId>5501264818</TaoBaoRoomTypeId>\n  <RoomTypeId>75480776835156</RoomTypeId>\n  <TaoBaoRatePlanId>4978679818</TaoBaoRatePlanId>\n  <RatePlanCode>75513705005058</RatePlanCode>\n  <TaoBaoGid>3824371818</TaoBaoGid>\n  <CheckIn>2021-09-10</CheckIn>\n  <CheckOut>2021-09-11</CheckOut>\n  <RoomNum>1</RoomNum>\n  <PaymentType>1</PaymentType>\n  <Extensions>{\"searchid\":\"22251178182015010620150107497867981843210904377\"}  </Extensions>\n  <TotalPrice>7949</TotalPrice>\n  <DailyInfos>\n    <DailyInfo>\n      <Day>2019-12-30</Day>\n      <Price>7949</Price>\n      <BreakFast>2</BreakFast>\n     </DailyInfo>\n  </DailyInfos>\n  <CurrencyCode>CNY</CurrencyCode>\n</ValidateRQ>"

	client := &http.Client{}
	// build a new request, but not doing the POST yet
	req, err := http.NewRequest("POST", "http://127.0.0.1:7777/output/flypighouse/order", bytes.NewBuffer([]byte(body)))
	if err != nil {
		fmt.Println(err)
	}
	// you can then set the Header here
	// I think the content-type should be "application/xml" like json...
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")
	req.Header.Add("X-Virtual-Env", "feature-flypighousev301-20210714")
	// now POST it
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	defer resp.Body.Close()

	bodyData, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("bodyData:%s\n", string(bodyData))
	var Data CheckOrderSuccessResponse
	err = xml.Unmarshal(bodyData, &Data)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	// Do something with resp.Body
	fmt.Printf("data:%+v", Data)
}

type CheckOrderSuccessResponse struct {
	//XMLName                xml.Name `xml:"Result"`
	Message                string `xml:"Message"`                //错误原因
	CreateOrderValidateKey string `xml:"CreateOrderValidateKey"` //该返回值用于接入方需要在下单时传入的试单信息
	ResultCode             string `xml:"ResultCode"`             //处理结果，0成功
	InventoryPrice         string `xml:"InventoryPrice"`         //价格日历
	CurrencyCode           string `xml:"CurrencyCode"`           //如果是外币，此项为必填，如不返回，默认为CNY
}
