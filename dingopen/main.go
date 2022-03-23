// This file is auto-generated, don't edit it. Thanks.
package main

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dingtalkoauth2_1_0 "github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"os"
)

/**
 * 使用 Token 初始化账号Client
 * @return Client
 * @throws Exception
 */
func CreateClient() (_result *dingtalkoauth2_1_0.Client, _err error) {
	config := &openapi.Config{}
	config.Protocol = tea.String("https")
	config.RegionId = tea.String("central")
	_result = &dingtalkoauth2_1_0.Client{}
	_result, _err = dingtalkoauth2_1_0.NewClient(config)
	return _result, _err
}

func _main(args []*string) (_err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}

	getCorpAccessTokenRequest := &dingtalkoauth2_1_0.GetCorpAccessTokenRequest{
		//SuiteKey: tea.String("suitep1f5lzyglm7fryxxxx"),
		//SuiteSecret: tea.String("_FP5PpZF3irDKjxxx"),
		//AuthCorpId: tea.String("ding123456"),
		//SuiteTicket: tea.String("1f5lzyglm7fryxxxx"),
		SuiteKey:    tea.String("dingnp0bq0506jxcyleh"),
		SuiteSecret: tea.String("uSDqPXrCgr6cubV-9gpNu5yOfvBxTRc1K_tLeFbiFKImUGJ_w-0lu5Poaxy1vFAo"),
		AuthCorpId:  tea.String("ding0cc7ee3de72c8bac35c2f4657eb6378f"),
		SuiteTicket: tea.String("1647233833"),
	}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, _err = client.GetCorpAccessToken(getCorpAccessTokenRequest)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var err = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			err = _t
		} else {
			err.Message = tea.String(tryErr.Error())
		}
		if !tea.BoolValue(util.Empty(err.Code)) && !tea.BoolValue(util.Empty(err.Message)) {
			// err 中含有 code 和 message 属性，可帮助开发定位问题
			fmt.Printf("sdk errr:%+v", err)
		}

	}
	return _err
}

func main() {
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		fmt.Printf("err:%+v", err)
	}
}
