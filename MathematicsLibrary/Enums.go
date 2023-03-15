package mathematicslibrary

type SimplexResult int

const (
	Unbounded   SimplexResult = iota + 1 // EnumIndex = 1
	Found                                // EnumIndex = 2
	NotYetFound                          // EnumIndex = 3
)

type OptimizationMethod int

const (
	Newton OptimizationMethod = iota + 1 // EnumIndex = 1
	LBFGS                                // EnumIndex = 2
	BFGS                                 // EnumIndex = 3
	CG
	GradientDescent
	NelderMead
)
