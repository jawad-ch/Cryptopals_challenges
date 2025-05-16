package main

import (
	"encoding/hex"
	"fmt"
)

func xorEncrypt(plaintext, key []byte) []byte {
	cipher := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		cipher[i] = plaintext[i] ^ key[i] // Key repeats if shorter
	}
	return cipher
}

func main() {
	buff1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	buff2, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	cph := xorEncrypt(buff1, buff2)

	fmt.Println("Ciphertext:", hex.EncodeToString(cph))
}
