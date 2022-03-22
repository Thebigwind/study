package variable

const OPEN_SALT = "xiaozhu&?f#/^z//|`_"

const STATE_VALID = "valid"
const STATE_DELETED = "deleted"

const WEB = "web"
const ANDROID = "android"
const HTML5 = "html5"
const H5 = "h5"
const IOS = "ios"
const ALIPAYAPPLET = "AlipayApplet"
const WECHATAPPLET = "WeChatApplet"
const TOUTIAOAPPLET = "toutiaoapplet"
const BAIDUAPPLET = "BaiduApplet"
const QUICKAPPHUAWEI = "QuickAppHuaWei"

const RAKE = "rake"
const FEISHUAPPLET = "feiShuApplet"
const XIAOHONGSHU = "xiaoHongShu"

var DefinesClient = map[string]string{
	WEB:            "网站",
	HTML5:          "html5",
	ANDROID:        "android",
	IOS:            "ios",
	RAKE:           "rake",
	ALIPAYAPPLET:   "AlipayApplet",
	WECHATAPPLET:   "WeChatApplet",
	TOUTIAOAPPLET:  "toutiaoapplet",
	BAIDUAPPLET:    "BaiduApplet",
	FEISHUAPPLET:   "feiShuApplet",
	XIAOHONGSHU:    "xiaoHongShu",
	QUICKAPPHUAWEI: "QuickAppHuaWei",
}

//UserSettings
const API = "api"
const USER_SESSION_SALT = "DeviceLoginSession.X$f25e6168-2ef7-4569-b655-OagT8PWybs"

func GetNotXZAppClient() []string {
	strArr := make([]string, 7)
	strArr = append(strArr, HTML5, ALIPAYAPPLET, WECHATAPPLET, TOUTIAOAPPLET, WEB, BAIDUAPPLET, QUICKAPPHUAWEI)
	return strArr
}

//login session expiretime

const USERNAME_PASSWORD = "username_password"
const MOBILE_PASSWORD = "mobile_password"
const MOBILE_CODE = "mobile_code"
const EMAIL_PASSWORD = "email_password"
const EMAIL_CODE = "email_code"
const OPENACCOUNT_QQ = "openaccount_qq"
const OPENACCOUNT_WECHAT = "openaccount_wechat"
const OPENACCOUNT_WECHATSL = "openaccount_wechatsl"
const OPENACCOUNT_WEIBO = "openaccount_weibo"
const OPENACCOUNT_FUWUCHUANG = "openaccount_fuwuchuang"
const MOBILE_SIM = "mobile_sim"

var LoginSessionMap = map[string]string{
	USERNAME_PASSWORD:      "用户名+密码登录",
	EMAIL_PASSWORD:         "邮箱+密码登录",
	MOBILE_PASSWORD:        "手机号+密码登录",
	MOBILE_CODE:            "邮箱+验证码登录",
	EMAIL_CODE:             "邮箱+验证码登录",
	OPENACCOUNT_QQ:         "第三方QQ授权登录",
	OPENACCOUNT_WECHAT:     "第三方微博授权登录",
	OPENACCOUNT_WECHATSL:   "第三方微博授权登录",
	OPENACCOUNT_WEIBO:      "第三方微博授权登录",
	OPENACCOUNT_FUWUCHUANG: "第三方支付宝服务窗授权登录",
	MOBILE_SIM:             "手机免密登录",
}

var LoginSessionExpireTimeMap = map[string]int{
	USERNAME_PASSWORD:      15552000,
	EMAIL_PASSWORD:         15552000,
	MOBILE_PASSWORD:        15552000,
	MOBILE_CODE:            15552000,
	EMAIL_CODE:             2592000,
	OPENACCOUNT_QQ:         2592000,
	OPENACCOUNT_WECHAT:     2592000,
	OPENACCOUNT_WECHATSL:   2592000,
	OPENACCOUNT_WEIBO:      2592000,
	OPENACCOUNT_FUWUCHUANG: 2592000,
	MOBILE_SIM:             2592000,
}

func GetExpireTime(key string) int {
	if value, ok := LoginSessionExpireTimeMap[key]; ok {
		return value
	}
	return 604800
}

//第三方
const QQZONE = "qqzone"
const SINAWB = "sinaweibo"
const WEIXIN = "weixin"
const QQ = "qq"
const WEIXIN_APPLET = "wxminiprgm"
const DOUBAN = "douban"
const RENREN = "renren"
const WEIXINDY = "weixindy"
const WEIXINSL = "weixinsl"
const WEIXINCLEANER = "weixincleaner"
const WEIXINOPEN = "weixinopen"
const FUWUCHUANG = "fuwuchuang"
const FRIENDCIRCLE = "friendcircle"
const XIAOBAI = "xiaobai"
const TAOBAO = "taobao"
const OTHER = "other"
const ALIPAY_MINI_PROGRAM = "alipayminiprgm"

//useraddress使用
var UserAddressBizlineMap = map[string]string{
	"bizrent_main_default": "100100100",
	"bizmall_main_default": "300100100",
}

/*
func VerifyParamsTest(sessionType string, user http.UserInfo) string {
	paramsArray := make([]string, 0)
	switch sessionType {
	case USERNAME_PASSWORD:
		paramsArray = append(paramsArray, user.RealName, user.PassportNo)
	case EMAIL_PASSWORD:
		paramsArray = append(paramsArray, user.Email, user.PassportNo)
	case MOBILE_PASSWORD:
		paramsArray = append(paramsArray, user.Mobile, user.PassportNo)
	case MOBILE_CODE:
		paramsArray = append(paramsArray, user.Mobile, "mobilecode")
	case EMAIL_CODE:
		paramsArray = append(paramsArray, user.Email, "mobilecode")
	case OPENACCOUNT_QQ: //to-do
		paramsArray = append(paramsArray, "", "qq") //userInfo 已经没有了QQ相关字段 QqOpenId
	case OPENACCOUNT_WECHAT:
		if user.WeixinAuth == "yes" {
			paramsArray = append(paramsArray, user.WeixinUnionID, WEIXIN)
		}
	case OPENACCOUNT_WEIBO:
		if user.SinaWeiboAuth == "yes" {
			paramsArray = append(paramsArray, user.SinaWeiboOpenID, SINAWB)
		}
	case FUWUCHUANG:
		if user.FuwuchuangAuth == "yes" {
			paramsArray = append(paramsArray, user.FuwuchuangOpenID, FUWUCHUANG)
		}
	case MOBILE_SIM:
		paramsArray = append(paramsArray, user.FuwuchuangOpenID, FUWUCHUANG)
	}

	result := ""
	if len(paramsArray) > 0 {

	}
	return result
}
*/

const LoginToken = "xiaozhu_2Tu3Uj1Qs0"

var DefaultImg = []int64{1009224, 1009225, 1009226, 1009227, 1009228, 1009229, 1009230, 1009231}

func IsDefaultHeadImg(headImgId int64) bool {
	for i := 0; i < len(DefaultImg); i++ {
		if DefaultImg[i] == headImgId {
			return true
		}
	}
	return false
}

///////
var SexMap = map[string]string{
	"man":   "男",
	"women": "女",
}

var BloodTypeMap = map[string]string{
	"a":  "A型",
	"b":  "B型",
	"ab": "AB型",
	"o":  "O型",
}

var EduMap = map[string]string{
	"doctor":        "博士",
	"master":        "硕士",
	"undergraduate": "本科",
	"college":       "大专",
	"secondary":     "中专",
	"senior":        "高中",
	"junior":        "初中",
}
