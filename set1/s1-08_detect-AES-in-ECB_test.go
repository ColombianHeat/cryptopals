package set1

import "testing"

func TestDetectAesInECB(t *testing.T) {
	got := DetectAesInECB("./data/s1-08.txt", "YELLOW SUBMARINE", 16)
	want := 132

	if got != want {
		t.Errorf("\ngot %q,\nwant %q", got, want)
	}
}
