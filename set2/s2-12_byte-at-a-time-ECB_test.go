package set_two

import (
	"testing"
)

func TestByteAtATimeECB(t *testing.T) {
	got := ByteAtATimeECB()
	want := "What I want"

	if got != want {
		t.Errorf("\ngot %s,\nwant %s", got, want)
	}
}
