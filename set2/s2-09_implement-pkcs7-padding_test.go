package set_two

import (
	"cryptopals/shared"
	"testing"
)

func TestImplementPKCSPadding(t *testing.T) {
	got := ImplementPKCSPadding([]byte("YELLOW SUBMARINE"), 20)
	want := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")

	if ! shared.CompareByteArrs(got, want) {
		t.Errorf("\ngot %q,\nwant %q", got, want)
	}

	got = ImplementPKCSPadding([]byte("Alan is cool"), 5)
	want = []byte("Alan is cool\x04\x04\x04")

	if ! shared.CompareByteArrs(got, want) {
		t.Errorf("\ngot %q,\nwant %q", got, want)
	}
}
