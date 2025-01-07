package set_two

import "testing"

func TestImplementPKCSPadding(t *testing.T) {
	got := ImplementPKCSPadding("YELLOW SUBMARINE", 20)
	want := "YELLOW SUBMARINE\x04\x04\x04\x04"

	if got != want {
		t.Errorf("\ngot %q,\nwant %q", got, want)
	}

	got = ImplementPKCSPadding("Alan is cool", 5)
	want = "Alan is cool\x04\x04\x04"

	if got != want {
		t.Errorf("\ngot %q,\nwant %q", got, want)
	}
}
