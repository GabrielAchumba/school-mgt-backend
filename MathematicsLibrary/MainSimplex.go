package mathematicslibrary

import (
	"fmt"
	"strconv"
)

func CreateVectorOfT(rows int) []float64 {
	ans := make([]float64, 0)
	for i := 0; i < rows; i++ {
		ans = append(ans, 0)
	}

	return ans
}

func CreateVectorOfStrings(rows int) []string {
	ans := make([]string, 0)
	for i := 0; i < rows; i++ {
		ans = append(ans, "")
	}

	return ans
}

func CreateMatrixOfStrings(rows int, cols int) [][]string {
	ans := make([][]string, 0)
	for j := 0; j < cols; j++ {
		Row := make([]string, 0)
		for i := 0; i < rows; i++ {
			Row = append(Row, "")
		}

		ans = append(ans, Row)
	}

	return ans
}

func FillConstraintsGrid(variablesCount int,
	constraintsCount int) [][]string {

	constraintsGridView := make([][]string, 0)

	variablesCountPlusTwo := variablesCount + 2
	for i := 0; i < constraintsCount; i++ {

		row := CreateVectorOfStrings(variablesCountPlusTwo)
		constraintsGridView = append(constraintsGridView, row)
	}
	return constraintsGridView
}

func FillFunctionGrid(variablesCount int, _constraintsGridView [][]string) ([]string, [][]string) {

	constraintsGridView := _constraintsGridView
	variablesCountPlusOne := variablesCount + 1
	functionGridView := CreateVectorOfStrings(variablesCountPlusOne)

	variablesCountPlusTwo := variablesCount + 2
	row := CreateVectorOfStrings(variablesCountPlusTwo)
	constraintsGridView = append(constraintsGridView, row)

	return functionGridView, constraintsGridView
}

func ToDouble(x string) float64 {
	ans, _ := strconv.ParseFloat(x, 64)
	return ans
}

type MainSimplex struct {
	ConstraintsCount       int
	VariablesCount         int
	VariableNames          []string
	DecisonVariables       []float64
	ConstraintsGridView    [][]string
	FunctionGridView       []string
	ResultsGridView        [][]string
	Unbounded              SimplexResult
	Found                  SimplexResult
	NotYetFound            SimplexResult
	OptimalSolution        float64
	OptimalSolutionMessage string
}

func NewMainSimplex(_constraintsCount int, _variablesCount int,
	_variableNames []string) MainSimplex {

	mainSimplex := MainSimplex{}
	mainSimplex.ConstraintsCount = _constraintsCount
	mainSimplex.VariablesCount = _variablesCount
	mainSimplex.VariableNames = _variableNames
	mainSimplex.Unbounded = Unbounded
	mainSimplex.Found = Found
	mainSimplex.NotYetFound = NotYetFound
	mainSimplex.OptimalSolution = 0
	mainSimplex.OptimalSolutionMessage = ""
	variableNamessize := len(mainSimplex.VariableNames)
	mainSimplex.DecisonVariables = CreateVectorOfT(variableNamessize)
	mainSimplex.ConstraintsGridView = FillConstraintsGrid(mainSimplex.VariablesCount,
		mainSimplex.ConstraintsCount)
	mainSimplex.FunctionGridView, mainSimplex.ConstraintsGridView =
		FillFunctionGrid(mainSimplex.VariablesCount,
			mainSimplex.ConstraintsGridView)

	return mainSimplex
}

func (impl MainSimplex) GetSolution(snap SimplexSnap) {

	impl.ResultsGridView = make([][]string, 0)

	snapCCsize := len(snap.CC)
	snapmatrixsizePlus4 := len(snap.Matrix) + 4
	for i := 0; i < snapCCsize; i++ {
		row := CreateVectorOfStrings(snapmatrixsizePlus4)

		for j := 0; j < snapmatrixsizePlus4; j++ {
			if j == 0 {
				row[j] = strconv.Itoa(i + 1)
			} else if j == 1 {
				row[j] = "A" + strconv.Itoa(snap.CC[i]+1)

			} else if j == 2 {
				row[j] = fmt.Sprintf("%f", snap.FVars[snap.CC[i]])
				if snap.MM[snap.CC[i]] {
					row[j] = "-M"
				}
			} else if j == 3 {
				row[j] = fmt.Sprintf("%f", snap.B[i])

				variableNamessize := len(impl.VariableNames)
				for ii := 0; ii < variableNamessize; ii++ {
					if row[j-2] == impl.VariableNames[ii] {
						impl.DecisonVariables[ii] = snap.B[i]
					}
				}
			} else {
				row[j] = fmt.Sprintf("%f", snap.Matrix[j-4][i])
			}
		}
	}
}

func (impl MainSimplex) Proceed() {
	constraints := make([]Constraint, 0)
	for i := 0; i < impl.ConstraintsCount; i++ {
		variables := CreateVectorOfT(impl.VariablesCount)
		b := ToDouble(impl.ConstraintsGridView[i][impl.VariablesCount+1])
		sign := impl.ConstraintsGridView[i][impl.VariablesCount]
		for j := 0; j < impl.VariablesCount; j++ {
			variables[j] = ToDouble(impl.ConstraintsGridView[i][j])
		}

		constraintTemp := NewConstraint(variables, b, sign)

		constraints = append(constraints, constraintTemp)
	}

	functionVariables := CreateVectorOfT(impl.VariablesCount)
	for i := 0; i < impl.VariablesCount; i++ {
		functionVariables[i] = ToDouble(impl.FunctionGridView[i])
	}
	cc := ToDouble(impl.FunctionGridView[impl.VariablesCount])

	isExtrMax := true // extrComboBox.SelectedIndex == 0;

	function := NewFunct(functionVariables, cc, isExtrMax)

	simplex := NewSimplex2(function, constraints)

	snaps, simplexResult := simplex.GetResult()
	lent := len(snaps)

	if simplexResult == impl.Found {
		/* extrStr := "min"
		if isExtrMax {
			extrStr = "max"
		} */

		impl.OptimalSolution = snaps[lent-1].FValue
		impl.OptimalSolutionMessage = "The optimal solution: P = " + fmt.Sprintf("%f", impl.OptimalSolution)
	}

	if simplexResult == impl.Unbounded {
		impl.OptimalSolutionMessage = "The domain of admissible solutions is unbounded"
	}

	if simplexResult == impl.NotYetFound {
		impl.OptimalSolutionMessage = "Algorithm has made 100 cycles and hasn't found any optimal solution."
	}

	snapssizeMinus1 := len(snaps) - 1
	simplexSnap := snaps[snapssizeMinus1]
	impl.GetSolution(simplexSnap)

}
