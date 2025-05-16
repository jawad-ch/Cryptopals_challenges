package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func xorSingleByte(plaintext []byte, key byte) []byte {
	cipher := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		cipher[i] = plaintext[i] ^ key // Key repeats if shorter
	}
	return cipher
}

func ScoreEnglishText(text []byte) float64 {
	freqMap := map[byte]float64{
		'a': 8.2, 'b': 1.5, 'c': 2.8, 'd': 4.3, 'e': 12.7, 'f': 2.2, 'g': 2.0,
		'h': 6.1, 'i': 7.0, 'j': 0.2, 'k': 0.8, 'l': 4.0, 'm': 2.4, 'n': 6.7,
		'o': 7.5, 'p': 1.9, 'q': 0.1, 'r': 6.0, 's': 6.3, 't': 9.1, 'u': 2.8,
		'v': 1.0, 'w': 2.4, 'x': 0.2, 'y': 2.0, 'z': 0.1, ' ': 13.0,
	}

	score := 0.0
	for i, b := range text {
		// Reward capital letters at the beginning
		if i == 0 && b >= 'A' && b <= 'Z' {
			score += 2.0
		}

		// Penalize unprintables
		if b < 32 || b > 126 {
			score -= 10
			continue
		}

		// Add small score for common punctuation
		switch b {
		case ' ', '.', ',', '\'', '-', ';', ':':
			score += 2.0
		}

		// Normalize for letter frequencies
		if b >= 'A' && b <= 'Z' {
			b += 32 // to lowercase
		}
		if freq, ok := freqMap[b]; ok {
			score += freq
		}
	}

	return score
}

func main() {
	hexBytes, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		log.Fatal(err)
	}

	var Ciphertext []byte
	var score float64
	var BestByte byte
	for i := 0; i <= 255; i++ {
		Ciphertext = xorSingleByte(hexBytes, byte(i))
		currScore := ScoreEnglishText(Ciphertext)

		if currScore > score {
			score = currScore
			BestByte = byte(i)
		}
	}

	Ciphertext = xorSingleByte(hexBytes, BestByte)
	fmt.Println(string(BestByte), "-->", string(Ciphertext))
}
