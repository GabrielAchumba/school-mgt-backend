package mathematicslibrary

import (
	"fmt"

	mat "gonum.org/v1/gonum/mat"
)

type LinearEquations struct {
}

func NewLinearEquations() LinearEquations {
	return LinearEquations{}
}

func (impl *LinearEquations) Solver(A [][]float64, B []float64) []float64 {

	//var x mat.Dense
	r := len(A)
	c := len(A[0])

	a := mat.NewDense(r, c, nil)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			a.Set(i, j, A[i][j])
		}
	}

	b := mat.NewVecDense(r, nil)
	for i := 0; i < r; i++ {
		b.SetVec(i, B[i])
	}

	var lu mat.LU
	lu.Factorize(a)
	var x mat.VecDense
	//x := mat.NewVecDense(r, nil)
	if err := lu.SolveVecTo(&x, false, b); err != nil {
		fmt.Println("Error encountered")
	}

	return x.RawVector().Data
}
