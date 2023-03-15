package cfdporomedia

import (
	mathematicsLibrary "github.com/GabrielAchumba/school-mgt-backend/MathematicsLibrary"
	DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"
)

type Coefficients struct {
	Blocks        []DataStructure.BlockData
	Block         DataStructure.BlockData
	interpolation *mathematicsLibrary.Interpolation
	PVTO          DataStructure.PVTO
	PVTG          DataStructure.PVTG
	PVTW          DataStructure.PVTW
}

func NewCoefficients(blocks []DataStructure.BlockData,
	block DataStructure.BlockData, PVTO DataStructure.PVTO,
	PVTG DataStructure.PVTG, PVTW DataStructure.PVTW) Coefficients {
	return Coefficients{
		Blocks:        blocks,
		Block:         block,
		interpolation: mathematicsLibrary.NewInterpolation(),
		PVTO:          PVTO,
		PVTG:          PVTG,
		PVTW:          PVTW,
	}
}

func (impl *Coefficients) CalculateCoefficientsOldTime() {

	n := len(impl.Blocks)

	for i := 0; i < n; i++ {

		impl.Blocks[i].FluidData.GasData.Pressure_old = impl.Blocks[i].PressureData.OldTimePressureData.GasPressure
		impl.Blocks[i].FluidData.GasData.FVF_old = impl.interpolation.LinearInterpolation(
			impl.PVTG.PGAS, impl.PVTG.BGAS, impl.Blocks[i].PressureData.OldTimePressureData.GasPressure)
		impl.Blocks[i].FluidData.GasData.Viscosity_old = impl.interpolation.LinearInterpolation(
			impl.PVTG.PGAS, impl.PVTG.VISGAS, impl.Blocks[i].PressureData.OldTimePressureData.GasPressure)
		impl.Blocks[i].FluidData.GasData.Density_old = impl.interpolation.LinearInterpolation(
			impl.PVTG.PGAS, impl.PVTG.DENGAS, impl.Blocks[i].PressureData.OldTimePressureData.GasPressure)
		impl.Blocks[i].FluidData.GasData.Compressibility_old = impl.interpolation.LinearInterpolation(
			impl.PVTG.PGAS, impl.PVTG.COMPRESSIBILITY, impl.Blocks[i].PressureData.OldTimePressureData.GasPressure)

		impl.Blocks[i].FluidData.OilData.Pressure_old = impl.Blocks[i].PressureData.OldTimePressureData.OilPressure
		impl.Blocks[i].FluidData.OilData.SolutionGOR_old = impl.interpolation.LinearInterpolation(
			impl.PVTO.POIL, impl.PVTO.RS, impl.Blocks[i].PressureData.OldTimePressureData.OilPressure)
		impl.Blocks[i].FluidData.OilData.FVF_old = impl.interpolation.LinearInterpolation(
			impl.PVTO.POIL, impl.PVTO.FVFO, impl.Blocks[i].PressureData.OldTimePressureData.OilPressure)
		impl.Blocks[i].FluidData.OilData.Viscosity_old = impl.interpolation.LinearInterpolation(
			impl.PVTO.POIL, impl.PVTO.VISO, impl.Blocks[i].PressureData.OldTimePressureData.OilPressure)
		impl.Blocks[i].FluidData.OilData.Density_old = impl.interpolation.LinearInterpolation(
			impl.PVTO.POIL, impl.PVTO.DENSITYOIL, impl.Blocks[i].PressureData.OldTimePressureData.OilPressure)
		impl.Blocks[i].FluidData.OilData.Compressibility_old = impl.interpolation.LinearInterpolation(
			impl.PVTO.POIL, impl.PVTO.COMPRESSIBILITY, impl.Blocks[i].PressureData.OldTimePressureData.OilPressure)

		impl.Blocks[i].FluidData.WaterData.Pressure_old = impl.Blocks[i].PressureData.OldTimePressureData.WaterPressure
		impl.Blocks[i].FluidData.WaterData.FVF_old = impl.interpolation.LinearInterpolation(
			impl.PVTW.PRES, impl.PVTW.FVF, impl.Blocks[i].PressureData.OldTimePressureData.WaterPressure)
		impl.Blocks[i].FluidData.WaterData.Viscosity_old = impl.interpolation.LinearInterpolation(
			impl.PVTW.PRES, impl.PVTW.VISCOSITY, impl.Blocks[i].PressureData.OldTimePressureData.WaterPressure)
		impl.Blocks[i].FluidData.WaterData.Density_old = impl.interpolation.LinearInterpolation(
			impl.PVTW.PRES, impl.PVTW.DENSITYWATER, impl.Blocks[i].PressureData.OldTimePressureData.WaterPressure)
		impl.Blocks[i].FluidData.WaterData.Compressibility_old = impl.interpolation.LinearInterpolation(
			impl.PVTW.PRES, impl.PVTW.COMPRESSIBILITY, impl.Blocks[i].PressureData.OldTimePressureData.WaterPressure)
	}
}

