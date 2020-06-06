package svgchart

import (
	"fmt"

	"github.com/blueworrybear/svg-charts/core"
)

// Options of the chart
type Options struct {
	// Series of data
	Series      []*SeriesOption
	Labels      []string
	LabelOption *LabelOptions
	// Series color
	Colors []string
	Chart  *ChartOptions
	Title  *TitleOptions
	Legend *LegendOptions
	XAxis  *AxisOptions
	YAxis  *AxisOptions
}

type SeriesOption struct {
	Name   string
	Data   []interface{}
	Colors []string
}

// ChartOptions for specified type
type ChartOptions struct {
	Type       string
	Height     int
	Width      int
	BackGround string
}

func (option *ChartOptions) ChartType() (core.ChartType, error) {
	switch option.Type {
	case "treemap":
		return core.TypeTreeMap, nil
	default:
		return 0, fmt.Errorf("Unsupported chart type %s", option.Type)
	}
}

type TitleOptions struct {
	Text      string
	Alignment core.Alignment
	FontSize  int
	Hide      bool
}

// LegendOptions for series
type LegendOptions struct {
	Position core.Position
	Hide     bool
}

// AxisOptions for chart
type AxisOptions struct {
	Title string
}

type LabelOptions struct {
	FontSize int
	Color    string
}
