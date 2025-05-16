package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	decodedHex, _ := hex.DecodeString(hexStr)
	b64 := base64.StdEncoding.EncodeToString(decodedHex)

	fmt.Println(string(b64))
}
