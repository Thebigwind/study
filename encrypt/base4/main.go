package main

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/blowfish"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
)

func encrypt(pt, key []byte) []byte {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Pad(pt) // pad last block of plaintext if block size less than block cipher size
	if err != nil {
		panic(err.Error())
	}
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return ct
}

func decrypt(ct, key []byte) []byte {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBDecrypter(block)
	pt := make([]byte, len(ct))
	mode.CryptBlocks(pt, ct)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Unpad(pt) // unpad plaintext after decryption
	if err != nil {
		panic(err.Error())
	}
	return pt
}

func example() {
	pt := []byte("路雪峰")
	key := []byte("a677c601941a4cfe9f1ab63432259f5f")

	ct := encrypt(pt, key)
	fmt.Printf("Ciphertext: %x\n", ct)
	fmt.Printf("Ciphertext base64: %x\n", base64.StdEncoding.EncodeToString(ct))

	//recoveredPt := decrypt(ct, key)
	//fmt.Printf("Recovered plaintext: %s\n", recoveredPt)
}

func main() {
	fmt.Println("Blowfish encryption with ECB and PKCS5 padding")
	example()
}
