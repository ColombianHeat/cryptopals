package set_two

import (
	"cryptopals/shared"
	"os"
	"testing"
)

func TestImplementCBCMode(t *testing.T) {
	plainText := ImplementPKCSPadding([]byte("This is my secret message. Read it at your own peril. Volcano..."), 16)
	encrypted := EncryptAesInECB([]byte(plainText), "YELLOW SUBMARINE") // encrypt plaintext
	decrypted := DecryptAesInECB(encrypted, "YELLOW SUBMARINE") // decrypt plaintext
	got := decrypted // Should be a padded version of original plaintext if both encrypt and decrypt fncs work
	want := []byte("This is my secret message. Read it at your own peril. Volcano...")

	if ! shared.CompareByteArrs(got, want) {
		t.Errorf("\ngot %q,\nwant %q", got, want)
	}

	got = DecryptCBCMode("./data/s2-10.txt", "YELLOW SUBMARINE")
	f, err := os.Create("./data/funkymusic.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(got)
	want = []byte(shared.ImportTxtFile("./../set1/data/funkymusic.txt"))

	if ! shared.CompareByteArrs(got, want) {
		t.Errorf("\ngot %q,\nwant %q", got, want)
	}
}
