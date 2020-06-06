package svgchart

import (
	"errors"
	"fmt"

	"github.com/blueworrybear/svg-charts/core"
)

var errSeriesFormat = errors.New("")

func SeriesData(data ...interface{}) []interface{} {
	return data
}

type baseSeries struct {
	core.Series
	name   string
	data   []interface{}
	colors []core.Hex
}

func newSeries(option *SeriesOption) (core.Series, error) {
	colors := make([]core.Hex, len(option.Colors))
	for i, v := range option.Colors {
		colors[i] = core.Hex(v)
	}
	return &baseSeries{
		name:   option.Name,
		data:   option.Data,
		colors: colors,
	}, nil
}

func (s *baseSeries) Name() string {
	return s.name
}

func (s *baseSeries) Data() []interface{} {
	return s.data
}

func (s *baseSeries) Float64Data() ([]float64, error) {
	values := make([]float64, len(s.data))
	for i, v := range s.data {
		switch n := v.(type) {
		case float64:
			values[i] = n
		case int:
			values[i] = float64(n)
		default:
			return values, fmt.Errorf("Value format error")
		}
	}
	return values, nil
}

func (s *baseSeries) Colors() []core.Hex {
	return s.colors
}
