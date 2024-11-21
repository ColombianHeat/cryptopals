package set_one

import "encoding/hex"

func SingleByteXor(toDecrypt string) string {
	bytes, err := hex.DecodeString(toDecrypt)
	if err != nil {
		panic(err)
	}

	keys := "1234567890abcdef"
	for i := 0; i < len(keys); i++ {
		// XOR bytes with each key, then compare the 16 outputs for most likely English phrase.
		// Try counting Es and Ts.
	}
	return ""
}