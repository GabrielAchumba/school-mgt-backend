package datastructure

type Residual struct {
	Oil   float64
	Gas   float64
	Water float64
}

func NewResidual() Residual {
	return Residual{}
}

type ResidualDerivatives struct {
	Oil   float64
	Gas   float64
	Water float64
}

func NewResidualDerivatives() ResidualDerivatives {
	return ResidualDerivatives{}
}
