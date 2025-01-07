package set1

import (
	"encoding/hex"
)

func XorBuffers(buf1, buf2 string) string {
	binary1, err := hex.DecodeString(buf1) // decode to bytes for xor operation
	if err != nil {
		panic(err)
	}
	binary2, err := hex.DecodeString(buf2) // decode to bytes for xor operation
	if err != nil {
		panic(err)
	}

	if len(binary1) != len(binary2) {
		panic("buffers are not the same length")
	}

	buffer_len := len(binary1)

	xor_bin := make([]byte, buffer_len) // empty byte slice

	for i := 0; i < buffer_len; i++ {
		xor_bin[i] = binary1[i] ^ binary2[i] // bytewise xor operation
	}

	xor_hex := hex.EncodeToString(xor_bin) // encode back to hex for output

	return xor_hex
}
