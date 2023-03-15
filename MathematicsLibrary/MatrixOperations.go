package mathematicslibrary

import (
	math "math"
)

type MatrixOperations struct {
}

func (matrixOperations MatrixOperations) CreateMatrix(rows int, cols int) [][]float64 {
	var ans [][]float64

	for i := 0; i < rows; i++ {

		var Row []float64
		for j := 0; j < cols; j++ {
			Row = append(Row, 0)
		}

		ans = append(ans, Row)
	}

	return ans
}

func (matrixOperations MatrixOperations) CreateVector(rows int) []float64 {
	var ans []float64

	for i := 0; i < rows; i++ {
		ans = append(ans, 0)
	}

	return ans
}

func (matrixOperations MatrixOperations) VecCum(vec []float64) []float64 {

	var SumVector []float64
	nrow := len(vec)
	SumVector = matrixOperations.CreateVector(nrow)
	SumVector[0] = vec[0]
	for i := 1; i < nrow; i++ {
		SumVector[i] = SumVector[i-1] + vec[i]
	}

	return SumVector
}

func (matrixOperations MatrixOperations) MatTranspose(Matrix [][]float64) [][]float64 {

	var MatrixT [][]float64
	nrow := len(Matrix)
	ncol := len(Matrix[0])

	MatrixT = matrixOperations.CreateMatrix(ncol, nrow)

	for j := 0; j < ncol; j++ {
		for i := 0; i < nrow; i++ {
			MatrixT[j][i] = Matrix[i][j]
		}
	}

	return MatrixT
}

func (matrixOperations MatrixOperations) Unitmatrix(Matrix [][]float64) [][]float64 {

	nrow := len(Matrix)
	ncol := len(Matrix[0])
	IMatrix := matrixOperations.CreateMatrix(nrow, ncol)
	for j := 0; j < ncol; j++ {
		for i := 0; i < nrow; i++ {
			if i == j {
				IMatrix[i][j] = 1
			} else {
				IMatrix[i][j] = 0
			}
		}
	}

	return IMatrix

}

func (matrixOperations MatrixOperations) ScalarMatrixMultiplication(scalar float64,
	Matrix [][]float64) [][]float64 {

	nrow := len(Matrix)
	ncol := len(Matrix[0])
	result := matrixOperations.CreateMatrix(nrow, ncol)
	for j := 0; j < ncol; j++ {
		for i := 0; i < nrow; i++ {
			result[i][j] = scalar * Matrix[i][j]
		}
	}

	return result
}

func (matrixOperations MatrixOperations) ConvertDiaMatrixtoVector(a [][]float64) []float64 {
	nrow := len(a)
	x := matrixOperations.CreateVector(nrow)
	for i := 0; i < nrow; i++ {
		x[i] = a[i][i]
	}

	return x
}

func (matrixOperations MatrixOperations) SubMatrix(MatrixA [][]float64,
	StartRow int, StartCol int, EndRow int, EndCol int) [][]float64 {

	nrowA := EndRow - StartRow + 1
	ncolB := EndCol - StartCol + 1
	CloneMatrix := matrixOperations.CreateMatrix(nrowA, ncolB)
	for j := StartCol; j <= EndCol; j++ {
		for i := StartRow; i <= EndRow; i++ {
			CloneMatrix[i-StartRow][j-StartCol] = MatrixA[i-1][j-1]
		}
	}

	return CloneMatrix
}

func (matrixOperations MatrixOperations) MatSubstract(MatrixA [][]float64,
	MatrixB [][]float64) [][]float64 {

	nrow := len(MatrixA)
	ncol := len(MatrixA[0])
	SubstractionMat := matrixOperations.CreateMatrix(nrow, ncol)
	for j := 0; j < ncol; j++ {
		for i := 0; i < nrow; i++ {
			SubstractionMat[i][j] = MatrixA[i][j] - MatrixB[i][j]
		}
	}

	return SubstractionMat
}

func (matrixOperations MatrixOperations) MatAdd(MatrixA [][]float64,
	MatrixB [][]float64) [][]float64 {

	nrow := len(MatrixA)
	ncol := len(MatrixA[0])
	AddMat := matrixOperations.CreateMatrix(nrow, ncol)
	for j := 0; j < ncol; j++ {
		for i := 0; i < nrow; i++ {
			AddMat[i][j] = MatrixA[i][j] + MatrixB[i][j]
		}
	}

	return AddMat
}

func (matrixOperations MatrixOperations) MatMult(MatrixA [][]float64,
	MatrixB [][]float64) [][]float64 {

	var sum float64 = 0
	nrowA := len(MatrixA)
	ncolB := len(MatrixB[0])
	nrowB := len(MatrixB)
	MultMat := matrixOperations.CreateMatrix(nrowA, ncolB)
	for j := 0; j < nrowA; j++ {
		for i := 0; i < ncolB; i++ {
			sum = 0
			for k := 0; k < nrowB; k++ {
				sum = sum + MatrixA[j][k]*MatrixB[k][i]
			}
			MultMat[j][i] = sum
		}
	}
	return MultMat
}

