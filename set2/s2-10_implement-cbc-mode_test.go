package set_two

import "testing"

func TestImplementCBCMode(t *testing.T) {
	got := ImplementCBCMode()
	want := "YELLOW SUBMARINE\x04\x04\x04\x04"

	if got != want {
		t.Errorf("\ngot %q,\nwant %q", got, want)
	}
}
