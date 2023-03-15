package cfdporomedia

import (
	"strconv"

	NumMthds "github.com/GabrielAchumba/school-mgt-backend/Simulation/NumericalMethods"
)

type DataType int

const (
	Scalar DataType = iota + 1
	Vector
)

type Rock struct {
	Porosity               float64
	PermeabilityX          float64
	PermeabilityY          float64
	PermeabilityZ          float64
	Compressibility        float64
	PorosityDataType       DataType
	PermeabilityXDataType  DataType
	PermeabilityYDataType  DataType
	PermeabilityZDataType  DataType
	RockCompDataType       DataType
	PermeabilityX_Vector   []float64
	PermeabilityY_Vector   []float64
	PermeabilityZ_Vector   []float64
	Porosity_Vector        []float64
	Compressibility_Vector []float64
}

type RockData struct {
	CartGrid CartGrid
}

func NewRockData(cG CartGrid) RockData {
	return RockData{
		CartGrid: cG,
	}
}

func (impl *RockData) SetRockData(_PORO NumMthds.MatD, _PERMX NumMthds.MatD,
	_PERMY NumMthds.MatD, _PERMZ NumMthds.MatD, CR NumMthds.MatD) {

	counter := -1
	zeroApprox := 0.000000001
	i := 0
	j := 0
	k := 0
	for k = 0; k < impl.CartGrid.Nz; k++ {
		for j = 0; j < impl.CartGrid.Ny; j++ {
			for i = 0; i < impl.CartGrid.Nx; i++ {
				counter++
				index := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k)
				porosity := _PORO.Mat[counter]
				permX := _PERMX.Mat[counter]
				permY := _PERMY.Mat[counter]
				permZ := _PERMZ.Mat[counter]
				rockComp := CR.Mat[counter]

				if porosity <= 0 {
					porosity = zeroApprox
				}
				if permX <= 0 {
					permX = zeroApprox
				}
				if permY <= 0 {
					permY = zeroApprox
				}
				if permZ <= 0 {
					permZ = zeroApprox
				}
				if rockComp <= 0 {
					rockComp = zeroApprox
				}

				item := impl.CartGrid.Blocks[index]
				item.RockData.Porosity = porosity
				item.RockData.Kx = permX
				item.RockData.Ky = permY
				item.RockData.Kz = permZ
				item.RockData.RockCompressibility = rockComp
				impl.CartGrid.Blocks[index] = item
			}
		}
	}
}
