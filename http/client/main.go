package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type CheckOrderSuccessResponse struct {
	//XMLName                xml.Name `xml:"Result"`
	Message                string `xml:"Message"`                //错误原因
	CreateOrderValidateKey string `xml:"CreateOrderValidateKey"` //该返回值用于接入方需要在下单时传入的试单信息
	ResultCode             string `xml:"ResultCode"`             //处理结果，0成功
	InventoryPrice         string `xml:"InventoryPrice"`         //价格日历
	CurrencyCode           string `xml:"CurrencyCode"`           //如果是外币，此项为必填，如不返回，默认为CNY
}

func main() {
	const myurl = "http://127.0.0.1:7777/output/flypighouse/order"
	const xmlbody = `
<ValidateRQ>
  <AuthenticationToken>
    <Username>556471e1</Username>
    <Password>0915350284be0668</Password>
    <CreateToken>22251178182015010620150107497867981843210904377</CreateToken>
  </AuthenticationToken>
  <TaoBaoHotelId>13577578181</TaoBaoHotelId>
  <HotelId>74551593009156</HotelId>
  <TaoBaoRoomTypeId>5501264818</TaoBaoRoomTypeId>
  <RoomTypeId>75480776835156</RoomTypeId>
  <TaoBaoRatePlanId>4978679818</TaoBaoRatePlanId>
  <RatePlanCode>75513705005058</RatePlanCode>
  <TaoBaoGid>3824371818</TaoBaoGid>
  <CheckIn>2021-09-10</CheckIn>
  <CheckOut>2021-09-11</CheckOut>
  <RoomNum>1</RoomNum>
  <PaymentType>1</PaymentType>
  <Extensions>{"searchid":"22251178182015010620150107497867981843210904377"}  </Extensions>
  <TotalPrice>7949</TotalPrice>
  <DailyInfos>
    <DailyInfo>
      <Day>2019-12-30</Day>
      <Price>7949</Price>
      <BreakFast>2</BreakFast>
     </DailyInfo>
  </DailyInfos>
  <CurrencyCode>CNY</CurrencyCode>
</ValidateRQ>`

	resp, err := http.Post(myurl, "application/xml", strings.NewReader(xmlbody)) //"text/xml"
	if err != nil {
		log.Fatal(err)
	}
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
