package svgchart

import (
	"fmt"
	"io"

	svgo "github.com/ajstarks/svgo"
	"github.com/blueworrybear/svg-charts/charts"
	"github.com/blueworrybear/svg-charts/core"
)

type SVGChart struct {
	context core.Context
	options Options
}

func NewSVGChart(options Options) (*SVGChart, error) {
	context, err := newContext(options)
	if err != nil {
		return nil, err
	}
	return &SVGChart{
		options: options,
		context: context,
	}, nil
}

func (svg *SVGChart) renderChart(w io.Writer) error {
	chartType, err := svg.options.Chart.ChartType()
	if err != nil {
		return err
	}
	switch chartType {
	case core.TypeTreeMap:
		c := charts.NewTreeMapChart(svg.context)
		if err := c.Render(w); err != nil {
			return err
		}
	default:
		return fmt.Errorf("%v not support", svg.options.Chart.Type)
	}
	return nil
}

func (svg *SVGChart) Render(w io.WriteCloser) error {
	defer w.Close()
	canvas := svgo.New(w)
	canvas.Start(svg.context.CanvasSize())
	if err := svg.renderChart(w); err != nil {
		return err
	}
	canvas.End()
	return nil
}
