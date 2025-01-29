package set_two

import (
	"cryptopals/shared"
	"fmt"
	"math/rand"
)

// RandAESKey returns a randomly generated AES key of length keySize
// The key is suitable for AES encryption and decryption in ECB and CBC modes.
func RandAESKey(keySize int) []byte {
	key := make([]byte, keySize)
	for i := 0; i < keySize; i++ {
		key[i] = byte(rand.Intn(256))
	}
	return key
}

func EncryptAESInCBC(plainText []byte, key string) []byte {
	blockSize := 16 // always for AES
	plainText = ImplementPKCSPadding(plainText, blockSize)
	iv := make([]byte, blockSize) // initialization vector (all zeroes)
	output := make([]byte, len(plainText))

	for bs, be := 0, blockSize; bs < len(plainText); bs, be = bs+blockSize, be+blockSize {
		plainTextBytesBlock := plainText[bs:be]
		// XOR with IV
		plainTextBytesBlock = shared.XorByteVectors(plainTextBytesBlock, iv)
		// Encrypt with ECB
		cipherTextBytesBlock := EncryptAesInECB(plainTextBytesBlock, key)
		// Update output
		copy(output[bs:be], cipherTextBytesBlock)
		// Update IV for next iteration
		iv = cipherTextBytesBlock
	}
	
	return output
}

func ECBorCBCOracle(plainText string) []byte {
	blockSize := 16 // always for AES
	key := RandAESKey(blockSize)
	mode := rand.Intn(2) // 0 = ECB, 1 = CBC

	nAppended := rand.Intn(5) + 5 // 5 to 10 bytes appended
	prefix := make([]byte, nAppended)
	for i := 0; i < nAppended; i++ {
		prefix[i] = byte(rand.Intn(256))
	}
	suffix := make([]byte, nAppended)
	for i := 0; i < nAppended; i++ {
		suffix[i] = byte(rand.Intn(256))
	}
	plainText = string(prefix) + plainText + string(suffix)

	if mode == 0 {
		fmt.Println("Challenge 11: ECB mode")
		plainText = string(ImplementPKCSPadding([]byte(plainText), blockSize))
		return EncryptAesInECB([]byte(plainText), string(key))
	} else if mode == 1 {
		fmt.Println("Challenge 11: CBC mode")
		return EncryptAESInCBC([]byte(plainText), string(key))
	}


	return []byte{}
}