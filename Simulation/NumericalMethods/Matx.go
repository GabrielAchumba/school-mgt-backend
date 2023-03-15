package numericalmethods

type MatD struct {
	Mat []float64
	r   int
}

func NewMatD() MatD {
	return MatD{
		Mat: make([]float64, 0),
		r:   5,
	}
}

func (impl *MatD) ArithmeticProgression(a float64, b float64, n int) {

	impl.Mat = make([]float64, 0)

	if n <= 1 {
		impl.Mat = append(impl.Mat, a)
	}

	commonDifference := (b - a) / float64(n)
	for i := 0; i < n; i++ {
		xi := a + float64(i)*commonDifference
		impl.Mat = append(impl.Mat, xi)
	}
}

func (impl *MatD) EqualSegments(a float64, b float64, n int) {

	if n <= 1 {
		impl.Mat = append(impl.Mat, b-a)
	}

	commonDifference := (b - a) / float64(n)
	for i := 0; i < n; i++ {
		impl.Mat = append(impl.Mat, commonDifference)
	}

}

func (impl *MatD) Duplicate(a float64, b float64, n int) {

	impl.Mat = make([]float64, 0)

	if n <= 1 {
		impl.Mat = append(impl.Mat, b-a)
	}

	commonDifference := (b - a)
	for i := 0; i < n; i++ {
		impl.Mat = append(impl.Mat, commonDifference)
	}
}

/* type Mat2D struct {
	Mat []MatD
}

func NewMat2D() Mat2D {
	return Mat2D{
		Mat: make([]MatD, 0),
	}
}

type Mat3D struct {
	Mat []Mat2D
}

func NewMat3D() Mat3D {
	return Mat3D{
		Mat: make([]Mat2D, 0),
	}
}
*/
