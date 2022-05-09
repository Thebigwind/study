package test4

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
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

	hash := Str2Md5Raw(salt + md5Pass)
	for i := 0; i < count; i++ {
		hash = Str2Md5Raw(hash + md5Pass)
	}
	output = dbPass[:12]
	output = output + encode64(hash, 16)
	return output
}

func encode64(input string, count int) string {
	//for i := 0; i < 16; i++ {
	//	fmt.Printf("value:%v\n", input[i])
	//}
	output := ""
	i := 0
	for {
		//fmt.Printf("i:%v,input[i]:%v\n", i, int(input[i]))
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
	if hash[:1] == "*" {
		has, err := PasswordHash(hash)
		if err == nil {
			hash = has
		}
	}
	//return PasswordVerify(hash, dbPass)
	return hash == dbPass
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
