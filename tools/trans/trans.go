package trans

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

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

func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal with error: %+v\n", err)
		return nil, err
	}
	return m, nil
}

// Convert map json string
func MapToJson(m map[string]interface{}) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Marshal with error: %+v\n", err)
		return "", nil
	}

	return string(jsonByte), nil
}

// vMap := FieldToMap(v, "json")
func FieldToMap(in interface{}, tagFlag string) map[string]interface{} {
	out := make(map[string]interface{})

	inT := reflect.TypeOf(in)
	inV := reflect.ValueOf(in)
	switch {
	case IsStruct(inT):
	case IsStructPtr(inT):
		inT = inT.Elem()
		inV = inV.Elem()
	default:
		return nil
	}

	for i := 0; i < inT.NumField(); i++ {
		if inV.Field(i).Kind() == reflect.Struct {
			out = mergeMap(out, FieldToMap(inV.Field(i).Interface(), tagFlag))
			continue
		}
		var field string
		tag := reflect.StructTag(inT.Field(i).Tag)
		if tag.Get(tagFlag) != "" {
			field = tag.Get(tagFlag)
		}
		if field == "" && tag.Get("json") != "" {
			field = tag.Get("json")
		}
		if field == "" && tag.Get("form") != "" {
			field = tag.Get("form")
		}

		if inV.Field(i).IsZero() {
			continue
		}
		// compatible support default value situation
		field = extractFieldSpec(field)
		out[field] = inV.Field(i).Interface()
	}

	return out
}

func extractFieldSpec(field string) string {
	var spec string
	index := strings.Index(field, ",")
	spec = field
	if index != -1 {
		spec = field[:index]
	}

	return spec
}

func IsStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}

func IsStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func mergeMap(dest, src map[string]interface{}) map[string]interface{} {
	for k, v := range src {
		dest[k] = v
	}

	return dest
}

/*
 * Golang 自实现的三元运算符, https://studygolang.com/articles/3248
 * 例: name := IF(age > 18, "man", "child").(string)
 */
func IF(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func ContainString(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func DiffSlice(sliA []string, sliB []string) []string {
	result := make([]string, 0)
	for _, v := range sliA {
		exist := ContainString(sliB, v)
		if !exist {
			result = append(result, v)
		}
	}
	return result
}

func JsonStructToMap(data interface{}) (map[string]interface{}, error) {
	// 结构体转json
	strRet, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	// json转map
	var mRet map[string]interface{}
	err1 := json.Unmarshal(strRet, &mRet)
	if err1 != nil {
		return nil, err1
	}
	return mRet, nil
}
