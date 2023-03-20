package datastructure

import (
	"strconv"
)

type BlockData struct {
	Geometry               Geometry
	RockData               RockData
	FluidData              FluidData
	SaturationData         SaturationData
	TransmisibilityData    TransmisibilityData
	Acumulation            Acumulation
	Wells                  []WellData
	PressureData           PressureData
	Residual               Residual
	ResidualDerivatives    ResidualDerivatives
	AdjResidualDerivatives map[int]ResidualDerivatives
	Flux                   Flux
	IJK                    int
	IMinusOneJK            int
	IPlusOneJK             int
	IJMinusOneK            int
	IJPlusOneK             int
	IJKMinusOne            int
	IJKPlusOne             int
	Qosc                   float64
	Qgsc                   float64
	Qwsc                   float64
	Pwf                    float64
	Zigma                  float64
	Wells_LHS              float64
	Wells_RHS              float64
	Boundary               Boundary
	VolumeInPlace          VolumeInPlace
	ProductivityIndex      float64
	I                      int
	J                      int
	K                      int
	KeyIndex               string
}

func NewBlockData() BlockData {
	return BlockData{
		Geometry:               NewGeometry(),
		RockData:               NewRockData(),
		IJK:                    -1,
		IMinusOneJK:            -1,
		IPlusOneJK:             -1,
		IJMinusOneK:            -1,
		IJPlusOneK:             -1,
		IJKMinusOne:            -1,
		IJKPlusOne:             -1,
		AdjResidualDerivatives: make(map[int]ResidualDerivatives),
		Wells:                  make([]WellData, 0),
		FluidData:              NewFluidData(),
		SaturationData:         NewSaturationData(),
		TransmisibilityData:    NewTransmisibilityData(),
		Acumulation:            NewAcumulation(),
		PressureData:           NewPressureData(),
		Residual:               NewResidual(),
		ResidualDerivatives:    NewResidualDerivatives(),
		Boundary:               NewBoundary(),
		VolumeInPlace:          NewVolumeInPlace(),
	}
}

type GridOrder int

const (
	XYZ GridOrder = iota + 1
	XZY
	YXZ
	YZX
	ZXY
	ZYX
)

func Blocks_OrderBy_XYZ(nx int, ny int, nz int) map[string]BlockData {

	k := 0
	j := 0
	i := 0
	num := -1
	indicies := make(map[string]int)
	for k = 0; k < nz; k++ {
		for j = 0; j < ny; j++ {
			for i = 0; i < nx; i++ {
				num++
				index := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k)
				indicies[index] = num
			}
		}
	}

	num = -1
	A := make(map[string]BlockData)
	for k = 0; k < nz; k++ {
		for j = 0; j < ny; j++ {
			for i = 0; i < nx; i++ {
				num++
				a := NewBlockData()
				index := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k)
				indexW := strconv.Itoa(i-1) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k)
				indexE := strconv.Itoa(i+1) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k)
				indexS := strconv.Itoa(i) + "_" + strconv.Itoa(j-1) + "_" + strconv.Itoa(k)
				indexN := strconv.Itoa(i) + "_" + strconv.Itoa(j+1) + "_" + strconv.Itoa(k)
				indexTop := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k-1)
				indexBottom := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k+1)

				a.IJK = indicies[index]
				a.I = i
				a.J = j
				a.K = k
				a.KeyIndex = index

				if i-1 < 0 {
					a.IMinusOneJK = -1
				} else {
					a.IMinusOneJK = indicies[indexW]
				}

				if i+1 > nx-1 {
					a.IPlusOneJK = -1
				} else {
					a.IPlusOneJK = indicies[indexE]
				}

				if j-1 < 0 {
					a.IJMinusOneK = -1
				} else {
					a.IJMinusOneK = indicies[indexS]
				}

				if j+1 > ny-1 {
					a.IJPlusOneK = -1
				} else {
					a.IJPlusOneK = indicies[indexN]
				}

				if k-1 < 0 {
					a.IJKMinusOne = -1
				} else {
					a.IJKMinusOne = indicies[indexTop]
				}

				if k+1 > nz-1 {
					a.IJKPlusOne = -1
				} else {
					a.IJKPlusOne = indicies[indexBottom]
				}

				A[index] = a
			}
		}
	}

	return A
}
