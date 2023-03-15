package cfdporomedia

import (
	"strconv"

	DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"
	NumMthds "github.com/GabrielAchumba/school-mgt-backend/Simulation/NumericalMethods"
)

/* type CartGridImpl interface {
	Compute(DX NumMthds.MatD,
		DY NumMthds.MatD, DZ NumMthds.MatD, Z float64)
}
*/
type CartGrid struct {
	CartDims           NumMthds.MatD
	IndexMap           NumMthds.MatD
	Volumes            NumMthds.MatD
	Nx                 int
	Ny                 int
	Nz                 int
	NGrid              int
	NGrid2             int
	DX                 NumMthds.MatD
	DY                 NumMthds.MatD
	DZ                 NumMthds.MatD
	Z                  NumMthds.MatD
	Blocks             map[string]DataStructure.BlockData
	BlockIndicesTexts  []string
	BlockIndicesTexts2 []string
	IList              []int
	JList              []int
	KList              []int
}

func NewCatGrid(_nx int, _ny int, _nz int) CartGrid {

	if _nx <= 0 {
		_nx = 1
	}
	if _ny <= 0 {
		_ny = 1
	}
	if _nz <= 0 {
		_nz = 1
	}

	return CartGrid{
		Nx:                 _nx,
		Ny:                 _ny,
		Nz:                 _nz,
		NGrid:              _nx * _ny * _nz,
		Blocks:             DataStructure.Blocks_OrderBy_XYZ(_nx, _ny, _nz),
		BlockIndicesTexts2: make([]string, 0),
	}
}

func (impl *CartGrid) Compute(DX NumMthds.MatD,
	DY NumMthds.MatD, DZ NumMthds.MatD, Z float64) {

	counter := -1
	i := 0
	j := 0
	k := 0
	for k = 0; k < impl.Nz; k++ {
		for j = 0; j < impl.Ny; j++ {
			for i = 0; i < impl.Nx; i++ {
				counter++
				index := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k)
				item := impl.Blocks[index]
				item.Geometry.Dx = DX.Mat[i]
				item.Geometry.Dy = DY.Mat[j]
				item.Geometry.Dz = DZ.Mat[k]
				item.Geometry.Z = Z + (item.Geometry.Dz / 2.0)
				item.Geometry.Area_x = item.Geometry.Dy * item.Geometry.Dz
				item.Geometry.Area_y = item.Geometry.Dx * item.Geometry.Dz
				item.Geometry.Area_z = item.Geometry.Dx * item.Geometry.Dy
				item.Geometry.Vb = item.Geometry.Dx * item.Geometry.Dy * item.Geometry.Dz
				impl.Blocks[index] = item
			}
		}
	}
}
