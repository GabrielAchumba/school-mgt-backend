package mathematicslibrary

import "math"

type Simplex struct {
	Function Funct

	FunctionVariables []float64
	Matrix            [][]float64
	B                 []float64
	MM                []bool
	M                 []float64
	F                 []float64
	CC                []int
	IsMDone           bool
	Unbounded         SimplexResult
	Found             SimplexResult
	NotYetFound       SimplexResult
}

func CreateMatrix(rows int, cols int) [][]float64 {
	ans := make([][]float64, 0)
	for j := 0; j < cols; j++ {
		Row := make([]float64, 0)
		for i := 0; i < rows; i++ {
			Row = append(Row, 0)
		}
		ans = append(ans, Row)
	}

	return ans
}

func CreateVector(rows int) []float64 {
	ans := make([]float64, 0)

	for i := 0; i < rows; i++ {
		ans = append(ans, 0)
	}

	return ans
}

func GetColumn(value float64, place int, length int) []float64 {
	newColumn := make([]float64, 0)

	for k := 0; k < length; k++ {
		var ans float64 = 0
		if k == place {
			ans = value
		}
		newColumn = append(newColumn, ans)
	}

	return newColumn
}

func AppendColumn(matrix [][]float64, column []float64) [][]float64 {
	newMatrix := make([][]float64, 0)
	cols := len(matrix)
	for i := 0; i < cols; i++ {
		Row := matrix[i]
		newMatrix = append(newMatrix, Row)
	}
	newMatrix = append(newMatrix, column)

	return newMatrix
}

func AppendBool(arrayData []bool, element bool) []bool {
	newArray := make([]bool, 0)
	arrayDatasize := len(arrayData)
	for i := 0; i < arrayDatasize; i++ {
		newArray = append(newArray, arrayData[i])
	}
	newArray = append(newArray, element)

	return newArray
}

func Canonize(_function Funct) Funct {
	newFuncVars := make([]float64, 0)
	_functionvariablessize := len(_function.Variables)
	for i := 0; i < _functionvariablessize; i++ {
		newFuncVars = append(newFuncVars, -_function.Variables[i])
	}

	tru := true
	Minus_functioncc := -_function.Cc
	ans := NewFunct(newFuncVars, Minus_functioncc, tru)

	return ans
}

func GetMatrix(constraints []Constraint) ([]float64, [][]float64, []int, []bool) {

	constraintssize := len(constraints)
	constraintsvariablessize := 0
	for i := 0; i < constraintssize; i++ {
		if constraints[i].B < 0 {
			cVars := make([]float64, 0)
			constraintsvariablessize := len(constraints[i].Variables)
			for j := 0; j < constraintsvariablessize; j++ {
				cVars = append(cVars, -constraints[i].Variables[j])
			}

			sign := constraints[i].Sign

			if sign == ">=" {
				sign = "<="
			} else if sign == "<=" {
				sign = ">="
			}

			minusconstraints_i_b := -constraints[i].B
			cNew := NewConstraint(cVars, minusconstraints_i_b, sign)
			constraints[i] = cNew
		}
	}

	constraintsvariablessize = len(constraints[0].Variables)
	constraintssize = len(constraints)
	matrix := CreateMatrix(constraintssize, constraintsvariablessize)

	for i := 0; i < constraintsvariablessize; i++ {
		for j := 0; j < constraintssize; j++ {
			matrix[i][j] = constraints[j].Variables[i]
		}
	}

	appendixMatrix := make([][]float64, 0)
	vecT := make([]float64, 0)
	Bs := CreateVector(constraintssize)
	for i := 0; i < constraintssize; i++ {
		current := constraints[i]

		Bs[i] = current.B

		if current.Sign == ">=" {
			var minusOne float64 = -1
			vecT = GetColumn(minusOne, i, constraintssize)
			appendixMatrix = AppendColumn(appendixMatrix, vecT)
		} else if current.Sign == "<=" {
			var plusOne float64 = 1
			vecT = GetColumn(plusOne, i, constraintssize)
			appendixMatrix = AppendColumn(appendixMatrix, vecT)
		}
	}

	newMatrix := make([][]float64, 0)
	constraintsvariablessize = len(constraints[0].Variables)
	for i := 0; i < constraintsvariablessize; i++ {
		newMatrix = append(newMatrix, matrix[i])
	}

	appendixMatrixsize := len(appendixMatrix)
	for i := constraintsvariablessize; i < constraintsvariablessize+appendixMatrixsize; i++ {
		newMatrix = append(newMatrix, appendixMatrix[i-constraintsvariablessize])
	}

	hasBasicVar := make([]bool, 0)

	constraintssize = len(constraints)
	for i := 0; i < constraintssize; i++ {
		hasBasicVar = append(hasBasicVar, false)
	}

	CC := make([]int, 0)

	newMatrixsize := len(newMatrix)
	for i := 0; i < newMatrixsize; i++ {

		hasOnlyNulls := true
		hasOne := false
		onePosition := make([]int, 2)
		for j := 0; j < constraintssize; j++ {

			if newMatrix[i][j] == 1 {
				if hasOne {
					hasOnlyNulls = false
					break
				} else {
					hasOne = true
					onePosition[0] = i
					onePosition[1] = j
				}
			} else if newMatrix[i][j] != 0 {
				hasOnlyNulls = false
				break
			}

		}

		if hasOnlyNulls && hasOne {
			hasBasicVar[onePosition[1]] = true
			CC = append(CC, onePosition[0])
		}

	}

	MM := make([]bool, 0)

	for i := 0; i < newMatrixsize; i++ {
		MM = append(MM, false)
	}

	for i := 0; i < constraintssize; i++ {

		if !hasBasicVar[i] {

			basicColumn := CreateVector(constraintssize)

			for j := 0; j < constraintssize; j++ {
				basicColumn[j] = 0
				if j == i {
					basicColumn[j] = 1
				}
			}

			newMatrix = AppendColumn(newMatrix, basicColumn)
			tru := true
			MM = AppendBool(MM, tru)
			CC = append(CC, len(newMatrix)-1)
		}

	}

	return Bs, newMatrix, CC, MM
}

