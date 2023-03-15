package datastructure

type SOW struct {
	SOIL []float64
	KROW []float64
}

func NewSOW() SOW {
	return SOW{
		SOIL: make([]float64, 0),
		KROW: make([]float64, 0),
	}
}
