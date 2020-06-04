package core

//go:generate mockgen -destination ../mock/config.go -package mock . Config

type Config interface {
	CanvasWidth() int
	CanvasHeight() int
	ChartTitle() string
	ChartHeight() int
	ChartWidth() int
	LegendTitle() string
	LegendPosition() Position
	YAxisTitle() string
	XAxisTitle() string
}
