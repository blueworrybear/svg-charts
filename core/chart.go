package core

import (
	"io"
)

type ChartGraphic interface {
	Context() Context
	Render(w io.Writer) error
}
