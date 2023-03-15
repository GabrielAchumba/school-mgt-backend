package datastructure

type BoundaryType int

const (
	Closed BoundaryType = iota + 1
	ConstantGradient
	KnownFlowRate
	ConstantPressure
)

type Boundary struct {
	West   West
	East   East
	South  South
	North  North
	Top    Top
	Bottom Bottom
}

func NewBoundary() Boundary {
	return Boundary{
		West:   NewWest(),
		East:   NewEast(),
		South:  NewSouth(),
		North:  NewNorth(),
		Top:    NewTop(),
		Bottom: NewBottom(),
	}
}

type West struct {
	BoundaryType     BoundaryType
	PressureGradient float64
	FlowRate         float64
	ConstantPressure float64
}

func NewWest() West {
	return West{
		BoundaryType: Closed,
	}
}

type East struct {
	BoundaryType     BoundaryType
	PressureGradient float64
	FlowRate         float64
	ConstantPressure float64
}

func NewEast() East {
	return East{
		BoundaryType: Closed,
	}
}

type South struct {
	BoundaryType     BoundaryType
	PressureGradient float64
	FlowRate         float64
	ConstantPressure float64
}

func NewSouth() South {
	return South{
		BoundaryType: Closed,
	}
}

type North struct {
	BoundaryType     BoundaryType
	PressureGradient float64
	FlowRate         float64
	ConstantPressure float64
}

func NewNorth() North {
	return North{
		BoundaryType: Closed,
	}
}

type Top struct {
	BoundaryType     BoundaryType
	PressureGradient float64
	FlowRate         float64
	ConstantPressure float64
}

func NewTop() Top {
	return Top{
		BoundaryType: Closed,
	}
}

type Bottom struct {
	BoundaryType     BoundaryType
	PressureGradient float64
	FlowRate         float64
	ConstantPressure float64
}

func NewBottom() Bottom {
	return Bottom{
		BoundaryType: Closed,
	}
}
