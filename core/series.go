package core

//go:generate mockgen -destination ../mock/series.go -package mock . Series

type Series interface {
	Name() string
	Data() []interface{}
	Colors() []RGB
}

type SeriesOption struct {
	Name   string
	Data   []interface{}
	Colors []string
}
