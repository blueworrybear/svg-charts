package core

import (
	"io"

	svg "github.com/ajstarks/svgo"
)

type ChartGraphic interface {
	Context() Context
	Render(w io.Writer) error
}

type CommonGraphic interface {
	Draw(canvas *svg.SVG) error
}
