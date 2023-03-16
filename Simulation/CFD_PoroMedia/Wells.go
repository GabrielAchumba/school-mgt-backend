package cfdporomedia

import (
	"math"

	DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"
)

type Wells struct {
	mBlock DataStructure.BlockData
}

func NewWells(mBlock DataStructure.BlockData) Wells {
	return Wells{
		mBlock: mBlock,
	}
}

func (impl *Wells) AverageHorizontalPermeability(Kx float64, Ky float64) float64 {
	KH := math.Sqrt(Kx * Ky)
	return KH
}

func (impl *Wells) DrainageRadius(_Kx float64, _Ky float64, Dx float64, Dy float64) float64 {
	Kx := _Kx / 1000
	Ky := _Ky / 1000
	KyKxSqr2 := math.Sqrt(Ky / Kx)
	KyKxSqr4 := math.Pow(Ky/Kx, 0.25)
	KxKySqr2 := math.Sqrt(Kx / Ky)
	KxKySqr4 := math.Pow(Kx/Ky, 0.25)
	DxPow2 := math.Pow(Dx, 2)
	DyPow2 := math.Pow(Dy, 2)
	numerator := math.Sqrt(KyKxSqr2*DxPow2 + KxKySqr2*DyPow2)

	re := 0.28 * numerator / (KyKxSqr4 + KxKySqr4)
	return re
}

func (impl *Wells) SumWellRates() {

	beta := 1.127 * math.Pow(10, -3)
	for i := 0; i < len(impl.mBlock.Wells); i++ {
		switch impl.mBlock.Wells[i].WellCondition {
		case DataStructure.WellCondition(DataStructure.ConstantRate):
			switch impl.mBlock.Wells[i].WellType {
			case DataStructure.GasProducer:
				impl.mBlock.Qgsc = impl.mBlock.Qgsc - impl.mBlock.Wells[i].GasRate
			case DataStructure.GasInjector:
				impl.mBlock.Qgsc = impl.mBlock.Qgsc + impl.mBlock.Wells[i].GasRate
			case DataStructure.OilProducer:
				impl.mBlock.Qosc = impl.mBlock.Qosc - impl.mBlock.Wells[i].OilRate
			case DataStructure.WaterInjector:
				impl.mBlock.Qwsc = impl.mBlock.Qwsc + impl.mBlock.Wells[i].Water
			}

		case DataStructure.WellCondition(DataStructure.ConstantBHP):
			kH := impl.AverageHorizontalPermeability(impl.mBlock.RockData.Kx, impl.mBlock.RockData.Ky)
			re := impl.DrainageRadius(impl.mBlock.RockData.Kx, impl.mBlock.RockData.Ky,
				impl.mBlock.Geometry.Dx, impl.mBlock.Geometry.Dy)
			switch impl.mBlock.Wells[i].WellType {
			case DataStructure.OilProducer:
				mobility := impl.mBlock.FluidData.RelPermData.Kro_old /
					(impl.mBlock.FluidData.OilData.Viscosity_old * impl.mBlock.FluidData.OilData.FVF_old)
				var WI float64 = 0
				for j := 0; j < len(impl.mBlock.Wells[i].PerforationIntervals); j++ {
					rerw := re / impl.mBlock.Wells[i].WellBoreRadius
					WI = WI + 2*math.Pi*beta*kH*impl.mBlock.Wells[i].PerforationIntervals[j].SegmentLength/
						(math.Log(rerw)+impl.mBlock.Wells[i].SkinFactor-0.5)
				}
				PI := WI * mobility
				impl.mBlock.Wells_LHS = impl.mBlock.Wells_LHS - PI
				impl.mBlock.Wells_RHS = impl.mBlock.Wells_RHS - PI*impl.mBlock.Wells[i].BottomHolePressureDatumDepth
				/* impl.mBlock.Qosc = impl.mBlock.Qosc -
				(PI * (impl.mBlock.PressureData.OldTimePressureData.OilPressure -
					impl.mBlock.Wells[i].BottomHolePressureDatumDepth)) */
			}

		}

	}
}

type WellsData struct {
	CartGrid CartGrid
	Wells    map[string]DataStructure.WellData
}

func NewWellsData(CartGrid CartGrid, Wells map[string]DataStructure.WellData) WellsData {
	return WellsData{
		CartGrid: CartGrid,
		Wells:    Wells,
	}
}

func (impl *WellsData) SetWellsData() {

	for key, val := range impl.Wells {
		block := impl.CartGrid.Blocks[key]
		block.Wells = append(impl.CartGrid.Blocks[key].Wells, val)
		impl.CartGrid.Blocks[key] = block
	}
}
