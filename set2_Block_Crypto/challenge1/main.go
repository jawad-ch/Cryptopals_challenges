package main

import "fmt"

// PKCS#7 padding
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - (len(data) % blockSize)
	padtext := make([]byte, padding)
	for i := 0; i < padding; i++ {
		padtext[i] = byte(padding)
	}
	return append(data, padtext...)
}

func main() {
	src := "YELLOW SUBMARINE"
	padded := pkcs7Pad([]byte(src), 16) // will padd 16 bytes of 16
	fmt.Println(padded)
}