func (impl *Coefficients) CalculateCoefficientOldTime() {

	impl.Block.FluidData.GasData.Pressure_old = impl.Block.PressureData.OldTimePressureData.GasPressure
	impl.Block.FluidData.GasData.FVF_old = impl.interpolation.LinearInterpolation(
		impl.PVTG.PGAS, impl.PVTG.BGAS, impl.Block.PressureData.OldTimePressureData.GasPressure)
	impl.Block.FluidData.GasData.Viscosity_old = impl.interpolation.LinearInterpolation(
		impl.PVTG.PGAS, impl.PVTG.VISGAS, impl.Block.PressureData.OldTimePressureData.GasPressure)
	impl.Block.FluidData.GasData.Density_old = impl.interpolation.LinearInterpolation(
		impl.PVTG.PGAS, impl.PVTG.DENGAS, impl.Block.PressureData.OldTimePressureData.GasPressure)
	impl.Block.FluidData.GasData.Compressibility_old = impl.interpolation.LinearInterpolation(
		impl.PVTG.PGAS, impl.PVTG.COMPRESSIBILITY, impl.Block.PressureData.OldTimePressureData.GasPressure)

	impl.Block.FluidData.OilData.Pressure_old = impl.Block.PressureData.OldTimePressureData.OilPressure
	impl.Block.FluidData.OilData.SolutionGOR_old = impl.interpolation.LinearInterpolation(
		impl.PVTO.POIL, impl.PVTO.RS, impl.Block.PressureData.OldTimePressureData.OilPressure)
	impl.Block.FluidData.OilData.FVF_old = impl.interpolation.LinearInterpolation(
		impl.PVTO.POIL, impl.PVTO.FVFO, impl.Block.PressureData.OldTimePressureData.OilPressure)
	impl.Block.FluidData.OilData.Viscosity_old = impl.interpolation.LinearInterpolation(
		impl.PVTO.POIL, impl.PVTO.VISO, impl.Block.PressureData.OldTimePressureData.OilPressure)
	impl.Block.FluidData.OilData.Density_old = impl.interpolation.LinearInterpolation(
		impl.PVTO.POIL, impl.PVTO.DENSITYOIL, impl.Block.PressureData.OldTimePressureData.OilPressure)
	impl.Block.FluidData.OilData.Compressibility_old = impl.interpolation.LinearInterpolation(
		impl.PVTO.POIL, impl.PVTO.COMPRESSIBILITY, impl.Block.PressureData.OldTimePressureData.OilPressure)

	impl.Block.FluidData.WaterData.Pressure_old = impl.Block.PressureData.OldTimePressureData.WaterPressure
	impl.Block.FluidData.WaterData.FVF_old = impl.interpolation.LinearInterpolation(
		impl.PVTW.PRES, impl.PVTW.FVF, impl.Block.PressureData.OldTimePressureData.WaterPressure)
	impl.Block.FluidData.WaterData.Viscosity_old = impl.interpolation.LinearInterpolation(
		impl.PVTW.PRES, impl.PVTW.VISCOSITY, impl.Block.PressureData.OldTimePressureData.WaterPressure)
	impl.Block.FluidData.WaterData.Density_old = impl.interpolation.LinearInterpolation(
		impl.PVTW.PRES, impl.PVTW.DENSITYWATER, impl.Block.PressureData.OldTimePressureData.WaterPressure)
	impl.Block.FluidData.WaterData.Compressibility_old = impl.interpolation.LinearInterpolation(
		impl.PVTW.PRES, impl.PVTW.COMPRESSIBILITY, impl.Block.PressureData.OldTimePressureData.WaterPressure)
}

