package series

import (
	"errors"

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

func (s *baseSeries) Colors() []core.RGB {
	colors := make([]core.RGB, len(s.colors))
	for i, hex := range s.colors {
		colors[i], _ = core.Hex2RGB(core.Hex(hex))
	}
	return colors
}
