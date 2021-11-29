package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type WeChat struct {
	AppId       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	AccessToken string `json:"access_token"`
}

type SnsOauth2 struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	Openid       string `json:"openid"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type AccessTokenErrorResponse struct {
	ErrMsg  string `json:"err_msg"`
	ErrCode string `json:"err_code"`
}

//授权
func (weChat *WeChat) GetAuthUrl(redirectUrl string) string {

	oauth2Url := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect",
		weChat.AppId, redirectUrl)
	return oauth2Url
}

//通过code换取网页授权access_token
func (weChat *WeChat) GetWxOpenIdFromOauth2(code string) (*SnsOauth2, error) {

	requestLine := strings.Join([]string{
		"https://api.weixin.qq.com/sns/oauth2/access_token",
		"?appid=", weChat.AppId,
		"&secret=", weChat.AppId,
		"&code=", code,
		"&grant_type=authorization_code"}, "")

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("发送get请求获取 openid 错误", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("发送get请求获取 openid 读取返回body错误", err)
		return nil, err
	}
	if bytes.Contains(body, []byte("errcode")) {
		ater := AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		if err != nil {
			fmt.Printf("发送get请求获取 openid 的错误信息 %+v\n", ater)
			return nil, err
		}
		return nil, fmt.Errorf("%s", ater.ErrMsg)
	} else {
		atr := SnsOauth2{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			fmt.Println("发送get请求获取 openid 返回数据json解析错误", err)
			return nil, err
		}
		return &atr, nil
	}
}
