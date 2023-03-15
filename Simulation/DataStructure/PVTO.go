package datastructure

type PVTO struct {
	RS              []float64
	POIL            []float64
	FVFO            []float64
	VISO            []float64
	DENSITYOIL      []float64
	COMPRESSIBILITY []float64
}

func NewPVTO() PVTO {
	return PVTO{
		RS:              make([]float64, 0),
		POIL:            make([]float64, 0),
		FVFO:            make([]float64, 0),
		VISO:            make([]float64, 0),
		DENSITYOIL:      make([]float64, 0),
		COMPRESSIBILITY: make([]float64, 0),
	}
}
