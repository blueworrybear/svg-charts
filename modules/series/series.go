package series

import (
	"errors"
	"fmt"

	"github.com/blueworrybear/svg-charts/core"
)

var errSeriesFormat = errors.New("")

type baseSeries struct {
	core.Series
	name   string
	data   []interface{}
	colors []string
}

func NewSeries(option *core.SeriesOption) (core.Series, error) {
	return &baseSeries{
		name:   option.Name,
		data:   option.Data,
		colors: option.Colors,
	}, nil
}

func (s *baseSeries) Name() string {
	return s.name
}

func (s *baseSeries) Data() []interface{} {
	return s.data
}

func (s *baseSeries) Float64Data() ([]float64, error) {
	var ok bool
	values := make([]float64, len(s.data))
	for i, v := range s.data {
		if values[i], ok = v.(float64); !ok {
			return values, fmt.Errorf("Value format error")
		}
	}
	return values, nil
}

func (s *baseSeries) Colors() []core.RGB {
	colors := make([]core.RGB, len(s.colors))
	for i, hex := range s.colors {
		colors[i], _ = core.Hex2RGB(core.Hex(hex))
	}
	return colors
}