func (matrixOperations MatrixOperations) Vector2Matrix(vec []float64) [][]float64 {
	nrow := len(vec)
	ncol := 1
	Matrix := matrixOperations.CreateMatrix(nrow, ncol)
	for i := 0; i < nrow; i++ {
		Matrix[i][0] = vec[i]
	}

	return Matrix
}

func (matrixOperations MatrixOperations) Mat2Vec(Matrix [][]float64) []float64 {
	nrow := len(Matrix)
	vec := matrixOperations.CreateVector(nrow)
	for i := 0; i < nrow; i++ {
		vec[i] = Matrix[i][0]
	}

	return vec
}

func (matrixOperations MatrixOperations) CopyVector(vec []float64) []float64 {
	nrow := len(vec)
	vec2 := matrixOperations.CreateVector(nrow)
	for i := 0; i < nrow; i++ {
		vec2[i] = vec[i]
	}

	return vec2
}

func (matrixOperations MatrixOperations) CopyMatrix(XX1 [][]float64) [][]float64 {
	nrow := len(XX1)
	ncol := len(XX1[0])
	xx := matrixOperations.CreateMatrix(nrow, ncol)
	for j := 0; j < ncol; j++ {
		for i := 0; i < nrow; i++ {
			xx[i][j] = XX1[i][j]
		}
	}

	return xx
}

func (matrixOperations MatrixOperations) CalcualatePolynomial(PolyCoeffs []float64, x []float64,
	PolynomialOrder int) []float64 {

	nrow := len(x)
	var sum1 float64 = 0
	ncol := PolynomialOrder
	vec := matrixOperations.CreateVector(nrow)
	for i := 0; i < nrow; i++ {
		sum1 = 0
		for j := 0; j < ncol; j++ {
			sum1 = sum1 + PolyCoeffs[j]*math.Pow(x[i], float64(ncol-j))
		}
		vec[i] = sum1
	}

	return vec
}

func (matrixOperations MatrixOperations) Horzcat(xx [][]float64,
	XX1 [][]float64) [][]float64 {

	nrow := len(xx)
	ncol1 := len(xx[0])
	ncol2 := len(XX1[0])
	ncol := ncol1 + ncol2
	XX2 := matrixOperations.CreateMatrix(nrow, ncol)
	for j := 0; j < ncol1; j++ {
		for i := 0; i < nrow; i++ {
			XX2[i][j] = xx[i][j]
		}
	}

	for j := 0; j < ncol2; j++ {
		for i := 0; i < nrow; i++ {
			XX2[i][j+ncol-1] = XX1[i][j]
		}
	}

	return XX2
}

func (matrixOperations MatrixOperations) CreatMatrixFromPolyCoeffs(PolyCoeffs []float64,
	PolynomialOrder int) [][]float64 {

	nrow := len(PolyCoeffs)
	ncol := PolynomialOrder
	Matrix := matrixOperations.CreateMatrix(nrow, ncol)
	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			Matrix[i][j] = math.Pow(PolyCoeffs[i], float64(ncol-j))
		}
	}

	return Matrix

}

func (matrixOperations MatrixOperations) ZeroMatrix(MatrixA [][]float64) [][]float64 {
	nrow := len(MatrixA)
	ZeroMat := matrixOperations.CreateMatrix(nrow, nrow)
	return ZeroMat
}

func (matrixOperations MatrixOperations) ZeroVector(vec []float64) []float64 {
	nrow := len(vec)
	return matrixOperations.CreateVector(nrow)
}

func (matrixOperations MatrixOperations) OveWriteMatrix(MatrixA [][]float64,
	StartRow int, StartCol int, EndRow int, EndCol int,
	ReplaceWith [][]float64) [][]float64 {

	MatrixA1 := matrixOperations.CopyMatrix(MatrixA)
	for j := StartCol; j <= EndCol; j++ {
		for i := StartRow; i <= EndRow; i++ {
			MatrixA1[i-1][j-1] = ReplaceWith[i-StartRow][j-StartCol]
		}
	}
	return MatrixA1
}

func (matrixOperations MatrixOperations) norm(Matrix [][]float64) float64 {
	var num float64 = 0
	var sum float64 = 0
	nRows := len(Matrix)
	nCols := len(Matrix[0])
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			sum = sum + math.Pow(Matrix[i][j], 2)
		}
	}

	num = math.Pow(sum, 0.5)
	return num
}

