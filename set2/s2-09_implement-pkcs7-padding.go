package set_two

func ImplementPKCSPadding(unpadded []byte, blockSize int) []byte {
	padded := string(unpadded)
	if len(unpadded)%blockSize != 0 { // if true, we need to add padding
		padSize := blockSize - len(unpadded)%blockSize
		for i := 0; i < padSize; i++ {
			// Add ASCII 4 to the end of the unpadded text until length is a multiple of blockSize
			padded += "\x04"
		}
		return []byte(padded)
	}
	return []byte(padded) // No changes necessary if length of unpadded text is a multiple of blockSize
}
