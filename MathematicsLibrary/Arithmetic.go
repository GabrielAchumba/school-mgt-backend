package mathematicslibrary

type Arithmetic struct {
}

func (arithmetic Arithmetic) min(values []float64) float64 {
	var x float64 = 0
	var length = len(values)
	x = values[0]
	for i := 0; i < length; i++ {
		if values[i] < x {
			x = values[i]
		}
	}

	return x
}

func (arithmetic Arithmetic) max(values []float64) float64 {
	var x float64 = 0
	var length = len(values)
	x = values[0]
	for i := 0; i < length; i++ {
		if values[i] > x {
			x = values[i]
		}
	}

	return x
}
