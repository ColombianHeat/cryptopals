package set_two

func ImplementPKCSPadding(unpadded string, blockSize int) string {
	padded := unpadded
	if len(unpadded)%blockSize != 0 { // if true, we need to add padding
		padSize := blockSize - len(unpadded)%blockSize
		for i := 0; i < padSize; i++ {
			// Add ASCII 4 to the end of the unpadded text until length is a multiple of blockSize
			padded += "\x04"
		}
		return padded
	}
	return padded // No changes necessary if length of unpadded text is a multiple of blockSize
}
