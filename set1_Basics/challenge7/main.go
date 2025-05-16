package main

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	ciphertext, err := os.ReadFile("src/file.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	decoded, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return
	}

	key := []byte("YELLOW SUBMARINE")

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return
	}

	plaintext := make([]byte, len(decoded))

	// Decrypt block by block (16 bytes each)
	for start := 0; start < len(decoded); start += 16 {
		block.Decrypt(plaintext[start:start+16], decoded[start:start+16])
	}

	fmt.Println(string(plaintext))
}
