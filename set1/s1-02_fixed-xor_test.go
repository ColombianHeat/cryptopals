package set1

import "testing"

func TestXorBuffers(t *testing.T) {
	got := XorBuffers("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	want := "746865206b696420646f6e277420706c6179"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
