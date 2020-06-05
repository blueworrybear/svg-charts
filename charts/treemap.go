package charts

import (
	"fmt"
	"io"

	svg "github.com/ajstarks/svgo"
	"github.com/blueworrybear/svg-charts/core"
	shape "github.com/blueworrybear/svg-charts/modules/common"
)

type TreeMapChart struct {
	contex core.Context
	core.ChartGraphic
}

func NewTreeMapChart(context core.Context) *TreeMapChart {
	return &TreeMapChart{
		contex: context,
	}
}

func (c *TreeMapChart) Context() core.Context {
	return c.contex
}

func (c *TreeMapChart) Render(w io.Writer) error {
	canvas := svg.New(w)
	s := c.contex.FirstSeries()
	data, err := s.Float64Data()
	colors := c.contex.SeriesColors(0)
	if err != nil {
		return err
	}
	root := tilingSlice(data)
	for i, box := range tilingBoxes(root, 0, 0, 600, 400, 0) {
		boxID := fmt.Sprintf("box%d", box.id)
		rect := shape.NewRectangle(box.w, box.h)
		rect.SetPosition(box.x, box.y)
		rect.SetPadding(1)
		rect.SetColor(colors[i])
		rect.SetID(boxID)
		rect.Draw(canvas)
	}
	return nil
}

type treeNode struct {
	left  *treeNode
	right *treeNode
	val   float64
	id    int
}

func tilingSlice(s []float64) *treeNode {
	nodes := make([]*treeNode, len(s))
	for i, v := range s {
		nodes[i] = &treeNode{
			val: v,
			id:  i,
		}
	}
	return tilingNodes(nodes)
}

func tilingNodes(nodes []*treeNode) *treeNode {
	if len(nodes) <= 0 {
		return nil
	} else if len(nodes) <= 1 {
		return nodes[0]
	}
	m := (len(nodes) + 1) / 2
	left := tilingNodes(nodes[0:m])
	right := tilingNodes(nodes[m:])
	s := 0.0
	if left != nil {
		s += left.val
	}
	if right != nil {
		s += right.val
	}
	return &treeNode{
		left:  left,
		right: right,
		val:   s,
	}
}

type treeBox struct {
	id int
	x  int
	y  int
	w  int
	h  int
}

func tilingBoxes(root *treeNode, x, y, w, h int, level int) []*treeBox {
	if root == nil {
		return []*treeBox{}
	} else if root.left == nil && root.right == nil {
		return []*treeBox{{
			id: root.id,
			x:  x,
			y:  y,
			w:  w,
			h:  h,
		}}
	}
	a := root.left.val / root.val
	boxes := make([]*treeBox, 0)
	if level%2 == 0 {
		s := int(float64(w) * a)
		boxes = append(boxes, tilingBoxes(root.left, x, y, s, h, level+1)...)
		boxes = append(boxes, tilingBoxes(root.right, x+s, y, w-s, h, level+1)...)
	} else {
		s := int(float64(h) * a)
		boxes = append(boxes, tilingBoxes(root.left, x, y, w, s, level+1)...)
		boxes = append(boxes, tilingBoxes(root.right, x, y+s, w, h-s, level+1)...)
	}
	return boxes
}
