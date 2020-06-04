package core

type ChartType int
type Position int

const (
	TypeTreeMap ChartType = iota
)

const (
	PositionTop Position = iota
	PositionBottom
	PositionLeft
	PositionRight
)

func (t ChartType) String() string {
	switch t {
	case TypeTreeMap:
		return "treemap"
	default:
		return "Unknown"
	}
}
