package set_two

import (
	"cryptopals/shared"
	"encoding/base64"
	"strings"
	"testing"
)

func TestEncryptAESInCBC(t *testing.T) {
	input := shared.ImportTxtFile("./../set1/data/funkymusic.txt")
	gotBytes := EncryptAESInCBC([]byte(input), "YELLOW SUBMARINE")
	gotStr := base64.StdEncoding.EncodeToString(gotBytes)

	wantNewlines := shared.ImportTxtFile("./data/s2-10.txt")
	want := strings.Replace(wantNewlines, "\n", "", -1)

	if gotStr != want {
		t.Errorf("\ngot %s,\nwant %s", gotStr, want)
	}
}

func TestECBorCBCOracle(t *testing.T) {
	got := ECBorCBCOracle("Sample input. How is everybody doing on this fine January afternoon?", 16)
	want := make([]byte, 16)

	if !shared.CompareByteArrs(got, want) {
		t.Errorf("\ngot %v,\nwant %v", got, want)
	}
}