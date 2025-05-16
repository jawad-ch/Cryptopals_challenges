package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func xorSingleByte(plaintext []byte, key byte) []byte {
	cipher := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		cipher[i] = plaintext[i] ^ key
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

	file, err := os.Open("src/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineindex := 1

	var score float64
	var BestByte byte
	var lineNum int
	var targetLine []byte

	for scanner.Scan() {
		hexBytes, err := hex.DecodeString(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		var Ciphertext []byte
		for i := 48; i <= 122; i++ {
			Ciphertext = xorSingleByte(hexBytes, byte(i))

			currScore := ScoreEnglishText(Ciphertext)

			if currScore > score {
				score = currScore
				BestByte = byte(i)
				targetLine = hexBytes
				lineNum = lineindex
			}
		}
		lineindex++
	}

	Ciphertext := xorSingleByte(targetLine, BestByte)
	fmt.Println("Line:", lineNum, "[", string(BestByte), "XOR", hex.EncodeToString(targetLine), "]", "-->", string(Ciphertext))

}
