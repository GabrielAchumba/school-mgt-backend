package datastructure

type SGFN struct {
	SGAS []float64
	KRG  []float64
	PCOG []float64
	KROG []float64
}

func NewSGFN() SGFN {
	return SGFN{
		SGAS: make([]float64, 0),
		KRG:  make([]float64, 0),
		PCOG: make([]float64, 0),
		KROG: make([]float64, 0),
	}
}
