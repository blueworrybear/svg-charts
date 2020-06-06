package common

import (
	"fmt"

	svg "github.com/ajstarks/svgo"
	"github.com/blueworrybear/svg-charts/core"
)

type TextLabel struct {
	text     string
	x        int
	y        int
	fontSize int
	color    core.Hex
}

func NewTextLabel(x, y int, text string) *TextLabel {
	return &TextLabel{
		x:        x,
		y:        y,
		text:     text,
		fontSize: 12,
		color:    "#000000",
	}
}

func (label *TextLabel) SetColor(hex core.Hex) {
	label.color = hex
}

func (label *TextLabel) SetFontSize(size int) {
	label.fontSize = size
}

func (label *TextLabel) Draw(canvas *svg.SVG) error {
	fontSize := fmt.Sprintf(`font-size="%dpx"`, label.fontSize)
	color, _ := core.Hex2RGB(label.color)
	x, y := label.x, label.y
	y += label.fontSize
	canvas.Text(
		x, y, label.text, fontSize,
		canvas.RGB(color.R, color.G, color.B))
	return nil
}
