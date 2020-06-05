package common

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
	"github.com/blueworrybear/svg-charts/core"
)

type Rectangle struct {
	Shape
	w       int
	h       int
	x       int
	y       int
	padding int
	styles  []string
}

func NewRectangle(w, h int) *Rectangle {
	return &Rectangle{
		w: w,
		h: h,
	}
}

func (r *Rectangle) SetPosition(x, y int) {
	r.x = x
	r.y = y
}

func (r *Rectangle) SetPadding(p int) {
	r.padding = p
}

func (r *Rectangle) SetColor(hex core.Hex) {
	rgb, _ := core.Hex2RGB(hex)
	r.AddAttr(fmt.Sprintf("fill:%s", rgb.String()))
}

func (r *Rectangle) getXY() (int, int) {
	x := r.x
	y := r.y
	if r.padding > 0 {
		x += r.padding
		y += r.padding
	}
	return x, y
}

func (r *Rectangle) getWH() (int, int) {
	w := r.w
	h := r.h
	if r.padding > 0 {
		w -= 2 * r.padding
		h -= 2 * r.padding
	}
	return w, h
}

func (r *Rectangle) Draw(canvas *svg.SVG) error {
	x, y := r.getXY()
	w, h := r.getWH()
	canvas.Rect(x, y, w, h, r.Attributes()...)
	return nil
}
