package set_two

import "testing"

func TestImplementPKCSPadding(t *testing.T) {
	got := ImplementPKCSPadding()
	want := "something"

	if got != want {
		t.Errorf("\ngot %q,\nwant %q", got, want)
	}
}
