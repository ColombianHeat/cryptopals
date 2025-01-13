package set_two

import (
	"crypto/aes"
	"cryptopals/shared"
	"encoding/base64"
)

func EncryptAesInECB(plainTextBytes []byte, key string) []byte {
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
	return encryptedBytes
}

// The ECB decrypt fnc in set 1 takes a file path as input. This takes a ciphertext string directly
func DecryptAesInECB(ciphertextBytes []byte, key string) []byte {
	blockSize := len([]byte(key))
	cipher, err := aes.NewCipher([]byte(key)) // AES cipher using key of length blocksize
	if err != nil {
		panic(err)
	}
	decryptedBytes := make([]byte, len(ciphertextBytes))

	for bs, be := 0, blockSize; bs < len(ciphertextBytes); bs, be = bs + blockSize , be + blockSize {
		// bs = start of block, be = end of block
		// increment by block size each iteration and update decryptedBytes
		cipher.Decrypt(decryptedBytes[bs:be], ciphertextBytes[bs:be])
	}
	return decryptedBytes
}

// unpadPKCS7 takes a byte slice that has been padded with the PKCS7 algorithm and returns a byte slice with the padding removed.
func unpadPKCS7(plainTextBytes []byte) []byte {
	padLen := 0
	for i := len(plainTextBytes) - 1; i > 0; i-- {
		if plainTextBytes[i] == 4 {
			padLen++
		} else {
			break
		}
	}
	return plainTextBytes[:len(plainTextBytes) - padLen]
}

func DecryptCBCMode(cipherTextPath string, key string) []byte {
	blockSize := len([]byte(key))
	iv := make([]byte, blockSize) // initialization vector (all zeroes)
	ciphertextB64 := shared.ImportTxtFile(cipherTextPath)
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertextB64)
	if err != nil {
		panic(err)
	}

	output := make([]byte, len(ciphertextBytes))
	// XOR IV with first block. Then encrypt with ECB. IV gets updated for next iteration.
	for bs, be := 0, blockSize; bs < len(ciphertextBytes); bs, be = bs+blockSize, be+blockSize {
		cipherTextBytesBlock := ciphertextBytes[bs:be]
		ecbDecrypted := DecryptAesInECB(cipherTextBytesBlock, key) // decrypt with ECB
		plainTextBlock := shared.XorByteVectors(ecbDecrypted, iv) // xor with initialization vector
		copy(output[bs:be], plainTextBlock) // update final fnc output
		iv = cipherTextBytesBlock // update IV for next iteration
	}
	output = unpadPKCS7(output)
	return output
}
