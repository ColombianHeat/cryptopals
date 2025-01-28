package set_two

// ImplementPKCSPadding takes a byte slice and block size as input and returns a slice of bytes
// with the PKCS#7 padding applied. If the length of the slice is already a multiple of the block size,
// the function will return the original slice. Otherwise, it will add the appropriate number of padding
// bytes (ASCII 0x04) to the end of the slice such that the length of the slice is a multiple of the block size.
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
