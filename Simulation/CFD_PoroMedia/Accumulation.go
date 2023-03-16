package cfdporomedia

import DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"

func WaterVolumeAccumulation(mGrid DataStructure.BlockData, dt float64) float64 {

	newPore := mGrid.RockData.Porosity * mGrid.SaturationData.Sw_new / mGrid.FluidData.WaterData.FVF_new
	oldPore := mGrid.RockData.Porosity * mGrid.SaturationData.Sw_old / mGrid.FluidData.WaterData.FVF_old
	acc := (mGrid.Geometry.Vb / dt) * (newPore - oldPore)
	return acc
}

func WaterVolumeAccumulationIteration(mGrid DataStructure.BlockData, dt float64) float64 {

	newPore := mGrid.RockData.Porosity * mGrid.SaturationData.Sw_itr / mGrid.FluidData.WaterData.FVF_itr
	oldPore := mGrid.RockData.Porosity * mGrid.SaturationData.Sw_old / mGrid.FluidData.WaterData.FVF_old
	acc := (mGrid.Geometry.Vb / dt) * (newPore - oldPore)
	return acc
}

func OilVolumeAccumulation(mGrid DataStructure.BlockData, dt float64) float64 {

	newPore := mGrid.RockData.Porosity * mGrid.SaturationData.So_new / mGrid.FluidData.OilData.FVF_new
	oldPore := mGrid.RockData.Porosity * mGrid.SaturationData.So_old / mGrid.FluidData.OilData.FVF_old
	acc := (mGrid.Geometry.Vb / dt) * (newPore - oldPore)
	return acc
}

func OilVolumeAccumulationIteration(mGrid DataStructure.BlockData, dt float64) float64 {

	newPore := mGrid.RockData.Porosity * mGrid.SaturationData.So_itr / mGrid.FluidData.OilData.FVF_itr
	oldPore := mGrid.RockData.Porosity * mGrid.SaturationData.So_old / mGrid.FluidData.OilData.FVF_old
	acc := (mGrid.Geometry.Vb / dt) * (newPore - oldPore)
	return acc
}

func GasVolumeAccumulation(mGrid DataStructure.BlockData, dt float64) float64 {

	newPore := mGrid.RockData.Porosity * mGrid.SaturationData.Sg_new / mGrid.FluidData.GasData.FVF_new
	oldPore := mGrid.RockData.Porosity * mGrid.SaturationData.Sg_old / mGrid.FluidData.GasData.FVF_old
	acc := (mGrid.Geometry.Vb / dt) * (newPore - oldPore)
	return acc
}

func GasVolumeAccumulationIteration(mGrid DataStructure.BlockData, dt float64) float64 {

	newPore := mGrid.RockData.Porosity * mGrid.SaturationData.Sg_itr / mGrid.FluidData.GasData.FVF_itr
	oldPore := mGrid.RockData.Porosity * mGrid.SaturationData.Sg_old / mGrid.FluidData.GasData.FVF_old
	acc := (mGrid.Geometry.Vb / dt) * (newPore - oldPore)
	return acc
}

func SlightlyCompresibleAcumulation(mGrid DataStructure.BlockData, dt float64) (float64, float64) {
	var ct float64 = mGrid.RockData.RockCompressibility + mGrid.FluidData.OilData.Compressibility_itr
	zigma := mGrid.Geometry.Vb * mGrid.RockData.Porosity * ct / (5.615 * mGrid.FluidData.OilData.FVF_itr * dt)
	acc := zigma * (mGrid.PressureData.IterationPressureData.OilPressure - mGrid.PressureData.OldTimePressureData.OilPressure)

	return acc, zigma
}
