package charts

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	svg "github.com/ajstarks/svgo"
	"github.com/blueworrybear/svg-charts/core"
	"github.com/blueworrybear/svg-charts/mock"
	"github.com/golang/mock/gomock"
)

func TestTreeMap(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	intNums := []int{100, 20, 30, 40, 50, 60, 40, 70, 80}
	colors := []core.Hex{"#C7243A", "#7A7616", "#7A0A19", "#165B7A", "#068AC7"}
	floatNums := make([]float64, len(intNums))
	seriesColors := make([]core.Hex, len(intNums))
	for i, v := range intNums {
		floatNums[i] = float64(v)
		seriesColors[i] = colors[i%len(colors)]
	}
	labels := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}

	mockSeries := mock.NewMockSeries(ctrl)
	mockSeries.EXPECT().Float64Data().Return(floatNums, nil)

	mockContext := mock.NewMockContext(ctrl)
	mockContext.EXPECT().FirstSeries().Return(mockSeries)
	mockContext.EXPECT().SeriesColors(gomock.Eq(0)).Return(seriesColors)
	mockContext.EXPECT().Labels().Return(labels)
	mockContext.EXPECT().LabelColor().AnyTimes().Return(core.Hex("#FF0000"))
	mockContext.EXPECT().LabelFontSize().AnyTimes().Return(15)
	c := NewTreeMapChart(mockContext)

	file, err := os.Create("tree.svg")
	defer file.Close()
	if err != nil {
		t.Error(err)
		return
	}
	buffer := bytes.NewBuffer([]byte{})
	if err := c.Render(buffer); err != nil {
		t.Error(err)
		return
	}
	canvas := svg.New(file)
	canvas.Start(800, 600)
	canvas.Gtransform(`translate(10, 10)`)
	canvas.Writer.Write(buffer.Bytes())
	canvas.Gend()
	canvas.End()
}

func TestTilingSlice(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50, 60}
	s := make([]float64, len(nums))
	for i, v := range nums {
		s[i] = float64(v)
	}

	root := tilingSlice(s)
	stack := make([]*treeNode, 0)
	leaves := make([]float64, 0)
	cur := root
	for {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		if len(stack) <= 0 {
			break
		}
		cur, stack = stack[len(stack)-1], stack[0:len(stack)-1]
		if cur.left == nil && cur.right == nil {
			leaves = append(leaves, cur.val)
		}
		cur = cur.right
	}
	for i, v := range s {
		if v != leaves[i] {
			t.Fail()
		}
	}
	if root.val != root.left.val+root.right.val {
		t.Fail()
	}
	for _, box := range tilingBoxes(root, 0, 0, 100, 100, 0) {
		fmt.Printf("%v\n", box)
	}
}
