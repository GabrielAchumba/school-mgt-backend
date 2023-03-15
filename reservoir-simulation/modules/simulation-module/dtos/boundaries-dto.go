package dtos

type Boundary struct {
	BoundaryCondition string  `json:"boundaryCondition"`
	FlowRate          float64 `json:"flowRate"`
	PressureGradient  float64 `json:"pressureGradient"`
	ConstantPressure  float64 `json:"constantPressure"`
	I                 int     `json:"i"`
	J                 int     `json:"j"`
	K                 int     `json:"k"`
}

type Boundaries struct {
	TypeOfBoundaries string     `json:"typeOfBoundaries"`
	West             []Boundary `json:"west"`
	East             []Boundary `json:"east"`
	South            []Boundary `json:"south"`
	North            []Boundary `json:"north"`
	Top              []Boundary `json:"top"`
	Bottom           []Boundary `json:"bottom"`
}
