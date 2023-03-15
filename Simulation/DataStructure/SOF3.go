package datastructure

type SOF3 struct {
	SOIL []float64
	KROW []float64
	KROG []float64
}

func NewSOF3() SOF3 {
	return SOF3{
		SOIL: make([]float64, 0),
		KROW: make([]float64, 0),
		KROG: make([]float64, 0),
	}
}
