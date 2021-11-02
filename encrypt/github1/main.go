package main

import (
	"crypto/aes"
	"fmt"
)

func DecryptAes128Ecb(data, key []byte) []byte {
	cipher, _ := aes.NewCipher([]byte(key))
	decrypted := make([]byte, len(data))
	size := 24

	for bs, be := 0, size; bs < len(data); bs, be = bs+size, be+size {
		cipher.Decrypt(decrypted[bs:be], data[bs:be])
	}

	return decrypted
}

func main() {
	aa := DecryptAes128Ecb([]byte("tpBncWHedsSVaLQ8ZN+inQ=="), []byte("a677c601941a4cfe9f1ab63432259f5f"))
	fmt.Println(string(aa))
}
