package set_one

import (
	"encoding/hex"
	"fmt"
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

// func sanitizeOutput(input string) string {
// 	sanitized := make([]rune, 0, len(input))
// 	for _, r := range input {
// 		if r >= 32 && r <= 126 { // Printable ASCII range
// 			sanitized = append(sanitized, r)
// 		}
// 	}
// 	return string(sanitized)
// }

func BreakRepeatingXor(fileDir string) string {
	lines, err := readLines(fileDir)
	if err != nil {
		panic(err)
	}
	var toDecrypt strings.Builder
	for _, line := range lines {
		toDecrypt.WriteString(line)
	}
	editDists := make(map[int]float64)
	for keySize := 2; keySize < 40; keySize++ {
		firstBytes := []byte(toDecrypt.String())[0:keySize]
		secondBytes := []byte(toDecrypt.String())[keySize : keySize*2]
		thirdBytes := []byte(toDecrypt.String())[keySize*2 : keySize*3]
		fourthBytes := []byte(toDecrypt.String())[keySize*3 : keySize*4]
		fifthBytes := []byte(toDecrypt.String())[keySize*4 : keySize*5]
		sixthBytes := []byte(toDecrypt.String())[keySize*5 : keySize*6]
		seventhBytes := []byte(toDecrypt.String())[keySize*6 : keySize*7]
		eighthBytes := []byte(toDecrypt.String())[keySize*7 : keySize*8]
		editDist1 := float64(CalcHammingDist(string(firstBytes), string(secondBytes))) / float64(keySize)
		editDist2 := float64(CalcHammingDist(string(thirdBytes), string(fourthBytes))) / float64(keySize)
		editDist3 := float64(CalcHammingDist(string(fifthBytes), string(sixthBytes))) / float64(keySize)
		editDist4 := float64(CalcHammingDist(string(seventhBytes), string(eighthBytes))) / float64(keySize)
		editDist := (editDist1 + editDist2 + editDist3 + editDist4) / 4.0
		editDists[keySize] = editDist
		fmt.Printf("KeySize: %d, EditDist: %f\n", keySize, editDist)
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

	// try to break the cipher with each of our three key sizes
	for _, keySize := range probableKeySizes {
		highestScore := 0
		var bestBlock []byte
		var bestKey int
		for i := 0; i < keySize; i++ {
			var block []byte
			for j := 0; j < len([]byte(toDecrypt.String())); j++ {
				if j%keySize == i {
					block = append(block, []byte(toDecrypt.String())[j])
				}
			}
			decrypted, key := SingleByteXor(hex.EncodeToString(block))
			score := ScoreString(decrypted)
			if score > highestScore {
				highestScore = score
				bestBlock = []byte(decrypted)
				bestKey = key
			}
		}
		fmt.Printf("%s\nKey: %d\n\n\n", string(bestBlock), bestKey)
	}

	fmt.Println(probableKeySizes)
	return ""
}
