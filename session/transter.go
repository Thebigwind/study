package tools

import (
	cacheutil "center/pkg/cache/util"
	"center/pkg/global"
	"center/pkg/variable"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gitlab.idc.xiaozhu.com/xz/lib/component/xiaozhu/constant"
	"html"
	"math/rand"
	"strconv"
	"time"

	"center/pkg/cache"
	"center/pkg/log"
)

func StrArr2Int64Arr(arr []string) []int64 {
	int64Arr := make([]int64, 0)
	for _, v := range arr {
		int64Obj, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fmt.Printf("%s transter to int64 error", v)
		}
		int64Arr = append(int64Arr, int64Obj)

	}
	return int64Arr
}

func ContainString(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func Str2Md5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Str2Md5Raw(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	arr := has[:]
	//fmt.Printf("%v\n", arr)

	return string(arr)
}

func GenerateRandomString(length int) string {
	characters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	cLength := len(characters)
	randomString := ""
	for i := 0; i < length; i++ {
		index := rand.Intn(cLength)
		randomString += string(characters[index])
	}
	return randomString
}

type SessionStruct struct {
	SsId   int64  `json:"ssId"`
	SsType string `json:"ssType"`
	Data   string `json:"data"`
	Expire int64  `json:"expire"`
}

func StateGenerate(ctx context.Context, data interface{}, salt string) string {
	//fmt.Printf("data:%+v", data)
	//期望： ssId   ssType data expire
	//now: data expire ssId ssType
	//  data:map[data:3c717b1da387f7f5a7116472037c2931 expire:1657691427 ssId:141159897497602 ssType:mobile_password]

	if salt == "" {
		salt = GenerateSalt()
	}
	/*
	   s - 秒，带前导零（00 到 59）
	   h - 12 小时制，带前导零（01 到 12）
	   m - 月份的数字表示（从 01 到 12）
	   i - 分，带前导零（00 到 59）
	   d - 一个月中的第几天（从 01 到 31）
	*/
	// 2410124930
	// 24秒 10小时 12月 49分 30号
	randString := GetDateshmid() + GenerateRandomString(4)
	Arr := make([]interface{}, 0)
	Arr = append(Arr, randString, data)
	value, err := json.Marshal(Arr)
	if err != nil {
		log.GetLogger(ctx).Debugf("err:%s", err.Error())
		return ""
	}
	md5Str := Str2Md5(string(value) + salt)

	Arr = append(Arr, md5Str)
	//fmt.Printf("Arr ... ..:%v", Arr)

	stateValue, err := json.Marshal(Arr)
	if err != nil {
		log.GetLogger(ctx).Debugf("state decode err:%s", err.Error())
		return ""
	}
	encodeString := base64.StdEncoding.EncodeToString(stateValue)
	return encodeString
}

func GetSessId4Html5(uid string, business string) string {
	prefix := "password#"
	if uid == "10828400" {
		return Str2Md5(prefix + uid + "salt4html5force")
	} else if business == variable.API {
		//一天过期
		return Str2Md5(prefix + uid + "salt4html5" + time.Now().Format(global.DateFmtYMD))
	} else {
		return Str2Md5(prefix + uid + "salt4html5")
	}
}

func StateDecode(ctx context.Context, state, salt string) map[string]interface{} {

	stateRes := html.UnescapeString(state)
	if salt == "" {
		salt = GenerateSalt()
	}

	dataArr := GetStateJson(ctx, stateRes)
	if len(dataArr) < 3 {
		return nil
	}
	if StateVerify(dataArr, salt) {
		return dataArr[1].(map[string]interface{})
	}

	return nil
}

func StateVerify(data []interface{}, salt string) bool {
	if len(data) < 3 {
		return false
	}
	Arr := make([]interface{}, 0)
	/*
		//fmt.Printf("data[1]:%v\n", data[1])

		rebuild原因：如session校验
		php中的data字段顺序：{"ssId":103159524491265,"ssType":"mobile_code","data":"aab68377280b9fb8a3faa57cbf343822","expire":1643195435}
		go中的data字段顺序：{"data":"aab68377280b9fb8a3faa57cbf343822","expire":1643195435,"ssId":103159524491265,"ssType":"mobile_code"}
		以下是为了将go中字段顺序和php保持一致，否则md5加密结果不一致
	*/

	rebuildData := RebuildDataAttrOrder(salt, data[1])
	Arr = append(Arr, data[0], rebuildData)

	value, err := json.Marshal(Arr)
	if err != nil {
		return false
	}

	//fmt.Printf("data[1]:%+v", data[1])
	//fmt.Printf("\n")
	//fmt.Printf("xxxx:%s", string(value)+salt)
	//fmt.Printf("\n")
	//fmt.Printf("md5:%s", Str2Md5(string(value)+salt))
	//fmt.Printf("\n")
	//fmt.Printf("data2:%s\n", data[2].(string))

	return Str2Md5(string(value)+salt) == data[2].(string)
}

func GetOpenApiState(ctx context.Context, md5 string) string {
	redis := cache.New(cache.UserSessionRedisServer)
	resValue, err := redis.Get(md5)
	if err != nil {
		log.GetLogger(ctx).Errorf("Get err:%s", err.Error())
		return ""
	}
	if resValue == nil {
		return ""
	}
	val := cacheutil.IfString(resValue)

	return val
}

func GetStateJson(ctx context.Context, state string) []interface{} {
	if len(state) == 32 {
		state = GetOpenApiState(ctx, state)
	}

	Arr := make([]interface{}, 0)
	stateDecode, err := Base64Decode(state) // php2go.Base64Decode(state)
	if err != nil {
		log.GetLogger(ctx).Errorf("Base64Decode err:%s", err.Error())
		return Arr
	}

	err = json.Unmarshal([]byte(stateDecode), &Arr)
	if err != nil {
		log.GetLogger(ctx).Debugf("Unmarshal err:%s", err.Error())
	}
	return Arr
}

func GenerateSalt() string {
	return variable.OPEN_SALT + GetDateYmdW() //和 php date('YmdW') 保持一致
}

func Base64Decode(str string) (string, error) {
	switch len(str) % 4 {
	case 2:
		str += "=="
	case 3:
		str += "="
	}

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//bool类型转int
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

//int转bool
func Itob(b int) bool {
	return b > 0
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func JsonToMap(jsonStr string) (interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		//log.GetLogger(context.Background()).Errorf("Unmarshal with error: %+v\n", err)
		return nil, err
	}
	return m, nil
}

// Convert map json string
func MapToJson(m interface{}) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		//log.GetLogger(context.Background()).Errorf("err:%s", err.Error())
		return "", nil
	}

	return string(jsonByte), nil
}

