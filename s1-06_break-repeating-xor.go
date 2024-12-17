package set_one

import (
	"encoding/base64"
	"encoding/hex"
	"sort"
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

func BreakRepeatingXor(fileDir string) (string, string) {
	lines, err := readLines(fileDir)
	if err != nil {
		panic(err)
	}
	var toDecrypt strings.Builder
	for _, line := range lines {
		toDecrypt.WriteString(line)
	}
	toDecryptBytes, err := base64.StdEncoding.DecodeString(toDecrypt.String())
	if err != nil {
		panic(err)
	}
	editDists := make(map[int]float64)
	for keySize := 2; keySize < 40; keySize++ {
		firstBytes := toDecryptBytes[0:keySize]
		secondBytes := toDecryptBytes[keySize : keySize*2]
		thirdBytes := toDecryptBytes[keySize*2 : keySize*3]
		fourthBytes := toDecryptBytes[keySize*3 : keySize*4]
		fifthBytes := toDecryptBytes[keySize*4 : keySize*5]
		sixthBytes := toDecryptBytes[keySize*5 : keySize*6]
		seventhBytes := toDecryptBytes[keySize*6 : keySize*7]
		eighthBytes := toDecryptBytes[keySize*7 : keySize*8]
		editDist1 := float64(CalcHammingDist(string(firstBytes), string(secondBytes))) / float64(keySize)
		editDist2 := float64(CalcHammingDist(string(thirdBytes), string(fourthBytes))) / float64(keySize)
		editDist3 := float64(CalcHammingDist(string(fifthBytes), string(sixthBytes))) / float64(keySize)
		editDist4 := float64(CalcHammingDist(string(seventhBytes), string(eighthBytes))) / float64(keySize)
		editDist := (editDist1 + editDist2 + editDist3 + editDist4) / 4.0
		editDists[keySize] = editDist
	}

	// sort editDists based on the edit distance
	type kv struct {
		key int
		val float64
	}

	var sortedEditDists []kv
	for k, v := range editDists {
		sortedEditDists = append(sortedEditDists, kv{k, v})
	}
	sort.Slice(sortedEditDists, func(i, j int) bool {
		return sortedEditDists[i].val < sortedEditDists[j].val
	})

	// get the 3 key sizes corresponding to the three lowest edit distances
	probableKeySizes := []int{}
	for _, kv := range sortedEditDists[:3] {
		probableKeySizes = append(probableKeySizes, kv.key)
	}

	// Try to break the cipher with each of our three key sizes
	var highestFinalDecrypted string
	var highestKeyString string
	finalDecryptedScore := 0
	for _, keySize := range probableKeySizes {
		fullKey := make([]int, keySize) // Store the derived key bytes for the current key size
		decryptedBlocks := make([][]byte, keySize) // Store the decrypted blocks for later reconstruction

		for i := 0; i < keySize; i++ {
			var block []byte
			// Transpose the ciphertext into blocks based on the current key size
			for j := 0; j < len(toDecryptBytes); j++ {
				if j%keySize == i {
					block = append(block, toDecryptBytes[j])
				}
			}
			// Decrypt the current block using single-byte XOR
			decryptedUnordered, key := SingleByteXor(hex.EncodeToString(block))
			decryptedBlocks[i] = []byte(decryptedUnordered) // Store the decrypted block
			fullKey[i] = key // Store the key byte for this block
		}

		// Generate the final decrypted string
		var finalDecrypted strings.Builder
		// Iterate through the blocks in an interleaved manner to reconstruct the plaintext
		for j := 0; j < len(toDecryptBytes); j++ {
			blockIndex := j % keySize // Determine which block this character belongs to
			charIndex := j / keySize // Determine the position within that block
			if charIndex < len(decryptedBlocks[blockIndex]) {
				finalDecrypted.WriteByte(decryptedBlocks[blockIndex][charIndex]) // Append the character to the final string
			}
		}

		// Print the reconstructed plaintext and the derived key
		keyString := string(convertToBytes(fullKey))
		if ScoreString(finalDecrypted.String()) > finalDecryptedScore {
			highestFinalDecrypted = finalDecrypted.String()
			finalDecryptedScore = ScoreString(finalDecrypted.String())
			highestKeyString = keyString
		}
	}
return highestFinalDecrypted, highestKeyString
}

func convertToBytes(ints []int) []byte {
	bytes := make([]byte, len(ints))
	for i, v := range ints {
		bytes[i] = byte(v) // Safely cast each integer to a byte
	}
	return bytes
}

