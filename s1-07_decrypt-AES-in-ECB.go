package set_one

import (
	"crypto/aes"
	"encoding/base64"
	"strings"
)

// import txt file as one long string
func importTxtFile(path string) string {
	lines, err := readLines(path)
	if err != nil {
		panic(err)
	}
	var contents strings.Builder
	for _, line := range lines {
		contents.WriteString(line)
	}
	return contents.String()
}

// decrypts AES-128 in ECB. Key size of 16
func DecryptAesInECB(ciphertextPath string, key string, blockSize int) string {
	b64ToDecrypt := importTxtFile(ciphertextPath)
	ciphertext, err := base64.StdEncoding.DecodeString(b64ToDecrypt)
	if err != nil {
		panic(err)
	}
	cipher, err := aes.NewCipher([]byte(key)) // AES cipher using key of length blocksize
	if err != nil {
		panic(err)
	}
	decryptedBytes := make([]byte, len(ciphertext)) // in AES, lenght of ciphertext bytes and decrypted bytes is always equal

	for bs, be := 0, blockSize; bs < len(ciphertext); bs, be = bs + blockSize , be + blockSize { 
		// bs = start of block, be = end of block
		// increment by block size each iteration and update decryptedBytes
		cipher.Decrypt(decryptedBytes[bs:be], ciphertext[bs:be])
	}
	return string(decryptedBytes)
}