func GetDateshmid() string {
	//location, _ := time.LoadLocation("UTC")
	// this should give you time in location
	//t := time.Now().In(location)

	return time.Now().Format("0515010402")
}

/*
   s - 秒，带前导零（00 到 59）
   h - 12 小时制，带前导零（01 到 12）
   m - 月份的数字表示（从 01 到 12）
   i - 分，带前导零（00 到 59）
   d - 一个月中的第几天（从 01 到 31）
*/

//YmdW'
//Y - 年份的四位数表示
//W - 用 ISO-8601 数字格式表示一年中的星期数字（每周从 Monday[星期一]开始）
func GetDateYmdW() string {
	t := time.Now()
	_, weekInt := t.ISOWeek()
	a := t.Format(constant.DateFmtShortYmd) + strconv.Itoa(weekInt)
	return a
}

//php src/Utils/Encryption/ICometTokenChecker.php:48
func MakeTokenV2(userId int64, clientType string) string {
	userIdStr := strconv.FormatInt(userId, 10)

	if clientType == variable.WEB {
		return Str2Md5("ef174b5fbaf02d23580c7c07dbb81c36" + userIdStr)
	} else if clientType == variable.H5 {
		return Str2Md5("44f1495014d93c75478412485f7dc286" + userIdStr)
	} else if clientType == variable.IOS {
		return Str2Md5("4e349e82d03e7b74bb233a76be6c34d2" + userIdStr)
	} else if clientType == variable.ANDROID {
		return Str2Md5("f6e22f5df99fcc7531f4030d99c3a115" + userIdStr)
	}

	return ""
}

// 因为go中解析出的map字段顺序和php中的不一致，需要将go中的字段调整，和保持一致，才能校验成功
func RebuildDataAttrOrder(salt string, data interface{}) interface{} {

	var result interface{}
	switch salt {
	case variable.USER_SESSION_SALT:

		var dd SessionData
		m, ok := data.(map[string]interface{})
		if ok {
			dd.SsId = m["ssId"]
			dd.SsType = m["ssType"]
			dd.Data = m["data"]
			dd.Expire = m["expire"]

			result = dd
		} else {
			result = data
		}

	case variable.LoginToken:

		var dd LoginTokenData
		m, ok := data.(map[string]interface{})
		if ok {
			dd.Userid = m["userid"]
			dd.Expire = m["expire"]
			dd.C = m["c"]

			result = dd
		} else {
			result = data
		}

	default:
		result = data

	}
	return result
}

type SessionData struct {
	SsId   interface{} `json:"ssId"`
	SsType interface{} `json:"ssType"`
	Data   interface{} `json:"data"`
	Expire interface{} `json:"expire"`
}

type LoginTokenData struct {
	Userid interface{} `json:"userid"`
	Expire interface{} `json:"expire"`
	C      interface{} `json:"c"`
}
