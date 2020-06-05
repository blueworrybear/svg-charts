package core

// Options of the chart
type Options struct {
	// Series of data
	Series []*SeriesOption
	Labels []string
	// Series color
	Colors []string
	Chart  *ChartOptions
	Title  *TitleOptions
	Legend *LegendOptions
	XAxis  *AxisOptions
	YAxis  *AxisOptions
}

// ChartOptions for specified type
type ChartOptions struct {
	Type       ChartType
	Height     int
	Width      int
	BackGround string
}

type TitleOptions struct {
	Text      string
	Alignment Alignment
	FontSize  int
	Hide      bool
}

// LegendOptions for series
type LegendOptions struct {
	Position Position
	Hide     bool
}

// AxisOptions for chart
type AxisOptions struct {
	Title string
}
