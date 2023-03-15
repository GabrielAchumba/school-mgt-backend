package datastructure

type PVTW struct {
	PRES            []float64
	FVF             []float64
	COMPRESSIBILITY []float64
	VISCOSITY       []float64
	DENSITYWATER    []float64
}

func NewPVTW() PVTW {
	return PVTW{
		PRES:            make([]float64, 0),
		FVF:             make([]float64, 0),
		COMPRESSIBILITY: make([]float64, 0),
		VISCOSITY:       make([]float64, 0),
		DENSITYWATER:    make([]float64, 0),
	}
}
