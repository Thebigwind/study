package main

import (
	"encoding/xml"
	"fmt"
)

var xmldata = `<?xml version="1.0" encoding="UTF-8"?>
<data>
    <message>
        <status>0</status>
        <value>处理成功</value>
    </message>
    <policeCheckInfos>
        <policeCheckInfo name="朱 XX" id="3412261982XXXXX">
            <message>
                <status>0</status>
                <value>查询成功</value>
            </message>
            <name desc="姓名">朱 XX</name>
            <identitycard desc="身份证号">341226198XXXXXX</identitycard>
            <compStatus desc="比对状态">3</compStatus>
            <compResult desc="比对结果">一致</compResult>
            <checkPhoto desc="照片">base64 码</checkPhoto>
            <no desc="唯一标识" />
        </policeCheckInfo>
    </policeCheckInfos>
</data>`

type GuozhengtongResp struct {
	XMLName xml.Name `xml:"data"`
	Message struct {
		Status string `xml:"status"`
		Value  string `xml:"value"`
	} `xml:"message"`
	PoliceCheckInfos struct {
		PoliceCheckInfo struct {
			AttrName string `xml:"name,attr"` //name
			ID       string `xml:"id,attr"`   //身份证号
			Message  struct {
				Status       string `xml:"status"`       //
				Value        string `xml:"message"`      //
				Name         string `xml:"name"`         //姓名
				Identitycard string `xml:"identitycard"` //身份证号
				CompStatus   string `xml:"compStatus"`   //比对状态
				CompResult   string `xml:"compResult"`   //比对结果
				CheckPhoto   string `xml:"checkPhoto"`   //照片
				No           string `xml:"no"`           //唯一标识
			} `xml:"policeCheckInfo"`
		} `xml:"policeCheckInfos"`
	}
}

type Data struct {
	XMLName xml.Name `xml:"data"`
	Text    string   `xml:",chardata"`
	Message struct {
		Text   string `xml:",chardata"`
		Status struct {
			Text string `xml:",chardata"`
		} `xml:"status"`
		Value struct {
			Text string `xml:",chardata"`
		} `xml:"value"`
	} `xml:"message"`
	PoliceCheckInfos struct {
		Text            string `xml:",chardata"`
		PoliceCheckInfo struct {
			Text     string `xml:",chardata"`
			AttrName string `xml:"name,attr"`
			ID       string `xml:"id,attr"`
			Message  struct {
				Text   string `xml:",chardata"`
				Status struct {
					Text string `xml:",chardata"`
				} `xml:"status"`
				Value struct {
					Text string `xml:",chardata"`
				} `xml:"value"`
			} `xml:"message"`
			Name struct {
				Text string `xml:",chardata"`
				Desc string `xml:"desc,attr"`
			} `xml:"name"`
			Identitycard struct {
				Text string `xml:",chardata"`
				Desc string `xml:"desc,attr"`
			} `xml:"identitycard"`
			CompStatus struct {
				Text string `xml:",chardata"`
				Desc string `xml:"desc,attr"`
			} `xml:"compStatus"`
			CompResult struct {
				Text string `xml:",chardata"`
				Desc string `xml:"desc,attr"`
			} `xml:"compResult"`
			CheckPhoto struct {
				Text string `xml:",chardata"`
				Desc string `xml:"desc,attr"`
			} `xml:"checkPhoto"`
			No struct {
				Text string `xml:",chardata"`
				Desc string `xml:"desc,attr"`
			} `xml:"no"`
		} `xml:"policeCheckInfo"`
	} `xml:"policeCheckInfos"`
}

func main() {
	var data Data
	if err := xml.Unmarshal([]byte(xmldata), &data); err != nil {
		fmt.Printf("err:%+v", err)
	}
	fmt.Printf("data:%+v", data)
}
