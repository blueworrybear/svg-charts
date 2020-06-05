package setting

import "github.com/blueworrybear/svg-charts/core"

func treeMapChartSize(options core.Options) (int, int) {
	w, h := options.Chart.Width, options.Chart.Height
	w -= globalPadding
	h -= globalPadding
	if !options.Title.Hide {
		h -= options.Title.FontSize
	}
	return w, h
}
