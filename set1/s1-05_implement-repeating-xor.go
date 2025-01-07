package set1

import (
	"encoding/hex"
)

func RepeatingXor(toEncrypt, key string) string {
	textBytes := []byte(toEncrypt) // convert to bytes for XOR operation
	keyBytes := []byte(key)

	for i := 0; i < len(textBytes); i++ {
		// XOR in place
		textBytes[i] = textBytes[i] ^ keyBytes[i%len(keyBytes)] // modulo ensures key bytes will wrap
	}

	return hex.EncodeToString(textBytes)
}
