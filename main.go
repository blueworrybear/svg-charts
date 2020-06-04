package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	Name string
	Data interface{}
}

func main() {
	t := &T{Data: []int{1, 2, 3}}
	data, _ := json.Marshal(t)
	fmt.Println(string(data))
	fmt.Println((1 + 1) / 2)
	// width := 500
	// height := 500
	// file, _ := os.OpenFile("tes	t.svg", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	// canvas := svg.New(file)
	// canvas.Start(width, height)
	// canvas.Circle(width/2, height/2, 100)
	// canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	// canvas.End()
}
