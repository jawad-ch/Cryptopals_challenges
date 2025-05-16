package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

func xorEncrypt(plaintext, key []byte) []byte {
	cipher := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		cipher[i] = plaintext[i] ^ key[i%len(key)] // Key repeats if shorter
	}
	return cipher
}

func xorSingleByte(plaintext []byte, key byte) []byte {
	cipher := make([]byte, len(plaintext))

	for i := 0; i < len(plaintext); i++ {
		cipher[i] = plaintext[i] ^ key
	}
	return cipher
}

func hamming_distance(b1, b2 []byte) int {

	distance := 0
	var xor byte
	for i, b := range b1 {
		xor = b ^ b2[i%len(b2)]
		for xor > 0 {
			distance += int(xor & 1)
			xor >>= 1
		}
	}

	return distance
}

func normalize_distance(data []byte, keysize int) float64 {

	distances := 0
	comparisons := 0
	// Compare multiple blocks to get a better average
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			block1 := data[i*keysize : (i+1)*keysize]
			block2 := data[j*keysize : (j+1)*keysize]
			distances += hamming_distance(block1, block2)
			comparisons++
		}
	}

	return float64(distances) / float64(comparisons*keysize)
}

func FindKeySize(data []byte) int {

	lowestDistance := float64(999999)
	var bestKeySize int
	for size := 2; size <= 40; size++ {

		distance := normalize_distance(data, size)

		if distance < lowestDistance {
			lowestDistance = distance
			bestKeySize = size
		}
	}

	return bestKeySize

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

func transposeBlocks(data []byte, keysize int) [][]byte {
	chunks := make([][]byte, keysize)

	blockID := 0
	for i, b := range data {
		blockID = i % keysize
		chunks[blockID] = append(chunks[blockID], b)
	}

	return chunks
}

func FindBestSingleKey(block []byte) byte {
	var xor []byte
	var bestKey byte

	bestScore := -1.0
	for key := 0; key <= 255; key++ {
		xor = xorSingleByte(block, byte(key))
		chunkScore := ScoreEnglishText(xor)
		if chunkScore > bestScore {
			bestScore = chunkScore
			bestKey = byte(key)
		}
	}

	return bestKey
}

func main() {
	// Read and decode input file
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

	fmt.Println("text in bytes", len(decoded))

	bestKeySize := FindKeySize(decoded)

	blocks := transposeBlocks(decoded, bestKeySize)

	decrypteKey := make([]byte, bestKeySize)
	for i, block := range blocks {
		decrypteKey[i] = FindBestSingleKey(block)
	}

	decryptedFile := xorEncrypt(decoded, decrypteKey)

	fmt.Println("key is : ", string(decrypteKey))
	fmt.Println("\ndecrypted file :\n\n", string(decryptedFile))
	fmt.Println("key is : ", string(decrypteKey))

}
