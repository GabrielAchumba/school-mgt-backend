package mathematicslibrary

import (
	"math"
)

type Interception struct {
	interpolation      *Interpolation
	nonLinearEquations *NonLinearEquations
}

func NewInterception() (interception *Interception) {
	interception = new(Interception)
	interception.interpolation = NewInterpolation()
	interception.nonLinearEquations = NewNonLinearEquations()
	return
}

func (interception *Interception) Intercept(
	Xs1 []float64, Ys1 []float64,
	Xs2 []float64, Ys2 []float64, x0 float64) (float64, float64) {

	fp := func(x float64) float64 {
		yline1 := interception.interpolation.LinearInterpolation(Xs1, Ys1, x)
		yline2 := interception.interpolation.LinearInterpolation(Xs2, Ys2, x)
		diff := math.Abs(yline1 - yline2)
		return diff
	}

	xAns := interception.nonLinearEquations.NewtonRaphson(fp, x0)
	yAns := interception.interpolation.LinearInterpolation(Xs1, Ys1, xAns)

	return xAns, yAns
}
