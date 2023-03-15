package mathematicslibrary

type Differentiation struct {
}

func NewDifferentiation() (differentiation *Differentiation) {
	differentiation = new(Differentiation)
	return
}

func (differentiation *Differentiation) Central(fun ObjectiveFunc1, x float64, stepsize float64) float64 {
	var Av, x1, x2 float64
	x1 = x + stepsize
	x2 = x - stepsize
	Av = (fun(x1) - fun(x2)) / (2 * stepsize)
	return Av
}
