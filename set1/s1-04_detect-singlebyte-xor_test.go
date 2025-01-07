package set1

import "testing"

func TestDetectXor(t *testing.T) {
	got := DetectXor("./data/s1-04.txt")
	want := "Now that the party is jumping\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}