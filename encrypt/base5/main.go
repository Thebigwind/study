package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	origData := []byte("路雪峰")                         // data to be encrypted
	key := []byte("a677c601941a4cfe9f1ab63432259f5f") // encrypted key
	fmt.Println("Original:", string(origData))

	fmt.Println("------------------ CBC mode --------------------")
	encrypted := AesEncryptCBC(origData, key)
	fmt.Println("Ciphertext(hex):", hex.EncodeToString(encrypted))
	fmt.Println("Ciphertext(base64):", base64.StdEncoding.EncodeToString(encrypted))
	decrypted := AesDecryptCBC(encrypted, key)
	fmt.Println("Decryption result:", string(decrypted))

	fmt.Println("------------------ ECB mode --------------------")
	encrypted = AesEncryptECB(origData, key)
	fmt.Println("Ciphertext(hex):", hex.EncodeToString(encrypted))
	fmt.Println("Ciphertext(base64):", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptECB(encrypted, key)
	fmt.Println("Decryption result:", string(decrypted))

	fmt.Println("------------------ CFB mode --------------------")
	encrypted = AesEncryptCFB(origData, key)
	fmt.Println("Ciphertext(hex):", hex.EncodeToString(encrypted))
	fmt.Println("Ciphertext(base64):", base64.StdEncoding.EncodeToString(encrypted))
	decrypted = AesDecryptCFB(encrypted, key)
	fmt.Println("Decryption result:", string(decrypted))
}

// =================== CBC ======================
func AesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	// // group key
	// NewCipher This function limits the length of input k to 16, 24 or 32
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // Get the length of the key block
	origData = pkcs5Padding(origData, blockSize)                // Completion code
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // encryption mode
	encrypted = make([]byte, len(origData))                     // create an array
	blockMode.CryptBlocks(encrypted, origData)                  // Encrypted
	return encrypted
}
func AesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)                              // grouping key
	blockSize := block.BlockSize()                              // Get the length of the key block
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // encryption mode
	decrypted = make([]byte, len(encrypted))                    // create an array
	blockMode.CryptBlocks(decrypted, encrypted)
	decrypted = pkcs5UnPadding(decrypted)
	return decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// =================== ECB ======================
func AesEncryptECB(origData []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// Block encryption
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// =================== CFB ======================
func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}
func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}
