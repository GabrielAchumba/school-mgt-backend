package mathematicslibrary

type SimplexSnap struct {
	B       []float64
	Matrix  [][]float64
	M       []float64
	F       []float64
	CC      []int
	FValue  float64
	FVars   []float64
	IsMDone bool
	MM      []bool
}

func NewSimplexSnap() SimplexSnap {
	return SimplexSnap{}
}

func NewSimplexSnap2(_b []float64, _matrix [][]float64, _M []float64, _F []float64,
	_CC []int, _fVars []float64, _isMDone bool, _m []bool) SimplexSnap {

	simplexSnap := SimplexSnap{}
	simplexSnap.B = _b
	simplexSnap.Matrix = _matrix
	simplexSnap.M = _M
	simplexSnap.F = _F
	simplexSnap.CC = _CC
	simplexSnap.IsMDone = _isMDone
	simplexSnap.MM = _m
	simplexSnap.FVars = _fVars
	simplexSnap.FValue = 0
	CCCount := len(simplexSnap.CC)
	for i := 0; i < CCCount; i++ {
		simplexSnap.FValue += simplexSnap.FVars[simplexSnap.CC[i]] * simplexSnap.B[i]
	}

	return simplexSnap
}
