package mathematicslibrary

import (
	"math"
)

type NonLinearEquations struct {
	differentiation *Differentiation
	interpolation   *Interpolation
}

func NewNonLinearEquations() (nonLinearEquations *NonLinearEquations) {
	nonLinearEquations = new(NonLinearEquations)
	nonLinearEquations.differentiation = NewDifferentiation()
	nonLinearEquations.interpolation = NewInterpolation()
	return
}

func (nonLinearEquations *NonLinearEquations) NewtonRaphson(fun ObjectiveFunc1, x0 float64) float64 {
	Tol := 10e-6
	MaxtIter := 100
	var diff float64 = 0

	x1 := x0
	var StepSize float64
	var ARE float64 = 10000
	Tol2 := math.Abs(Tol * 100.0)

	for i := 0; i <= MaxtIter; i++ {
		StepSize = x0 / math.Pow(10, 3)

		diff = nonLinearEquations.differentiation.Central(fun,
			x0, StepSize)

		if diff == 0 {
			diff = 10e-5
		}

		x1 = x0 - fun(x0)/diff

		ARE = math.Abs((x1-x0)/x1) * 100.0

		if ARE > Tol2 {
			x0 = x1
		} else {
			break
		}
	}

	return x0

}

func (nonLinearEquations *NonLinearEquations) GetValue(x []float64, y []float64,
	x1 float64) float64 {
	a := nonLinearEquations.interpolation.LinearInterpolation(x, y, x1)
	return a
}

func (nonLinearEquations *NonLinearEquations) ObjFunc(x1 []float64, y1 []float64,
	x2 []float64, y2 []float64,
	xx float64) float64 {

	aa := nonLinearEquations.GetValue(x1, y1, xx)
	bb := nonLinearEquations.GetValue(x2, y2, xx)
	Residual := aa - bb
	return Residual
}
func (nonLinearEquations *NonLinearEquations) Bisectionfzero(x2 []float64,
	y2 []float64, x1 []float64, y1 []float64, xL float64, xU float64) float64 {

	nmax := 500
	var xr, err1 float64
	err1 = 10e-6
	for i := 1; i <= nmax; i++ {
		xr = (xL + xU) / 2.0
		fxu := nonLinearEquations.ObjFunc(x1, y1, x2, y2, xU)
		fxr := nonLinearEquations.ObjFunc(x1, y1, x2, y2, xr)
		if math.Abs(fxr) < err1 {
			break
		}

		if (fxu * fxr) < 0 {
			xL = xr
		} else {
			xU = xr
		}
	}
	return xr
}
