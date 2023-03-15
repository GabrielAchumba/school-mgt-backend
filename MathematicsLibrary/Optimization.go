package mathematicslibrary

import (
	"fmt"

	"gonum.org/v1/gonum/optimize"
)

type UnconstrainedTest struct {
	// name is the name of the test.
	Name string
	// p is the optimization problem to be solved.
	P optimize.Problem
	// x is the initial guess.
	X []float64
	// gradTol is the absolute gradient tolerance for the test. If gradTol == 0,
	// the default value of 1e-12 will be used.
	GradTol float64
	// fAbsTol is the absolute function convergence for the test. If fAbsTol == 0,
	// the default value of 1e-12 will be used.
	FAbsTol float64
	// fIter is the number of iterations for function convergence. If fIter == 0,
	// the default value of 20 will be used.
	FIter int
	// long indicates that the test takes long time to finish and will be
	// excluded if testing.Short returns true.
	Long bool
}

func (t UnconstrainedTest) String() string {
	dim := len(t.X)
	if dim <= 10 {
		// Print the initial X only for small-dimensional problems.
		return fmt.Sprintf("F: %v\nDim: %v\nInitial X: %v\nGradientThreshold: %v",
			t.Name, dim, t.X, t.GradTol)
	}
	return fmt.Sprintf("F: %v\nDim: %v\nGradientThreshold: %v",
		t.Name, dim, t.GradTol)
}

type Optimization struct {
	UnconstrainedTest  UnconstrainedTest
	OptimizationMethod OptimizationMethod
	Result             *optimize.Result
	Error              error
}

func NewOptimization() (optimization *Optimization) {

	optimization = new(Optimization)

	return
}
func (optimization *Optimization) Run() {

	switch optimization.OptimizationMethod {
	case Newton:
		optimization.Result, optimization.Error = optimization.run(&optimize.Newton{})

	case LBFGS:
		optimization.Result, optimization.Error = optimization.run(&optimize.LBFGS{})

	case BFGS:
		optimization.Result, optimization.Error = optimization.run(&optimize.BFGS{})

	case CG:
		optimization.Result, optimization.Error = optimization.run(&optimize.CG{})

	case GradientDescent:
		optimization.Result, optimization.Error = optimization.run(&optimize.GradientDescent{})

	case NelderMead:
		optimization.Result, optimization.Error = optimization.run(&optimize.NelderMead{})

	}

}

func (optimization *Optimization) Errorf(format string, args ...any) {
	fmt.Sprintf(format, args...)
}

func (optimization *Optimization) defaultFunctionConverge() *optimize.FunctionConverge {
	return &optimize.FunctionConverge{
		Absolute:   1e-10,
		Iterations: 100,
	}
}

func (optimization *Optimization) availFromProblem(prob optimize.Problem) optimize.Available {
	return optimize.Available{Grad: prob.Grad != nil, Hess: prob.Hess != nil}
}

func (optimization *Optimization) run(method optimize.Method) (*optimize.Result, error) {

	test := optimization.UnconstrainedTest
	settings := &optimize.Settings{}
	settings.Converger = optimization.defaultFunctionConverge()
	var uses optimize.Available
	if method != nil {
		var err error
		has := optimization.availFromProblem(test.P)
		uses, err = method.Uses(has)
		if err != nil {
			optimization.Errorf("problem and method mismatch: %v", err)

		}
	}
	if method != nil {
		// Turn off function convergence checks for gradient-based methods.
		if uses.Grad {
			settings.Converger = optimize.NeverTerminate{}
		}
	} else {
		if test.FIter == 0 {
			test.FIter = 20
		}
		c := settings.Converger.(*optimize.FunctionConverge)
		c.Iterations = test.FIter
		if test.FAbsTol == 0 {
			test.FAbsTol = 1e-12
		}
		c.Absolute = test.FAbsTol
		settings.Converger = c
	}
	if test.GradTol == 0 {
		test.GradTol = 1e-12
	}
	settings.GradientThreshold = test.GradTol

	result, err := optimize.Minimize(test.P, test.X, settings, method)
	//result, err := optimize.Minimize(test.P, test.X, nil, nil)
	/* if err != nil {
		Errorf("Case %d: error finding minimum (%v) for:\n%v", err, test)

	}
	if result == nil {
		Errorf("Case %d: nil result without error for:\n%v",  test)

	}

	// Check that the function value at the found optimum location is
	// equal to result.F.
	optF := test.p.Func(result.X)
	if optF != result.F {
		Errorf("Case %d: Function value at the optimum location %v not equal to the returned value %v for:\n%v",
			optF, result.F, test)
	}
	if result.Gradient != nil {
		// Evaluate the norm of the gradient at the found optimum location.
		g := make([]float64, len(test.x))
		test.p.Grad(g, result.X)

		if !floats.Equal(result.Gradient, g) {
			Errorf("Case %d: Gradient at the optimum location not equal to the returned value for:\n%v", test)
		}

		optNorm := floats.Norm(g, math.Inf(1))
		// Check that the norm of the gradient at the found optimum location is
		// smaller than the tolerance.
		if optNorm >= settings.GradientThreshold {
			Errorf("Case %d: Norm of the gradient at the optimum location %v not smaller than tolerance %v for:\n%v",
				optNorm, settings.GradientThreshold, test)
		}
	}

	if !uses.Grad && !uses.Hess {
		// Gradient-free tests can correctly terminate only with
		// FunctionConvergence status.
		if result.Status != optimize.FunctionConvergence {
			Errorf("Status not %v, %v instead", optimize.FunctionConvergence, result.Status)
		}
	}

	// We are going to restart the solution using known initial data, so
	// evaluate them.
	settings.InitValues = &optimize.Location{}
	settings.InitValues.F = test.p.Func(test.x)
	if uses.Grad {
		settings.InitValues.Gradient = resize(settings.InitValues.Gradient, len(test.x))
		test.p.Grad(settings.InitValues.Gradient, test.x)
	}
	if uses.Hess {
		settings.InitValues.Hessian = mat.NewSymDense(len(test.x), nil)
		test.p.Hess(settings.InitValues.Hessian, test.x)
	}

	// Rerun the test again to make sure that it gets the same answer with
	// the same starting condition. Moreover, we are using the initial data.
	result2, err2 := optimize.Minimize(test.p, test.x, settings, method) */
	/* if err2 != nil {
		Errorf("error finding minimum second time (%v) for:\n%v", err2, test)

	}
	if result2 == nil {
		Errorf("second time nil result without error for:\n%v", test)
	}

	// At the moment all the optimizers are deterministic, so check that we
	// get _exactly_ the same answer second time as well.
	if result.F != result2.F || !floats.Equal(result.X, result2.X) {
		Errorf("Different minimum second time for:\n%v", test)
	}

	// Check that providing initial data reduces the number of evaluations exactly by one.
	if result.FuncEvaluations != result2.FuncEvaluations+1 {
		Errorf("Providing initial data does not reduce the number of Func calls for:\n%v", test)
	}
	if uses.Grad {
		if result.GradEvaluations != result2.GradEvaluations+1 {
			Errorf("Providing initial data does not reduce the number of Grad calls for:\n%v", test)
		}
	}
	if uses.Hess {
		if result.HessEvaluations != result2.HessEvaluations+1 {
			Errorf("Providing initial data does not reduce the number of Hess calls for:\n%v", test)
		}
	} */

	return result, err
}
