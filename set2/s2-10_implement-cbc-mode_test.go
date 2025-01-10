package set_two

import (
	"testing"
)

func TestImplementCBCMode(t *testing.T) {
	plainText := ImplementPKCSPadding("This is my secret message. Read it at your own peril. Volcano...", 16)
	encrypted := EncryptAesInECB(plainText, "YELLOW SUBMARINE") // encrypt plaintext
	decrypted := DecryptAesInECB(encrypted, "YELLOW SUBMARINE") // decrypt plaintext
	got := decrypted // Should be a padded version of original plaintext if both encrypt and decrypt fncs work
	want := "This is my secret message. Read it at your own peril. Volcano..."

	if got != want {
		t.Errorf("\ngot %q,\nwant %q", got, want)
	}
}
