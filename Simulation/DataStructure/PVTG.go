package datastructure

type PVTG struct {
	PGAS            []float64
	BGAS            []float64
	VISGAS          []float64
	DENGAS          []float64
	COMPRESSIBILITY []float64
}

func NewPVTG() PVTG {
	return PVTG{
		PGAS:            make([]float64, 0),
		BGAS:            make([]float64, 0),
		VISGAS:          make([]float64, 0),
		DENGAS:          make([]float64, 0),
		COMPRESSIBILITY: make([]float64, 0),
	}
}