func (matrixOperations MatrixOperations) Vectornorm(vec []float64) float64 {
	var num float64 = 0
	var sum float64 = 0
	nRows := len(vec)
	for i := 0; i < nRows; i++ {
		sum = sum + math.Pow(vec[i], 2)
	}

	num = math.Pow(sum, 0.5)
	return num
}

func (matrixOperations MatrixOperations) DiagonalMatrix(MatrixA [][]float64) [][]float64 {
	nRows := len(MatrixA)
	nCols := len(MatrixA[0])
	DiagMatrixA := matrixOperations.CreateMatrix(nRows, nCols)
	for i := 0; i < nRows; i++ {
		DiagMatrixA[i][i] = MatrixA[i][i]
	}
	return DiagMatrixA
}

func (matrixOperations MatrixOperations) VectorMinimun(a []float64) float64 {
	nRows := len(a)
	xmin := a[0]
	for i := 0; i < nRows; i++ {
		if a[i] < xmin {
			xmin = a[i]
		}
	}

	return xmin
}

func (matrixOperations MatrixOperations) VectorMinIndex(a []float64) int {
	nRows := len(a)
	var ii int = 0
	xmin := a[0]
	for i := 0; i < nRows; i++ {
		if a[i] < xmin {
			xmin = a[i]
			ii = i
		}
	}

	return ii
}

func (matrixOperations MatrixOperations) VectorMaximum(a []float64) float64 {
	nRows := len(a)
	xmax := a[0]
	for i := 0; i < nRows; i++ {
		if a[i] > xmax {
			xmax = a[i]
		}
	}

	return xmax
}

func (matrixOperations MatrixOperations) MatrixMaximum(a [][]float64) float64 {
	xmax := a[0][0]
	nRows := len(a)
	nCols := len(a[0])
	for j := 0; j < nCols; j++ {
		for i := 0; i < nRows; i++ {
			if a[i][j] > xmax {
				xmax = a[i][j]
			}
		}
	}

	return xmax
}