func (impl *Coefficients) CalculateCoefficientsNewTime() {

	n := len(impl.Blocks)

	for i := 0; i < n; i++ {

		impl.Blocks[i].FluidData.GasData.Pressure_itr = impl.Blocks[i].PressureData.IterationPressureData.GasPressure
		impl.Blocks[i].FluidData.GasData.FVF_itr = impl.interpolation.LinearInterpolation(
			impl.PVTG.PGAS, impl.PVTG.BGAS, impl.Blocks[i].PressureData.IterationPressureData.GasPressure)
		impl.Blocks[i].FluidData.GasData.Viscosity_itr = impl.interpolation.LinearInterpolation(
			impl.PVTG.PGAS, impl.PVTG.VISGAS, impl.Blocks[i].PressureData.IterationPressureData.GasPressure)
		impl.Blocks[i].FluidData.GasData.Density_itr = impl.interpolation.LinearInterpolation(
			impl.PVTG.PGAS, impl.PVTG.DENGAS, impl.Blocks[i].PressureData.IterationPressureData.GasPressure)
		impl.Blocks[i].FluidData.GasData.Compressibility_itr = impl.interpolation.LinearInterpolation(
			impl.PVTG.PGAS, impl.PVTG.COMPRESSIBILITY, impl.Blocks[i].PressureData.IterationPressureData.GasPressure)

		impl.Blocks[i].FluidData.OilData.Pressure_itr = impl.Blocks[i].PressureData.IterationPressureData.OilPressure
		impl.Blocks[i].FluidData.OilData.SolutionGOR_itr = impl.interpolation.LinearInterpolation(
			impl.PVTO.POIL, impl.PVTO.RS, impl.Blocks[i].PressureData.IterationPressureData.OilPressure)
		impl.Blocks[i].FluidData.OilData.FVF_itr = impl.interpolation.LinearInterpolation(
			impl.PVTO.POIL, impl.PVTO.FVFO, impl.Blocks[i].PressureData.IterationPressureData.OilPressure)
		impl.Blocks[i].FluidData.OilData.Viscosity_itr = impl.interpolation.LinearInterpolation(
			impl.PVTO.POIL, impl.PVTO.VISO, impl.Blocks[i].PressureData.IterationPressureData.OilPressure)
		impl.Blocks[i].FluidData.OilData.Density_itr = impl.interpolation.LinearInterpolation(
			impl.PVTO.POIL, impl.PVTO.DENSITYOIL, impl.Blocks[i].PressureData.IterationPressureData.OilPressure)
		impl.Blocks[i].FluidData.OilData.Compressibility_itr = impl.interpolation.LinearInterpolation(
			impl.PVTO.POIL, impl.PVTO.COMPRESSIBILITY, impl.Blocks[i].PressureData.IterationPressureData.OilPressure)

		impl.Blocks[i].FluidData.WaterData.Pressure_itr = impl.Blocks[i].PressureData.IterationPressureData.WaterPressure
		impl.Blocks[i].FluidData.WaterData.FVF_itr = impl.interpolation.LinearInterpolation(
			impl.PVTW.PRES, impl.PVTW.FVF, impl.Blocks[i].PressureData.IterationPressureData.WaterPressure)
		impl.Blocks[i].FluidData.WaterData.Viscosity_itr = impl.interpolation.LinearInterpolation(
			impl.PVTW.PRES, impl.PVTW.VISCOSITY, impl.Blocks[i].PressureData.IterationPressureData.WaterPressure)
		impl.Blocks[i].FluidData.WaterData.Density_itr = impl.interpolation.LinearInterpolation(
			impl.PVTW.PRES, impl.PVTW.DENSITYWATER, impl.Blocks[i].PressureData.IterationPressureData.WaterPressure)
		impl.Blocks[i].FluidData.WaterData.Compressibility_itr = impl.interpolation.LinearInterpolation(
			impl.PVTW.PRES, impl.PVTW.COMPRESSIBILITY, impl.Blocks[i].PressureData.IterationPressureData.WaterPressure)
	}
}

