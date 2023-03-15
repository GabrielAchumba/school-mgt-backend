package mathematicslibrary

type Averaging struct {
}

func NewAveraging() Averaging {
	return Averaging{}
}

func (impl *Averaging) Arithmentic(a []float64) float64 {
	var avg float64 = 0
	var sum float64 = 0
	n := len(a)
	for i := 0; i < n; i++ {
		sum = sum + a[i]
	}

	avg = sum / float64(n)
	return avg
}

func (impl *Averaging) Harmonic(a []float64) float64 {
	var avg float64 = 0
	var sum float64 = 0
	n := len(a)
	for i := 0; i < n; i++ {
		var val float64 = 0
		if a[i] != 0 {
			val = 1 / a[i]
		}
		sum = sum + val
	}

	avg = 1 / (sum / float64(n))
	return avg
}
