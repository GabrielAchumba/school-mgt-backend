package datastructure

type SWFN struct {
	SWAT []float64
	KRW  []float64
	PCOW []float64
	KROW []float64
}

func NewSWFN() SWFN {
	return SWFN{
		SWAT: make([]float64, 0),
		KRW:  make([]float64, 0),
		PCOW: make([]float64, 0),
		KROW: make([]float64, 0),
	}
}
