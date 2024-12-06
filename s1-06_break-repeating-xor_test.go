package set_one

import "testing"

func TestCalcHammingDist(t *testing.T) {
	got := CalcHammingDist("this is a test", "wokka wokka!!!")
	want := 37

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestBreakRepeatingXor(t *testing.T) {
	got := BreakRepeatingXor("./data/s1-06.txt")
	want := `WHO KNOWS!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
