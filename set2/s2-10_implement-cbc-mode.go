package set_two

import (
	"crypto/aes"
	"cryptopals/shared"
	"encoding/base64"
)

func EncryptAesInECB(plainText string, key string) string {
	plainTextBytes := []byte(plainText)
	blockSize := len([]byte(key))
	cipher, err := aes.NewCipher([]byte(key)) // AES cipher using key of length blocksize. 16 bytes in this case
	if err != nil {
		panic(err)
	}
	encryptedBytes := make([]byte, len(plainTextBytes))
	
	for bs, be := 0, blockSize; bs < len(plainTextBytes); bs, be = bs+blockSize, be+blockSize {
		// bs = start of block, be = end of block
		// increment by block size each iteration and update decryptedBytes
		cipher.Encrypt(encryptedBytes[bs:be], plainTextBytes[bs:be])
	}
	return string(encryptedBytes)
}

// The ECB decrypt fnc in set 1 takes a file path as input. This takes a ciphertext string directly
func DecryptAesInECB(ciphertext string, key string) string {
	ciphertextBytes := []byte(ciphertext)
	blockSize := len([]byte(key))
	cipher, err := aes.NewCipher([]byte(key)) // AES cipher using key of length blocksize
	if err != nil {
		panic(err)
	}
	decryptedBytes := make([]byte, len(ciphertext))

	for bs, be := 0, blockSize; bs < len(ciphertext); bs, be = bs + blockSize , be + blockSize { 
		// bs = start of block, be = end of block
		// increment by block size each iteration and update decryptedBytes
		cipher.Decrypt(decryptedBytes[bs:be], ciphertextBytes[bs:be])
	}
	return string(decryptedBytes)
}

func XorByteVectors(a []byte, b []byte) []byte {
	if len(a) != len(b) {
		panic("buffers are not the same length")
	}
	vectorLen := len(a)
	xorBin := make([]byte, vectorLen) // empty byte slice
	for i := 0; i < vectorLen, i++ {
		xorBin[i] = a[i] ^ b[i]
	}
	return xorBin
}

func ImplementCBCMode(ciphertextPath string, key string) string {
	blockSize := len([]byte(key))
	iv := make([]byte, blockSize)
	for i := 0; i < blockSize; i++ {
		iv[i] = byte(0) // initialization vector is all ASCII 0
	}
	ciphertextB64 := shared.ImportTxtFile(ciphertextPath)
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertextB64)
	if err != nil {
		panic(err)
	}

	// FIXME: Decode using CBC. Still not sure what the steps are. Research.
	for bs, be := 0, blockSize; bs < len(ciphertextBytes); bs, be = bs+blockSize, be+blockSize {
		ecbEncrypted := EncryptAesInECB(string(ciphertextBytes[bs:be]), key) // encrypt to ECB
		iv = XorByteVectors([]byte(ecbEncrypted), iv) // xor with initialization vector
	}
	return ""
}
