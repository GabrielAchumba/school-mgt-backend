package mathematicslibrary

type Funct struct {
	Variables []float64
	Cc        float64
	IsExtrMax bool
}

func NewFunct(_variables []float64, _cc float64, _isExtrMax bool) Funct {
	return Funct{
		Variables: _variables,
		Cc:        _cc,
		IsExtrMax: _isExtrMax,
	}
}
