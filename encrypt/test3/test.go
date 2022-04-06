package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

const itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

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

	hash := Str2Md5(salt + md5Pass)[8:24]
	for i := 0; i < count; i++ {
		hash = Str2Md5(hash + md5Pass)[8:24]
	}
	output = dbPass[:12]
	// :$P$Bi9sq6/ml
	fmt.Printf("111:%s\n", output)
	fmt.Printf("hash:%s\n", hash) //67f08934c146dac19579296829d4c1c8
	//当前 qQnBrMaNa/1AkU1CsYHCtAnAnE1BoAqMX3HAlE1BoMXBqE4NY3KMVBqMX3HAl.， 期望 2tapV278beH1Gt6kMbOKo0
	fmt.Printf("222:%s\n", encode64(hash, 16))
	output = output + encode64(hash, 16)
	return output
}

func encode64(input string, count int) string {
	output := ""
	i := 0
	for {
		r := input[i]
		i = i + 1
		value := int(r)
		output = output + string(rune(itoa64[value&0x3f]))

		if i < count {
			value |= int(input[i]) << 8
		}

		output = output + string(rune(itoa64[(value>>6)&0x3f]))
		if i >= count {
			i = i + 1
			break
		}
		if i < count {
			value |= int(input[i]) << 16
		}
		output = output + string(rune(itoa64[(value>>12)&0x3f]))
		if i >= count {
			i = i + 1
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

func Str2Md5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
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

func Str2Md5Raw(str string) string {
	h := md5.New()
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
	//h.Write([]byte("hello"))

	//fmt.Println(base64.StdEncoding.EncodeToString(h.Sum(nil)))

	//fmt.Println("Hello, playground")
}

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
