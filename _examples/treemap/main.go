package main

import (
	"log"
	"os"

	charts "github.com/blueworrybear/svg-charts"
)

func main() {
	file, _ := os.Create("test.svg")
	svg, err := charts.NewSVGChart(charts.Options{
		Series: []*charts.SeriesOption{
			{
				Data:   charts.SeriesData(10, 20, 30),
				Colors: []string{"#FF0000", "#00FF00", "#0000FF"},
			},
		},
		Chart: &charts.ChartOptions{
			Type: "treemap",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if err := svg.Render(file); err != nil {
		log.Fatal(err)
	}
}
