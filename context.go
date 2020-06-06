package svgchart

import (
	"github.com/blueworrybear/svg-charts/core"
)

const globalPadding = 10

type baseContext struct {
	core.Context
	options Options
	series  []core.Series
}

func newContext(options Options) (core.Context, error) {
	var err error
	seriesList := make([]core.Series, len(options.Series))
	for i, option := range options.Series {
		if seriesList[i], err = newSeries(option); err != nil {
			return nil, err
		}
	}
	return &baseContext{
		options: options,
		series:  seriesList,
	}, nil
}

func (c *baseContext) CanvasSize() (int, int) {
	if c.options.Chart.Width <= 0 || c.options.Chart.Height <= 0 {
		return 800, 600
	}
	return c.options.Chart.Width, c.options.Chart.Height
}

func (c *baseContext) Title() string {
	return c.options.Title.Text
}

func (c *baseContext) TitleFontSize() int {
	return c.options.Title.FontSize
}

func (c *baseContext) ChartSize() (int, int) {
	chartType, err := c.options.Chart.ChartType()
	if err != nil {
		return c.CanvasSize()
	}
	switch chartType {
	case core.TypeTreeMap:
		return treeMapChartSize(c.options)
	default:
		return c.CanvasSize()
	}
}

func (c *baseContext) Labels() []string {
	return c.options.Labels
}

func (c *baseContext) LabelFontSize() int {
	if c.options.LabelOption.FontSize <= 0 {
		return 12
	}
	return c.options.LabelOption.FontSize
}

func (c *baseContext) LabelFontColor() core.Hex {
	if c.options.LabelOption.Color == "" {
		return "#000000"
	}
	return core.Hex(c.options.LabelOption.Color)
}

func (c *baseContext) SeriesCount() int {
	return len(c.options.Series)
}

func (c *baseContext) Series(index int) core.Series {
	return c.series[index]
}

func (c *baseContext) FirstSeries() core.Series {
	return c.series[0]
}

func (c *baseContext) SeriesColors(index int) []core.Hex {
	s := c.series[index]
	colors := s.Colors()
	if len(colors) > 0 {
		return colors
	}
	colors = s.Colors()
	if len(colors) > 0 {
		return colors
	}
	return []core.Hex{"#000000"}
}

func treeMapChartSize(options Options) (int, int) {
	w, h := options.Chart.Width, options.Chart.Height
	w -= globalPadding
	h -= globalPadding
	if !options.Title.Hide {
		h -= options.Title.FontSize
	}
	return w, h
}
