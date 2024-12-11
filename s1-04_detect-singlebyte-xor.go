package set_one

import (
	"bufio"
	"os"
)

// import the txt file as an array of strings
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func DetectXor(fileDir string) string {
	strings, err := readLines(fileDir)
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
