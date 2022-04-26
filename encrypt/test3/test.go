package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/imroc/biu"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

// https://www.csdn.net/tags/MtTaMg1sMTE0NTM3LWJsb2cO0O0O.html
const itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func main() {
	reqPass := "Lu123456"
	dbPass := "$P$Bi9sq6/ml2tapV278beH1Gt6kMbOKo0"
	//P$Bi9sq6/ml2tapV278beH1Gt6kMbOKo0
	fmt.Printf("Str2Md5(reqPass):%s\n", Str2Md5(reqPass))
	if Str2Md5(reqPass) != dbPass && !CheckPassword(Str2Md5(reqPass), dbPass) {
		fmt.Println(false)
	} else {
		fmt.Println(true)
	}
}
func Str2Md5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Str2Md5Raw2(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	//fmt.Printf("%v\n", has)
	//md5str := fmt.Sprintf("%B", has)
	arr := []byte{}
	arr = has[:]

	md5str := biu.ToBinaryString(arr)
	//md5str = strings.Trim(md5str,"[]")
	//fmt.Printf("md5url:%s\n",md5str)
	return md5str
}

func Str2Md5Raw(str string) string {
	data := []byte(str)
	has := md5.Sum(data)

	arr := has[:]
	//fmt.Printf("%v\n", arr)

	return string(arr)
}

func CryptPassword(md5Pass, dbPass string) string {
	output := "*0"
	if dbPass[:2] == output {
		output = "*1"
	}
	id := dbPass[:3]
	if id != "$P$" && id != "$H$" {
		return output
	}

	count_index := strings.Index(itoa64, dbPass[3:4])
	if count_index < 7 || count_index > 30 {
		return output
	}

	count := 1 << count_index

	salt := dbPass[4:12]
	if len(salt) != 8 {
		return output
	}

	fmt.Printf("salt + md5Pass:%s\n", salt+md5Pass) //i9sq6/ml10d50a1e43a718a9c31bc450a0985131
	hash := Str2Md5Raw(salt + md5Pass)
	fmt.Printf("count:%d\n", count)
	for i := 0; i < count; i++ {
		hash = Str2Md5Raw(hash + md5Pass)
	}
	output = dbPass[:12]
	// :$P$Bi9sq6/ml
	//fmt.Printf("111:%s\n", output)
	fmt.Printf("hash:%s,len:%d\n", hash, len(hash)) //67f08934c146dac19579296829d4c1c8
	//当前 qQnBrMaNa/1AkU1CsYHCtAnAnE1BoAqMX3HAlE1BoMXBqE4NY3KMVBqMX3HAl.， 期望 2tapV278beH1Gt6kMbOKo0
	encodeStr := encode64(hash, 16)
	fmt.Printf("encodeStr:%s\n", encodeStr) //2taPiNhpK5G6V2NYFW08cQudbeXCuoE1B6ZIGtcXC0Ak.XBqMbOedaJKNF9ho0
	output = output + encodeStr
	return output
}

func encode64(input string, count int) string {
	for i := 0; i < 16; i++ {
		fmt.Printf("value:%v\n", input[i])
	}
	output := ""
	i := 0
	for {
		fmt.Printf("i:%v,input[i]:%v\n", i, int(input[i]))
		r := input[i]
		i = i + 1
		value := int(r)

		//fmt.Printf("value:%d\n", value)

		output = output + string(rune(itoa64[value&0x3f]))

		if i < count {
			value |= int(input[i]) << 8
		}

		output = output + string(rune(itoa64[(value>>6)&0x3f]))
		i = i + 1
		if i-1 >= count {
			//i = i + 1
			break
		}
		if i < count {
			value |= int(input[i]) << 16
		}
		output = output + string(rune(itoa64[(value>>12)&0x3f]))
		i = i + 1
		if i-1 >= count {
			//i = i + 1
			break
		}

		output = output + string(rune(itoa64[(value>>18)&0x3f]))

		if i >= count {
			break
		}
	}
	return output
}

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckPassword(md5Pass, dbPass string) bool {
	if len(md5Pass) > 4096 {
		return false
	}

	hash := CryptPassword(md5Pass, dbPass)
	fmt.Printf("hash :%v\n", hash)
	if hash[:1] == "*" {
		has, err := PasswordHash(hash)
		fmt.Printf("has :%v\n", has)
		if err == nil {
			hash = has
		}
		return PasswordVerify(hash, dbPass)

	}
	return hash == dbPass
}

func Hash2bin(hash string) (string, int, error) {
	binary_string := ""
	for _, char := range hash {
		char_hex, err := strconv.ParseInt(string(char), 16, 8)
		if err != nil {
			return "", 0, err
		}
		char_bin := ""
		for ; char_hex > 0; char_hex /= 2 {
			b := char_hex % 2
			char_bin = strconv.Itoa(int(b)) + char_bin
		}
		fill := 4 - len(char_bin)
		for fill > 0 {
			char_bin = "0" + char_bin
			fill -= 1
		}
		binary_string += char_bin
	}
	return binary_string, len(binary_string), nil
}