func GetFunctionArray(function Funct,
	Matrix [][]float64) []float64 {

	matrixsize := len(Matrix)
	funcVars := CreateVector(matrixsize)
	functionvariablessize := len(function.Variables)
	for i := 0; i < matrixsize; i++ {
		funcVars[i] = 0
		if i < functionvariablessize {
			funcVars[i] = function.Variables[i]
		}
	}

	return funcVars
}

func GetMandF(FunctionVariables []float64,
	CC []int,
	MM []bool,
	Matrix [][]float64) ([]float64, []float64) {

	matrixsize := len(Matrix)
	M := CreateVector(matrixsize)
	F := CreateVector(matrixsize)

	for i := 0; i < matrixsize; i++ {
		var sumF float64 = 0
		var sumM float64 = 0
		matrix_i_size := len(Matrix[0])
		for j := 0; j < matrix_i_size; j++ {
			if MM[CC[j]] {
				sumM -= Matrix[i][j]
			} else {
				sumF += FunctionVariables[CC[j]] * Matrix[i][j]
			}
		}
		M[i] = sumM
		if MM[i] {
			M[i] = sumM + 1
		}
		F[i] = sumF - FunctionVariables[i]
	}

	return M, F
}

func NewSimplex() Simplex {
	return Simplex{
		IsMDone:     false,
		Unbounded:   Unbounded,
		Found:       Found,
		NotYetFound: NotYetFound,
	}
}

func NewSimplex2(_function Funct, _constraints []Constraint) Simplex {
	simplex := Simplex{
		IsMDone:     false,
		Unbounded:   Unbounded,
		Found:       Found,
		NotYetFound: NotYetFound,
	}

	if _function.IsExtrMax {
		simplex.Function = _function
	} else {
		simplex.Function = Canonize(_function)
	}

	Bs, newMatrix, CC, MM := GetMatrix(_constraints)
	simplex.B = Bs
	simplex.Matrix = newMatrix
	simplex.CC = CC
	simplex.MM = MM

	simplex.FunctionVariables = GetFunctionArray(simplex.Function, newMatrix)
	simplex.M, simplex.F = GetMandF(simplex.FunctionVariables,
		simplex.CC, simplex.MM, newMatrix)

	FSize := len(simplex.F)
	for i := 0; i < FSize; i++ {
		simplex.F[i] = -simplex.FunctionVariables[i]
	}

	return simplex
}

func (impl Simplex) GetIndexOfNegativeElementWithMaxAbsoluteValue(arrayData []float64) int {
	index := -1
	arrayDataSize := len(arrayData)
	for i := 0; i < arrayDataSize; i++ {
		if arrayData[i] < 0 {
			if !impl.IsMDone || (impl.IsMDone && !impl.MM[i]) {
				if index == -1 {
					index = i
				} else if math.Abs(arrayData[i]) > math.Abs(arrayData[index]) {
					index = i
				}
			}

		}
	}
	return index
}

