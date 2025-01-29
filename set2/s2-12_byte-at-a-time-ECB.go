package set_two

import (
	"cryptopals/shared"
	"encoding/base64"
	"fmt"
)

func S2_12_AES_128_ECB(plainText string, key []byte) []byte {
	blockSize := 16 // always for AES
	toAppendB64 := shared.ImportTxtFile("./data/s2-12.txt")
	toAppendStr := base64.StdEncoding.EncodeToString([]byte(toAppendB64))

	plainText = plainText + string(toAppendStr)

	
	plainText = string(ImplementPKCSPadding([]byte(plainText), blockSize))
	return EncryptAesInECB([]byte(plainText), string(key))
}

func ByteAtATimeECB(yourString string, key []byte) []byte {
	// TODO:  Feed identical bytes of your-string to the function 1 at a time --- start with 1 byte ("A"), then 
	// "AA", then "AAA" and so on. Discover the block size of the cipher. You know it, but do this step anyway.
	// DONE
	outputLen := 0
	nBlockIncreases := 0
	outputLen1 := 0
	outputLen2 := 0
	for i := 0; i < len(yourString); i++ {
		str := yourString[:i]
		encryptedStr := S2_12_AES_128_ECB(str, key) 
		if len(encryptedStr) > outputLen {
			outputLen = len(encryptedStr)
			nBlockIncreases++ // first iteration will always increase this to 1, since outputLen is initialized at 0
		}
		if nBlockIncreases == 2 {
			outputLen1 = outputLen
		}
		if nBlockIncreases == 3 {
			outputLen2 = outputLen
		}
	}
	blockSize := outputLen2 - outputLen1
	fmt.Printf("\n\nLength of output increased at input lengths of %d and %d.\nTherefore the block size of the cipher is %d. \n\n", outputLen1, outputLen2, blockSize)
	
	
	// TODO: Detect that the function is using ECB. You already know, but do this step anyways. 
	
	// TODO: Knowing the block size, craft an input block that is exactly 1 byte short (for instance, if the 
	// block size is 8 bytes, make "AAAAAAA"). Think about what the oracle function is going to put in that 
	// last byte position. 
	
	// TODO: Make a dictionary of every possible last byte by feeding different strings to the oracle; for 
	// instance, "AAAAAAAA", "AAAAAAAB", "AAAAAAAC", remembering the first block of each invocation.
	
	// TODO: Match the output of the one-byte-short input to one of the entries in your dictionary. You've now
	// discovered the first byte of unknown-string.
	
	// TODO: Repeat for the next byte.

	return []byte{}
}
