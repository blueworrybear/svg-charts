package core

import "testing"

func TestHex2RGB(t *testing.T) {
	rgb, err := Hex2RGB("#FF0000")
	if err != nil {
		t.Error(err)
		return
	}
	if rgb.R != 255 || rgb.G != 0 || rgb.B != 0 {
		t.Fail()
	}
}
