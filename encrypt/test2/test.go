package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Str2Md5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Str2Md5Raw(str string) string {
	data := []byte(str)
	has := md5.Sum(data)

	fmt.Printf("%v\n", has)
	md5str := fmt.Sprintf("%x", has)
	fmt.Printf("%v\n", md5str)
	return md5str
}

func main() {
	Str2Md5Raw("abcd")
	fmt.Printf(Str2Md5("abcd"))
	fmt.Println()
	fmt.Println("------------")

	//fmt.Println(string(226))

	password := "Lu123456"
	hash, _ := PasswordHash(password)

	fmt.Println("password:", password)
	fmt.Println("hash:", hash)

	match := PasswordVerify(password, hash)
	fmt.Println("Verify:", match)
}
