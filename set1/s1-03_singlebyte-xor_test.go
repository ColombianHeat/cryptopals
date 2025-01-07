package set1

import "testing"

func TestSingleByteXor(t *testing.T) {
	got, _ := SingleByteXor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	want := "Cooking MC's like a pound of bacon"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