func (impl Simplex) GetIndexOfMinimalRatio(column []float64, b []float64) int {
	index := -1
	columnsize := len(column)
	for i := 0; i < columnsize; i++ {
		if column[i] > 0 && b[i] > 0 {
			if index == -1 {
				index = i
			} else if b[i]/column[i] < b[index]/column[index] {
				index = i
			}
		}
	}

	return index
}

func (impl Simplex) CalculateSimplexTableau(Xij []int) {

	J := Xij[1]
	I := Xij[0]
	matrixsize := len(impl.Matrix)
	bsize := len(impl.B)

	impl.CC[J] = I

	newJRow := CreateVector(matrixsize)

	for i := 0; i < matrixsize; i++ {
		newJRow[i] = impl.Matrix[i][J] / impl.Matrix[I][J]
	}

	newB := CreateVector(bsize)

	for i := 0; i < bsize; i++ {
		if i == J {
			newB[i] = impl.B[i] / impl.Matrix[I][J]
		} else {
			newB[i] = impl.B[i] - impl.B[J]/impl.Matrix[I][J]*impl.Matrix[I][i]
		}
	}

	impl.B = newB
	CCsize := len(impl.CC)
	newMatrix := CreateMatrix(CCsize, matrixsize)

	for i := 0; i < matrixsize; i++ {
		for j := 0; j < CCsize; j++ {
			if j == J {
				newMatrix[i][j] = newJRow[i]
			} else {
				newMatrix[i][j] = impl.Matrix[i][j] - newJRow[i]*impl.Matrix[I][j]
			}
		}
	}

	impl.Matrix = newMatrix
	impl.M, impl.F = GetMandF(impl.FunctionVariables,
		impl.CC, impl.MM, newMatrix)
}

func (impl Simplex) NextStep() SimplexIndexResult {

	columnM := impl.GetIndexOfNegativeElementWithMaxAbsoluteValue(impl.M)

	if impl.IsMDone || columnM == -1 {
		//M doesn't have negative values
		impl.IsMDone = true
		columnF := impl.GetIndexOfNegativeElementWithMaxAbsoluteValue(impl.F)

		if columnF != -1 {
			//Has at least 1 negative value
			row := impl.GetIndexOfMinimalRatio(impl.Matrix[columnF], impl.B)

			if row != -1 {
				tupleObject := []int{columnF, row}
				simplexIndexResult := NewSimplexIndexResult(tupleObject, impl.NotYetFound)

				return simplexIndexResult
			} else {
				tupleObject := []int{0, 0}
				simplexIndexResult := NewSimplexIndexResult(tupleObject, impl.Unbounded)
				return simplexIndexResult
			}
		} else {
			tupleObject := []int{0, 0}
			simplexIndexResult := NewSimplexIndexResult(tupleObject, impl.Found)
			return simplexIndexResult
		}

	} else {
		row := impl.GetIndexOfMinimalRatio(impl.Matrix[columnM], impl.B)

		if row != -1 {
			tupleObject := []int{columnM, row}
			simplexIndexResult := NewSimplexIndexResult(tupleObject, impl.NotYetFound)
			return simplexIndexResult
		} else {
			tupleObject := []int{0, 0}
			simplexIndexResult := NewSimplexIndexResult(tupleObject, impl.Unbounded)
			return simplexIndexResult
		}
	}
}

func (impl Simplex) GetResult() ([]SimplexSnap, SimplexResult) {
	snaps := make([]SimplexSnap, 0)
	simplexSnap := NewSimplexSnap2(impl.B, impl.Matrix, impl.M, impl.F,
		impl.CC, impl.FunctionVariables, impl.IsMDone, impl.MM)

	snaps = append(snaps, simplexSnap)

	result := impl.NextStep()
	i := 0
	for result.Result == impl.NotYetFound && i < 100 {
		impl.CalculateSimplexTableau(result.Index)
		simplexSnap2 := NewSimplexSnap2(impl.B, impl.Matrix, impl.M, impl.F,
			impl.CC, impl.FunctionVariables, impl.IsMDone, impl.MM)
		//snaps.push_back(simplexSnap2);
		snaps = append(snaps, simplexSnap2)
		result = impl.NextStep()
		i++
	}

	return snaps, result.Result
}
