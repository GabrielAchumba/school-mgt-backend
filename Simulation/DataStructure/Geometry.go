package datastructure

type Geometry struct {
	Dx     float64
	Dy     float64
	Dz     float64
	Vb     float64
	Z      float64
	Area_x float64
	Area_y float64
	Area_z float64
}

func NewGeometry() Geometry {
	return Geometry{}
}
