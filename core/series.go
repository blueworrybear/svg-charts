package core

//go:generate mockgen -destination ../mock/series.go -package mock . Series

type SeriesGroup []*Series

type Series interface {
	Name() string
	Data() []interface{}
	Float64Data() ([]float64, error)
	Colors() []Hex
}

func (g SeriesGroup) Size() int {
	return len(g)
}

func (g SeriesGroup) Series(i int) *Series {
	return g[i]
}
