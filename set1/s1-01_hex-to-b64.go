package set1

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexString string) string {
	decodedHex, err := hex.DecodeString(hexString) // decode hex to bytes

	if err != nil {
		panic(err)
	}

	var encodedBase64 string = base64.StdEncoding.EncodeToString(decodedHex) // encode bytes to base64
	return encodedBase64
}
