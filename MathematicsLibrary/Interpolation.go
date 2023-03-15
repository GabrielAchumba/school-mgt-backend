package mathematicslibrary

type Interpolation struct {
}

func NewInterpolation() (interpolation *Interpolation) {
	interpolation = new(Interpolation)
	return
}

func (interpolation *Interpolation) LinearInterpolationScalar(X1 float64, X2 float64, Y1 float64, Y2 float64, X float64) float64 {
	var denom = X2 - X1
	if denom == 0 {
		return Y2
	}

	m := (Y2 - Y1) / denom
	c := Y1
	x := X - X1

	Y := m*x + c

	return Y
}

func (interpolation *Interpolation) LinearInterpolation(Xs []float64, Ys []float64, X float64) float64 {
	var i int = 0
	var ii int = 0
	var check = false
	Y := Ys[0]
	XsCount := len(Xs)

	for i = 1; i < XsCount; i++ {
		if X >= Xs[i-1] && X <= Xs[i] {
			ii = i
			check = true
			break
		}
	}

	if check == true {
		Y = interpolation.LinearInterpolationScalar(Xs[ii-1], Xs[ii], Ys[ii-1], Ys[ii], X)
	}

	return Y
}
