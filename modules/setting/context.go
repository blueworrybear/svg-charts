package setting

import (
	"github.com/blueworrybear/svg-charts/core"
	"github.com/blueworrybear/svg-charts/modules/series"
)

const globalPadding = 10

type baseContext struct {
	core.Context
	options core.Options
	series  []core.Series
}

func NewContext(options core.Options) (core.Context, error) {
	var err error
	seriesList := make([]core.Series, len(options.Series))
	for i, option := range options.Series {
		if seriesList[i], err = series.NewSeries(option); err != nil {
			return nil, err
		}
	}
	return &baseContext{
		options: options,
		series:  seriesList,
	}, nil
}

func (c *baseContext) CanvasSize() (int, int) {
	return c.options.Chart.Width, c.options.Chart.Height
}

func (c *baseContext) Title() string {
	return c.options.Title.Text
}

func (c *baseContext) TitleFontSize() int {
	return c.options.Title.FontSize
}

func (c *baseContext) ChartSize() (int, int) {
	switch c.options.Chart.Type {
	case core.TypeTreeMap:
		return treeMapChartSize(c.options)
	default:
		return c.CanvasSize()
	}
}