/* double MatrixOperations::MatrixMinimum(vector<vector<double>> a){
    int i = 0, j = 0;
    double xmin = a[0][0];
    int nRows = a.size();
    int nCols = a[0].size();
    for(j = 0; j < nCols; j++){
        for(i = 0; i < nRows; i++){
            if(a[i][j] < xmin){
                xmin = a[i][j];
            }
        }
    }

    return xmin;
}

vector<vector<double>> MatrixOperations::VecMatAdd(vector<vector<double>>& Matrix,
       vector<double>& vec){
    int i, j, ncol, nrow;
    nrow = Matrix.size();
    ncol = Matrix[0].size();
    vector<vector<double>> AddMat = createMatrix(nrow, ncol);
    for(j = 0; j < ncol; j++){
        for(i = 0; i < nrow; i++){
            AddMat[i][j] = Matrix[i][j] + vec[i];
        }
    }
    return AddMat;
}

vector<vector<double>> MatrixOperations::VecMatSubtract(vector<vector<double>>& Matrix,
       vector<double>& vec){
    int i, j, ncol, nrow;
    nrow = Matrix.size();
    ncol = Matrix[0].size();
    vector<vector<double>> AddMat = createMatrix(nrow, ncol);
    for(j = 0; j < ncol; j++){
        for(i = 0; i < nrow; i++){
            AddMat[i][j] = vec[i] - Matrix[i][j];
        }
    }
    return AddMat;
}

vector<vector<double>> MatrixOperations::MatVecSubtract(vector<vector<double>>& Matrix,
       vector<double>& vec){
    int i, j, ncol, nrow;
    nrow = Matrix.size();
    ncol = Matrix[0].size();
    vector<vector<double>> AddMat = createMatrix(nrow, ncol);
    for(j = 0; j < ncol; j++){
        for(i = 0; i < nrow; i++){
            AddMat[i][j] = Matrix[i][j] - vec[i];
        }
    }
    return AddMat;
}

double MatrixOperations::SumofSquares(vector<vector<double>>& Matrix){
   int i = 0;
   double  sum = 0;
   int nRows = Matrix.size();
   for(i = 0; i < nRows; i++){
       sum = sum + pow(Matrix[i][0],2);
   }

   return sum;
}

vector<vector<double>> MatrixOperations::BubbleSortMatrix( vector<vector<double>>& Matrix, int& sortcolumn){
    int icol, inner, outer, ncol;
    double Temp;
    ncol = Matrix[0].size();
    vector<vector<double>> Matrix2 = CopyMatrix(Matrix);
    for(outer = ncol; outer >= 0; outer--){
        for(inner  = 0; inner < outer-1; inner++){
            if( Matrix2[inner][sortcolumn] > Matrix2[inner + 1][sortcolumn]){
              for(icol = 0; icol < ncol; icol++){
                  Temp = Matrix2[inner][icol];
                Matrix2[inner][icol] = Matrix2[inner + 1][icol];
                Matrix2[inner + 1][icol] = Temp;
              }
            }
        }
    }
    return Matrix2;
}

vector<double> MatrixOperations::BubbleSortVector(vector<double>& vec){
    int inner = 0, outer = 0;
    vector<double> vec2 = CopyVector(vec);
    int nRows = vec2.size();
    double Temp = 0;
    for(outer = nRows; outer >= 0; outer--){
        for(inner = 0; inner < outer-1; inner++){
            if(vec2[inner] > vec2[inner + 1]){
                Temp = vec2[inner];
                vec2[inner] = vec2[inner + 1];
                vec2[inner + 1] = Temp;
            }
        }
    }

    return vec2;
}

vector<double> MatrixOperations::Get_s(vector<vector<double>>& a){
    int i = 0, j = 0;
    int n = a.size();
    vector<double> s = createVector(n);
    for(i = 0; i < n; i++){
        s[i] = abs(a[i][1]);
        for(j = 0; j < n; j++){
            if (abs(a[i][j]) > s[i]){
                 s[i] = abs(a[i][j]);
            }
        }
    }

    return s;
}

LinSysResult MatrixOperations::LowerUpperTriangularMatrix(vector<vector<double>>& A)
{
    vector<vector<double>> UT = CopyMatrix(A);
    vector<vector<double>> TempRow1;
    vector<vector<double>> TempRow2;
    vector<vector<double>> Row1, Row2;
    int N = UT.size(), nCols = UT[0].size(), k, i;
    double multiplier = 0;
    vector<vector<double>> LT = unitmatrix(UT);

    for (k = 0; k < N - 1; k++)
    {

        Row1 = SubMatrix(UT, k+1, 1, k+1, nCols);// UT[k + 1, ":"];
        for (i = k + 1; i < N; i++)
        {
            Row2 = SubMatrix(UT, i+1, 1, i+1, nCols); // UT[i + 1, ":"];
            multiplier = Row2[0][k] / Row1[0][k];
            TempRow1 = ScalarMatrixMultiplication(multiplier, Row1);
            TempRow2 = MatSubstract(Row2, TempRow1);
            //vector<vector<double>> BB;
            UT = OveWriteMatrix(UT, i+1, 1, i+1, nCols, TempRow2);  //UT[i + 1, ":"] = TempRow2;
            LT[i][k] = multiplier;

        }
    }
    LinSysResult linSysResult;
    linSysResult.LowerTriangularMatrix = LT;
    linSysResult.UpperTriangularMatrix = UT;
    return linSysResult;
}

vector<vector<double>> MatrixOperations::ForwardSubstitution(vector<vector<double>>& LT, vector<vector<double>>& ColMat)
{
    int i, j, N; N = LT.size();
    vector<vector<double>> Z = createMatrix(N, 1); double sum = 0;
    Z[0][0] = ColMat[0][0] / LT[0][0];
    for (i = 1; i < N; i++)
    {
        sum = 0;
        for (j = 0; j <= (i - 1); j++)
        {
            sum = sum + LT[i][j] * Z[j][0];
        }
        Z[i][0] = (ColMat[i][0] - sum) / LT[i][i];
    }

    return Z;
}

  vector<vector<double>> MatrixOperations::BackSubstitution(vector<vector<double>>& HCat)
        {
            int nrows = HCat.size();
            int ncols = HCat[0].size();
            vector<vector<double>> A = SubMatrix(HCat, 1, 1, nrows, ncols - 1);
            vector<vector<double>> B = SubMatrix(HCat, 1, ncols, nrows, ncols);
            int N = A.size(), NN = N - 1;
            vector<vector<double>> x = createMatrix(N, 1); double summ = 0;

            if (A[NN][NN] == 0)
            {
                x[NN][0] = 0;
            }
            else
            {
                x[NN][0] = B[NN][0] / A[NN][NN];
            }

            for (int i = NN - 1; i >= 0; i--)
            {
                summ = 0;
                for (int j = i + 1; j <= NN; j++)
                {
                    summ = summ + A[i][j] * x[j][0];
                }
                x[i][0] = (B[i][0] - summ) / A[i][i];
            }

            return x;
        }

vector<vector<double>> MatrixOperations::LU_Decomposition(vector<vector<double>>& LHS3, vector<vector<double>>& RHS3){

    LinSysResult LU = LowerUpperTriangularMatrix(LHS3);
    vector<vector<double>> LT = LU.LowerTriangularMatrix;
    vector<vector<double>> UT = LU.UpperTriangularMatrix;
    vector<vector<double>>  Z = ForwardSubstitution(LT, RHS3);
    vector<vector<double>> HC = horzcat(UT, Z);
    vector<vector<double>> X = BackSubstitution(HC);
    return X;
} */
