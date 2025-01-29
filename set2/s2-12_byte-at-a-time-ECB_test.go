package set_two

import (
	"testing"
)

func TestByteAtATimeECB(t *testing.T) {
	yourString := "THIS IS MY PLAINTEXT. I LIKE IT VERY MUCHZZZZZZZZZZZZZZZZZZZZZZZZ"
	randomKey := RandAESKey(16)
	got := string(ByteAtATimeECB(yourString, randomKey))
	want := "What I want"

	if got != want {
		t.Errorf("\ngot %s,\nwant %s", got, want)
	}
}
