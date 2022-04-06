package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html"
	"math/rand"
	"strconv"
	"study/tools/util"
	"study/variable"
	"time"
)

func main() {
	//123942367395843

	key1 := "WyIxNjA0MDI0NTI4bm9nMyIseyJzc0lkIjoxNTEzNzk1NTE3MTUzMzIsInNzVHlwZSI6Im1vYmlsZV9jb2RlIiwiZGF0YSI6IjQxMjUyZTBiZTMwNjAxNWFjNDBlYzUzODZiNmJkM2VhIiwiZXhwaXJlIjoxNjYxNTg5OTE2fSwiNjIyOTZjMTY5MjJkOWVjODcwMWViZjllM2VmMDYyNGYiXQ=="
	// userId:123942367395843,session的ssId:151379551715332,expirTime:1661589916
	key2 := "WyIxNjA0MDI0NTI4b1JVZiIseyJzc0lkIjoxNTEzNzk1NTE3MTUzMjksInNzVHlwZSI6Im1vYmlsZV9jb2RlIiwiZGF0YSI6IjQxMjUyZTBiZTMwNjAxNWFjNDBlYzUzODZiNmJkM2VhIiwiZXhwaXJlIjoxNjYxNTg5OTE2fSwiNTNiZWI2NTZlYmJhMGFiYmIwOTViNmVjNWMxMTU5MmQiXQ=="
	//userId:123942367395843,session的ssId:151379551715329,expirTime:1661589916
	key3 := "WyI0NjAzMDExMzI3TWZtWiIseyJzc0lkIjoxNDQxMTc0MDQyNzA1OTUsInNzVHlwZSI6Im1vYmlsZV9jb2RlIiwiZGF0YSI6ImFmMjIyMTQyMWMxOWU1NWE3MzJiYmI3ZDY2NmUxMDBhIiwiZXhwaXJlIjoxNjU4ODE5NjI2fSwiMTNmNzFkMGI1YzMxYmZlNzYyNDNmNWFmN2IyNTlkZjIiXQ=="
	key4 := "WyIyODA3MDQyNDAxMG1EMyIseyJzc0lkIjoxNTg2NTIzNTAzMzMwOTAsInNzVHlwZSI6Im1vYmlsZV9jb2RlIiwiZGF0YSI6IjhkZmIwZmFlZTA3ZDU2ZTY2MjBmNjYxZjcyYzgwNDAwIiwiZXhwaXJlIjoxNjY0MzY0MjY4fSwiNmFlYzc4NDMzMDdiMmY1NDU4NGVmNDg0MjdjYmUxNzEiXQ=="
	VerifySessionKey(key1)
	VerifySessionKey(key2)
	VerifySessionKey(key3)
	VerifySessionKey(key4)
}
func VerifySessionKey(sessionKey string) {

	data := StateDecode(sessionKey, variable.USER_SESSION_SALT) //to-do
	//log.GetLogger(context.Background()).Infof("data:%+v\n", data)
	ssId, err := util.IfInt64(data["ssId"])
	if err != nil {
		fmt.Errorf("err:%+v\n", err)
	}

	expirTime, err := util.IfInt64(data["expire"])
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}

	//方便排查问题
	fmt.Printf("userId:%v,session的ssId:%+v,expirTime:%v\n", "11", ssId, expirTime)

	if data == nil || expirTime < time.Now().Unix() {
		fmt.Errorf("session expired,userId:%d,expire time:%d\n", "122", expirTime)
	}

}

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

func StateGenerate(data interface{}, salt string) string {
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
		fmt.Errorf("err:%s", err.Error())
		return ""
	}
	md5Str := Str2Md5(string(value) + salt)

	Arr = append(Arr, md5Str)
	//fmt.Printf("Arr ... ..:%v", Arr)

	stateValue, err := json.Marshal(Arr)
	if err != nil {
		fmt.Errorf("err:%s\n", err.Error())
		return ""
	}
	encodeString := base64.StdEncoding.EncodeToString(stateValue)
	return encodeString
}

func GetSessId4Html5(uid string, business string) string {
	prefix := "password#"
	if uid == "10828400" {
		return Str2Md5(prefix + uid + "salt4html5force")
	} else if business == "api" {
		//一天过期
		return Str2Md5(prefix + uid + "salt4html5" + time.Now().Format("2006-01-02"))
	} else {
		return Str2Md5(prefix + uid + "salt4html5")
	}
}

func StateDecode(state, salt string) map[string]interface{} {

	stateRes := html.UnescapeString(state)
	if salt == "" {
		salt = GenerateSalt()
	}

	dataArr := GetStateJson(stateRes)
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

func GetStateJson(state string) []interface{} {

	Arr := make([]interface{}, 0)
	stateDecode, err := Base64Decode(state) // php2go.Base64Decode(state)
	if err != nil {
		fmt.Errorf("err:%s", err.Error())
		return Arr
	}

	err = json.Unmarshal([]byte(stateDecode), &Arr)
	if err != nil {
		fmt.Errorf("err:%s", err.Error())
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
	a := t.Format("20060102") + strconv.Itoa(weekInt)
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
