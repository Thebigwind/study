package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

type AesEncrypt struct {
	keyStr string
}

func NewAesEncrypt(key string) *AesEncrypt {
	return &AesEncrypt{
		keyStr: key,
	}
}

func (this *AesEncrypt) getKey() []byte {
	strKey := this.keyStr
	keyLen := len(strKey)
	if keyLen < 16 {
		panic("Key length is less than 16!")
	}
	arrKey := []byte(strKey)
	if keyLen >= 32 {
		// first 32 bytes
		return arrKey[:32]
	}
	if keyLen >= 24 {
		// first 24 bytes
		return arrKey[:24]
	}
	// first 16 bytes
	return arrKey[:16]
}

// Encrypt string to a secret string
func (this *AesEncrypt) Encrypt(strMesg string) (string, error) {
	key := this.getKey()
	var iv = []byte(key)[:aes.BlockSize]
	encrypted := make([]byte, len(strMesg))
	aesBlockEncrypter, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(encrypted, []byte(strMesg))
	return hex.EncodeToString(encrypted), nil
}

func main() {
	ak := NewAesEncrypt("a677c601941a4cfe9f1ab63432259f5f")
	aaa, _ := ak.Encrypt("路雪峰")
	fmt.Println(aaa)
}
