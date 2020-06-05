package core

//go:generate mockgen -destination ../mock/context.go -package mock . Context

type Context interface {
	CanvasSize() (int, int)
	Title() string
	TitleFontSize() int
	ChartSize() (int, int)
	ChartOffset() (int, int)
	Labels() string
	SeriesCount() int
	Series(index int) Series
	FirstSeries() Series
	SeriesColors(index int) []Hex
	LegendTitle() string
	LegendPosition() Position
	YAxisTitle() string
	XAxisTitle() string
}
