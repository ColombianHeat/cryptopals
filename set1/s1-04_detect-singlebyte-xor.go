package set1

import (
	"cryptopals/shared"
)

func DetectXor(fileDir string) string {
	strings, err := shared.ReadLines(fileDir)
	if err != nil {
		panic(err)
	}
	highestScore := 0
	highestScoreIdx := 0
	// decrypt and score each string
	for idx, str := range strings {
		decrypted_str, _ := SingleByteXor(str)
		score := ScoreString(decrypted_str)
		if score > highestScore {
			highestScore = score
			highestScoreIdx = idx
		}
	}
	// return the decrypted string with the highest score
	winningStr := strings[highestScoreIdx] // "Now that the party is jumping\n"
	decrypted, _ := SingleByteXor(winningStr)
	return decrypted
}
