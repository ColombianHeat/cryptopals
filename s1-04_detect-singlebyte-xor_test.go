package set_one

import "testing"

func TestDetectXor(t *testing.T) {
	got := DetectXor("DIR GOES HERE")
	want := "secret message"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}