package set_one

import (
	"encoding/hex"
	"strings"
)

func ScoreString(s string) int {
	score := 0
	goodChars := "aeioustAEIOUST"
	badChars := "!@#$%^&*()-_=+/?><[{]}|]"
	for _, char := range s {
		if strings.ContainsAny(string(char), goodChars) {
			score++
		}
		if strings.ContainsAny(string(char), badChars) {
			score--
		}
	}
	return score
}

func SingleByteXor(toDecrypt string) string {
	bytes, err := hex.DecodeString(toDecrypt)
	if err != nil {
		panic(err)
	}

	highestScore := 0
	// highestScoreIdx := 0
	var englishString string

	for key := 0; key < 256; key++ {
		// XOR each character between 0 and 255 with the input hex string
		output := make([]byte, len(bytes))
		for j := 0; j < len(bytes); j++ {
			xor := bytes[j] ^ byte(key)
			output[j] = xor
		}
		// Determine which string is most likely to be English plaintext
		strScore := ScoreString(string(output))
		if strScore > highestScore {
			highestScore = strScore
			// highestScoreIdx = key
			englishString = string(output)
		}
	}
	// fmt.Printf("And the winning string is:\n%d: %s\nScore of %d!!\n\n", highestScoreIdx, englishString, highestScore)

	return englishString
}
