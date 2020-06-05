package core

type ChartType int
type Position int
type Alignment int

const (
	TypeTreeMap ChartType = iota
)

const (
	PositionTop Position = iota
	PositionBottom
	PositionLeft
	PositionRight
)

const (
	AlignmentLeft Alignment = iota
	AlignmentRight
	AlignmentCenter
)

func (t ChartType) String() string {
	switch t {
	case TypeTreeMap:
		return "treemap"
	default:
		return "Unknown"
	}
}