func (impl *Coefficients) CalculateCoefficientNewTime() {

	impl.Block.FluidData.GasData.Pressure_itr = impl.Block.PressureData.IterationPressureData.GasPressure
	impl.Block.FluidData.GasData.FVF_itr = impl.interpolation.LinearInterpolation(
		impl.PVTG.PGAS, impl.PVTG.BGAS, impl.Block.PressureData.IterationPressureData.GasPressure)
	impl.Block.FluidData.GasData.Viscosity_itr = impl.interpolation.LinearInterpolation(
		impl.PVTG.PGAS, impl.PVTG.VISGAS, impl.Block.PressureData.IterationPressureData.GasPressure)
	impl.Block.FluidData.GasData.Density_itr = impl.interpolation.LinearInterpolation(
		impl.PVTG.PGAS, impl.PVTG.DENGAS, impl.Block.PressureData.IterationPressureData.GasPressure)
	impl.Block.FluidData.GasData.Compressibility_itr = impl.interpolation.LinearInterpolation(
		impl.PVTG.PGAS, impl.PVTG.COMPRESSIBILITY, impl.Block.PressureData.IterationPressureData.GasPressure)

	impl.Block.FluidData.OilData.Pressure_itr = impl.Block.PressureData.IterationPressureData.OilPressure
	impl.Block.FluidData.OilData.SolutionGOR_itr = impl.interpolation.LinearInterpolation(
		impl.PVTO.POIL, impl.PVTO.RS, impl.Block.PressureData.IterationPressureData.OilPressure)
	impl.Block.FluidData.OilData.FVF_itr = impl.interpolation.LinearInterpolation(
		impl.PVTO.POIL, impl.PVTO.FVFO, impl.Block.PressureData.IterationPressureData.OilPressure)
	impl.Block.FluidData.OilData.Viscosity_itr = impl.interpolation.LinearInterpolation(
		impl.PVTO.POIL, impl.PVTO.VISO, impl.Block.PressureData.IterationPressureData.OilPressure)
	impl.Block.FluidData.OilData.Density_itr = impl.interpolation.LinearInterpolation(
		impl.PVTO.POIL, impl.PVTO.DENSITYOIL, impl.Block.PressureData.IterationPressureData.OilPressure)
	impl.Block.FluidData.OilData.Compressibility_itr = impl.interpolation.LinearInterpolation(
		impl.PVTO.POIL, impl.PVTO.COMPRESSIBILITY, impl.Block.PressureData.IterationPressureData.OilPressure)

	impl.Block.FluidData.WaterData.Pressure_itr = impl.Block.PressureData.IterationPressureData.WaterPressure
	impl.Block.FluidData.WaterData.FVF_itr = impl.interpolation.LinearInterpolation(
		impl.PVTW.PRES, impl.PVTW.FVF, impl.Block.PressureData.IterationPressureData.WaterPressure)
	impl.Block.FluidData.WaterData.Viscosity_itr = impl.interpolation.LinearInterpolation(
		impl.PVTW.PRES, impl.PVTW.VISCOSITY, impl.Block.PressureData.IterationPressureData.WaterPressure)
	impl.Block.FluidData.WaterData.Density_itr = impl.interpolation.LinearInterpolation(
		impl.PVTW.PRES, impl.PVTW.DENSITYWATER, impl.Block.PressureData.IterationPressureData.WaterPressure)
	impl.Block.FluidData.WaterData.Compressibility_itr = impl.interpolation.LinearInterpolation(
		impl.PVTW.PRES, impl.PVTW.COMPRESSIBILITY, impl.Block.PressureData.IterationPressureData.WaterPressure)
}
