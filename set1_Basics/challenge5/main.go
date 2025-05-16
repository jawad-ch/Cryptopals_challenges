package main

import (
	"encoding/hex"
	"fmt"
)

func xorRepeatingKey(plaintext, key []byte) []byte {
	cipher := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		cipher[i] = plaintext[i] ^ key[i%len(key)]
	}
	return cipher
}

var input = `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

func main() {
	plaintext := []byte(input)
	key := []byte("ICE")

	cipher := xorRepeatingKey(plaintext, key)
	fmt.Println("Ciphertext:", hex.EncodeToString(cipher))
}
