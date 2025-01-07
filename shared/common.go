package shared

import (
	"bufio"
	"os"
	"strings"
)

// import the txt file as an array of strings
func ReadLines(path string) ([]string, error) {
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

// import txt file as one long string
func ImportTxtFile(path string) string {
	lines, err := ReadLines(path)
	if err != nil {
		panic(err)
	}
	var contents strings.Builder
	for _, line := range lines {
		contents.WriteString(line)
	}
	return contents.String()
}

// converts a slice of integers to a slice of bytes.
func IntsToBytes(ints []int) []byte {
	bytes := make([]byte, len(ints))
	for i, v := range ints {
		bytes[i] = byte(v) // Safely cast each integer to a byte
	}
	return bytes
}
