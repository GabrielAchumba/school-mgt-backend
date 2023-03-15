package mathematicslibrary

import "log"

type Constraint struct {
	Variables []float64
	B         float64
	Sign      string
}

func NewConstraint(_variables []float64, _b float64, _sign string) Constraint {

	if _sign != "=" && _sign != "<=" && _sign != ">=" {
		log.Panic("...Wrong sign...")
	}
	return Constraint{
		Variables: _variables,
		B:         _b,
		Sign:      _sign,
	}
}
