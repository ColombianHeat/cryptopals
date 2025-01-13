package shared

import "testing"

func TestCompareByteArrs(t *testing.T) {
	arr1 := []byte{1, 3, 7, 2}
	arr2 := []byte{1, 3, 7, 2}
	got := CompareByteArrs(arr1, arr2)
	want := true

	if got != want {
		t.Errorf("\ngot  %t \nwant %t", got, want)
	}

	arr1 = []byte{1, 3, 7, 2}
	arr2 = []byte{1, 3, 7, 3}
	got = CompareByteArrs(arr1, arr2)
	want = false

	if got != want {
		t.Errorf("\ngot  %t \nwant %t", got, want)
	}
}

func TestReadLines(t *testing.T) {
	want := []string{"testing TeStInG", "This is the second line", "and the third!"}
	got, _ := ReadLines("./data/test.txt")
	for i := 0; i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("\ngot  %q \nwant %q", got[i], want[i])
		}
	}
}

func TestImportTxtFile(t *testing.T) {
	got := ImportTxtFile("./data/test.txt")
	want := "testing TeStInG\nThis is the second line\nand the third!"

	if got != want {
		t.Errorf("\ngot %q \nwant %q", got, want)
	}
}