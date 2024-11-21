package set_one

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexString string) string {
	decodedHex, err := hex.DecodeString(hexString)

	if err != nil {
		panic(err)
	}

	var encodedBase64 string = base64.StdEncoding.EncodeToString(decodedHex)
	return encodedBase64
}
