package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func countRepeats(blocks [][]byte) (int, []byte) {
	seen := make(map[string]bool)
	var repeatedBytes []byte
	repeats := 0
	for _, b := range blocks {
		key := string(b)
		if seen[key] {
			repeats++
			repeatedBytes = b
		} else {
			seen[key] = true
		}

	}
	return repeats + 1, repeatedBytes
}

func main() {
	file, err := os.Open("src/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	maxRepeats := 0
	suspectLine := 0
	var repeatedBlock []byte
	for scanner.Scan() {
		line := scanner.Text()
		data, err := hex.DecodeString(line)
		if err != nil {
			log.Fatalf("Failed to decode line %d: %v", lineNum, err)
		}

		var blocks [][]byte
		for i := 0; i+16 <= len(data); i += 16 {
			blocks = append(blocks, data[i:i+16])
		}

		repeats, repeatedBytes := countRepeats(blocks)
		if repeats > maxRepeats {
			maxRepeats = repeats
			suspectLine = lineNum
			repeatedBlock = repeatedBytes

		}

		lineNum++
	}

	if suspectLine > 0 {
		fmt.Printf("Most likely ECB encrypted line: %d (with %d repeated blocks) %v\n", suspectLine, maxRepeats, hex.EncodeToString(repeatedBlock))
	} else {
		fmt.Println("No ECB pattern detected.")
	}
}
