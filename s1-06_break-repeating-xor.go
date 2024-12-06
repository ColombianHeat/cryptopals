package set_one

import (
	"strings"
)

func CalcHammingDist(str1, str2 string) int {
	bytes1 := []byte(str1)
	bytes2 := []byte(str2)
	if len(bytes1) != len(bytes2) {
		return -1 // str1 and str2 should be of equal length
	}
	hammingBytes := make([]byte, len(bytes1))

	n_differingBits := 0
	for i := 0; i < len(bytes1); i++ {
		hammingBytes[i] = bytes1[i] ^ bytes2[i] // binary number has a 1 in positions where the bits are different
		for j := 0; j < 8; j++ {
			// Now we simply count how many 1s are in all of the hammingBytes array
			// We perform a logical AND with every power of 2. If the result is not 0, then there is a 1
			if hammingBytes[i]&(1<<j) != 0 { // << is bitwise left shift. 1 << j equals 2^j
				n_differingBits++
			}
		}
	}
	return n_differingBits
}

func BreakRepeatingXor(fileDir string) string {
	lines, err := readLines(fileDir)
	if err != nil {
		panic(err)
	}
	var toDecrypt strings.Builder
	for _, line := range lines {
		toDecrypt.WriteString(line)
	}
	for keySize := 0; keySize < 50; keySize++ {
		// For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes, 
		// and find the edit distance between them. Normalize this result by dividing by KEYSIZE. 
	}
	return ""
}
