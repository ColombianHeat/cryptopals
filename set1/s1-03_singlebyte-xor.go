package set1

import (
	"encoding/hex"
	"strings"
)

func ScoreString(s string) int {
	score := 0
	// goodChars := "aeioustAEIOUST "
	goodChars := "aeiourstlmn AEIOURSTLMN"
	badChars := "!@#$%^&*()-_=+/?><[{]}|]"
	for _, char := range s {
		if strings.ContainsAny(string(char), goodChars) {
			score += 2
		}
		if strings.ContainsAny(string(char), badChars) {
			score--
		}
	}
	return score
}

// SingleByteXor takes a string of hex and decrypts it, assuming it was encrypted by a single byte XOR.
// It returns the decrypted string.
func SingleByteXor(toDecrypt string) (string, int) {
	bytes, err := hex.DecodeString(toDecrypt) // Ensure that the input string is in hexadecimal form!!
	if err != nil {
		panic(err)
	}

	highestScore := 0
	bestKey := 0
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
			bestKey = key
			// highestScoreIdx = key
			englishString = string(output)
		}
	}
	// fmt.Printf("And the winning string is:\n%d: %s\nScore of %d!!\n\n", highestScoreIdx, englishString, highestScore)

	return englishString, bestKey
}
