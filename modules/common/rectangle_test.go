package common

import "testing"

func TestRectangleStyle(t *testing.T) {
	r := NewRectangle(10, 10)
	r.AddAttr("fill")
	r.AddAttr("stroke")
	if r.AttrString() != "fill;stroke" {
		t.Fail()
	}
	r.ClearAttr()
	if r.AttrString() != "" {
		t.Fail()
	}
	r.AddAttr("fill")
	if r.AttrString() != "fill" {
		t.Fail()
	}

}
