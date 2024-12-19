package set_one

import "testing"

func TestDetectAesInECB(t *testing.T) {
	got := DetectAesInECB("./data/s1-08.txt", "YELLOW SUBMARINE", 16)
	want := "something"

	if got != want {
		t.Errorf("\ngot %q,\nwant %q", got, want)
	}
}
