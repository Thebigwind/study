package main

import (
	"bytes"
	"crypto/aes"
	//"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var Key = "a677c601941a4cfe9f1ab63432259f5f"

func main() {
	fmt.Println("------------test1---------")

	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("a677c601941a4cfe9f1ab63432259f5f")
	result, err := AesEncrypt([]byte("路雪峰"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	fmt.Println("------------test2---------")

	result = EcbEncrypt([]byte("路雪峰"), key)
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	fmt.Println("------------test3---------")
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	//fmt.Println("blockSize:%d",blockSize)
	origData = PKCS5Padding(origData, blockSize)
	fmt.Println("origin:", string(origData))
	blockMode := NewECBEncrypter(block)
	//blockMode := cipher.NewCBCDecrypter (block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	blockMode.CryptBlocks(crypted, origData)

	return crypted, nil
}

///////////////////////////////////////////////////////

func EcbEncrypt(data, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	data = PKCS5Padding(data, block.BlockSize())
	decrypted := make([]byte, len(data))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		block.Encrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

//////////////////////////////////////////

func AESECB(ciphertext []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(Key))
	fmt.Println("AESing the data")
	bs := 24
	if len(ciphertext)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	plaintext := make([]byte, len(ciphertext))
	for len(plaintext) > 0 {
		cipher.Decrypt(plaintext, ciphertext)
		plaintext = plaintext[bs:]
		ciphertext = ciphertext[bs:]
	}
	return plaintext
}

//"tpBncWHedsSVaLQ8ZN+inQ==" cbc
//o+miQe3vw6YWnqTFJlWpvg==  ecb
func testAes() {
	// AES-128。key长度：16, 24, 32 bytes 对应 AES-128, AES-192, AES-256
	key := []byte("a677c601941a4cfe9f1ab63432259f5f")
	result, err := AesEncrypt([]byte("路雪峰"), key)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	fmt.Println("-------------")

}
