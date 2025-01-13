package set1

import (
	"cryptopals/shared"
	"encoding/hex"
)

// Input many lines of hex, detect which one is most likely to be encrypted using AES in ECB, returns
// the line index of the most likely line as an int
func DetectAesInECB(ciphertextPath string, key string, blockSize int) int {
	hexLines, err := shared.ReadLines(ciphertextPath) // import file to an array of strings, or "lines"
	if err != nil {
		panic(err)
	}
	idxsBlocksRepeated := []int{}
	for idx, line := range hexLines {
		bytesLine, err := hex.DecodeString(line) // hex to bytes
		if err != nil {
			panic(err)
		}
		// split each line into blocks of size blockSize
		nBlocks := len(bytesLine) / blockSize
		blocksArr := make([][]byte, nBlocks)
		for i, bs, be := 0, 0, blockSize; bs < len(bytesLine); i, bs, be = i + 1, bs + blockSize, be + blockSize {
			blocksArr[i] = bytesLine[bs:be]
		}
		// check each block array for repeated blocks. This will indicate lines likely to be encrypted using ECB
		for i, block := range blocksArr {
			if i == nBlocks {
				break
			}
			for j := i + 1; j < nBlocks; j++ {
				if shared.CompareByteArrs(block, blocksArr[j]) {
					idxsBlocksRepeated = append(idxsBlocksRepeated, idx) // contains line numbers, one per detected repetition
				}
			}
		}
	}
	// Create a map of all line numbers which contained repeated blocks
	// Also count how many repeated blocks were present in each of these lines
	repeatedBlocksCountMap := make(map[int]int)
	for _, idx := range idxsBlocksRepeated {
		value, ok := repeatedBlocksCountMap[idx]
		if ok {
			repeatedBlocksCountMap[idx] = value + 1
		} else {
			repeatedBlocksCountMap[idx] = 1
		}
	}
	
	// Which line number from the original file had the most repeated blocks?
	highestCount := 0
	mostRepeatedIdx := 0
	for idx, count := range repeatedBlocksCountMap {
		if count > highestCount {
			mostRepeatedIdx = idx
		}
	}
	return mostRepeatedIdx
}
