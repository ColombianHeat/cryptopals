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
		line = line + "\n"
		contents.WriteString(line)
	}
	output := contents.String()
	// remove trailing newline from output
	// if output[len(output)-1:] == "\n" {
	// 	output = output[:len(output)-1]
	// }
	return output
}

// converts a slice of integers to a slice of bytes.
func IntsToBytes(ints []int) []byte {
	bytes := make([]byte, len(ints))
	for i, v := range ints {
		bytes[i] = byte(v) // Safely cast each integer to a byte
	}
	return bytes
}

func XorByteVectors(a []byte, b []byte) []byte {
	if len(a) != len(b) {
		panic("buffers are not the same length")
	}
	vectorLen := len(a)
	xorBin := make([]byte, vectorLen) // empty byte slice
	for i := 0; i < vectorLen; i++ {
		xorBin[i] = a[i] ^ b[i]
	}
	return xorBin
}

// compareByteArrs takes two byte arrays and returns true if they are equal, false if not
func CompareByteArrs(arr1, arr2 []byte) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}