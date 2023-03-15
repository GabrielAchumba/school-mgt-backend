package datastructure

type SOG struct {
	SOIL []float64
	KROG []float64
}

func NewSOG() SOG {
	return SOG{
		SOIL: make([]float64, 0),
		KROG: make([]float64, 0),
	}
}
