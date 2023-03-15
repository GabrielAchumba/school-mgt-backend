package cfdporomedia

import DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"

type Residual_Multiphase struct {
	Block DataStructure.BlockData
}

func NewResidual(Block DataStructure.BlockData) Residual_Multiphase {
	return Residual_Multiphase{
		Block: Block,
	}
}

func (impl *Residual_Multiphase) Water(blocks []DataStructure.BlockData, mBlock DataStructure.BlockData, dt float64) {

	block_w := DataStructure.NewBlockData()
	block_e := DataStructure.NewBlockData()
	block_s := DataStructure.NewBlockData()
	block_n := DataStructure.NewBlockData()
	block_t := DataStructure.NewBlockData()
	block_b := DataStructure.NewBlockData()

	if mBlock.IMinusOneJK > -1 {
		block_w = blocks[mBlock.IJK]
	} else {
		block_w = blocks[mBlock.IJKMinusOne]
	}

	if mBlock.IPlusOneJK > -1 {
		block_e = blocks[mBlock.IPlusOneJK]
	} else {
		block_e = blocks[mBlock.IJK]
	}

	if mBlock.IJMinusOneK > -1 {
		block_s = blocks[mBlock.IJMinusOneK]
	} else {
		block_s = blocks[mBlock.IJK]
	}

	if mBlock.IJPlusOneK > -1 {
		block_n = blocks[mBlock.IJPlusOneK]
	} else {
		block_n = blocks[mBlock.IJK]
	}

	if mBlock.IJKMinusOne > -1 {
		block_t = blocks[mBlock.IJKMinusOne]
	} else {
		block_t = blocks[mBlock.IJK]
	}

	if mBlock.IJKPlusOne > -1 {
		block_b = blocks[mBlock.IJKPlusOne]
	} else {
		block_b = blocks[mBlock.IJK]
	}

	impl.Block.Acumulation.WaterVolumeStored = 0

	impl.Block.Acumulation.WaterVolumeStored = WaterVolumeAccumulationIteration(mBlock, dt)

	impl.Block.TransmisibilityData.WaterTransmisibility = TransmisibilityWater(mBlock, block_w, block_e,
		block_s, block_n, block_t, block_b)

	impl.Block.Flux.WaterFlux.Flux_West = impl.Block.TransmisibilityData.WaterTransmisibility.Tw_West *
		(block_w.PressureData.IterationPressureData.WaterPressure - impl.Block.PressureData.IterationPressureData.WaterPressure)

	impl.Block.Flux.WaterFlux.Flux_East = impl.Block.TransmisibilityData.WaterTransmisibility.Tw_East *
		(block_e.PressureData.IterationPressureData.WaterPressure - impl.Block.PressureData.IterationPressureData.WaterPressure)

	impl.Block.Flux.WaterFlux.Flux_South = impl.Block.TransmisibilityData.WaterTransmisibility.Tw_South *
		(block_s.PressureData.IterationPressureData.WaterPressure - impl.Block.PressureData.IterationPressureData.WaterPressure)

	impl.Block.Flux.WaterFlux.Flux_South = impl.Block.TransmisibilityData.WaterTransmisibility.Tw_North *
		(block_n.PressureData.IterationPressureData.WaterPressure - impl.Block.PressureData.IterationPressureData.WaterPressure)

	impl.Block.Flux.WaterFlux.Flux_Top = impl.Block.TransmisibilityData.WaterTransmisibility.Tw_Top *
		(block_t.PressureData.IterationPressureData.WaterPressure - impl.Block.PressureData.IterationPressureData.WaterPressure)

	impl.Block.Flux.WaterFlux.Flux_Bottom = impl.Block.TransmisibilityData.WaterTransmisibility.Tw_Bottom *
		(block_b.PressureData.IterationPressureData.WaterPressure - impl.Block.PressureData.IterationPressureData.WaterPressure)

	//Calculate water rates from wells

	impl.Block.Residual.Water = impl.Block.Acumulation.WaterVolumeStored -
		impl.Block.Flux.WaterFlux.Flux_West - impl.Block.Flux.WaterFlux.Flux_East -
		impl.Block.Flux.WaterFlux.Flux_South - impl.Block.Flux.WaterFlux.Flux_North -
		impl.Block.Flux.WaterFlux.Flux_Top - impl.Block.Flux.WaterFlux.Flux_Bottom

}

func (impl *Residual_Multiphase) Oil(blocks []DataStructure.BlockData, mBlock DataStructure.BlockData, dt float64) {

	block_w := DataStructure.NewBlockData()
	block_e := DataStructure.NewBlockData()
	block_s := DataStructure.NewBlockData()
	block_n := DataStructure.NewBlockData()
	block_t := DataStructure.NewBlockData()
	block_b := DataStructure.NewBlockData()

	if mBlock.IMinusOneJK > -1 {
		block_w = blocks[mBlock.IJK]
	} else {
		block_w = blocks[mBlock.IJKMinusOne]
	}

	if mBlock.IPlusOneJK > -1 {
		block_e = blocks[mBlock.IPlusOneJK]
	} else {
		block_e = blocks[mBlock.IJK]
	}

	if mBlock.IJMinusOneK > -1 {
		block_s = blocks[mBlock.IJMinusOneK]
	} else {
		block_s = blocks[mBlock.IJK]
	}

	if mBlock.IJPlusOneK > -1 {
		block_n = blocks[mBlock.IJPlusOneK]
	} else {
		block_n = blocks[mBlock.IJK]
	}

	if mBlock.IJKMinusOne > -1 {
		block_t = blocks[mBlock.IJKMinusOne]
	} else {
		block_t = blocks[mBlock.IJK]
	}

	if mBlock.IJKPlusOne > -1 {
		block_b = blocks[mBlock.IJKPlusOne]
	} else {
		block_b = blocks[mBlock.IJK]
	}

	impl.Block.Acumulation.OilVolumeStored = OilVolumeAccumulationIteration(mBlock, dt)
	impl.Block.TransmisibilityData.OilTransmisibiity = TransmisibilityOil(mBlock, block_w, block_e,
		block_s, block_n, block_t, block_b)

	impl.Block.Flux.OilFlux.Flux_West = impl.Block.TransmisibilityData.OilTransmisibiity.To_West *
		(block_w.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	impl.Block.Flux.OilFlux.Flux_East = impl.Block.TransmisibilityData.OilTransmisibiity.To_East *
		(block_e.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	impl.Block.Flux.OilFlux.Flux_South = impl.Block.TransmisibilityData.OilTransmisibiity.To_South *
		(block_s.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	impl.Block.Flux.OilFlux.Flux_South = impl.Block.TransmisibilityData.OilTransmisibiity.To_North *
		(block_n.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	impl.Block.Flux.OilFlux.Flux_Top = impl.Block.TransmisibilityData.OilTransmisibiity.To_Top *
		(block_t.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	impl.Block.Flux.OilFlux.Flux_Bottom = impl.Block.TransmisibilityData.OilTransmisibiity.To_Bottom *
		(block_b.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	//Calculate water rates from wells

	impl.Block.Residual.Oil = impl.Block.Acumulation.OilVolumeStored -
		impl.Block.Flux.OilFlux.Flux_West - impl.Block.Flux.OilFlux.Flux_East -
		impl.Block.Flux.OilFlux.Flux_South - impl.Block.Flux.OilFlux.Flux_North -
		impl.Block.Flux.OilFlux.Flux_Top - impl.Block.Flux.OilFlux.Flux_Bottom

}

func (impl *Residual_Multiphase) Gas(blocks []DataStructure.BlockData, mBlock DataStructure.BlockData, dt float64) {

	block_w := DataStructure.NewBlockData()
	block_e := DataStructure.NewBlockData()
	block_s := DataStructure.NewBlockData()
	block_n := DataStructure.NewBlockData()
	block_t := DataStructure.NewBlockData()
	block_b := DataStructure.NewBlockData()

	if mBlock.IMinusOneJK > -1 {
		block_w = blocks[mBlock.IJK]
	} else {
		block_w = blocks[mBlock.IJKMinusOne]
	}

	if mBlock.IPlusOneJK > -1 {
		block_e = blocks[mBlock.IPlusOneJK]
	} else {
		block_e = blocks[mBlock.IJK]
	}

	if mBlock.IJMinusOneK > -1 {
		block_s = blocks[mBlock.IJMinusOneK]
	} else {
		block_s = blocks[mBlock.IJK]
	}

	if mBlock.IJPlusOneK > -1 {
		block_n = blocks[mBlock.IJPlusOneK]
	} else {
		block_n = blocks[mBlock.IJK]
	}

	if mBlock.IJKMinusOne > -1 {
		block_t = blocks[mBlock.IJKMinusOne]
	} else {
		block_t = blocks[mBlock.IJK]
	}

	if mBlock.IJKPlusOne > -1 {
		block_b = blocks[mBlock.IJKPlusOne]
	} else {
		block_b = blocks[mBlock.IJK]
	}

	impl.Block.Acumulation.GasVolumeStored = GasVolumeAccumulationIteration(mBlock, dt)
	impl.Block.TransmisibilityData.GasTransmisibiity = TransmisibilityGas(mBlock, block_w, block_e,
		block_s, block_n, block_t, block_b)

	impl.Block.Flux.GasFlux.Flux_West = impl.Block.TransmisibilityData.GasTransmisibiity.Tg_West *
		(block_w.PressureData.IterationPressureData.GasPressure - impl.Block.PressureData.IterationPressureData.GasPressure)

	impl.Block.Flux.GasFlux.Flux_East = impl.Block.TransmisibilityData.GasTransmisibiity.Tg_East *
		(block_e.PressureData.IterationPressureData.GasPressure - impl.Block.PressureData.IterationPressureData.GasPressure)

	impl.Block.Flux.GasFlux.Flux_South = impl.Block.TransmisibilityData.GasTransmisibiity.Tg_South *
		(block_s.PressureData.IterationPressureData.GasPressure - impl.Block.PressureData.IterationPressureData.GasPressure)

	impl.Block.Flux.GasFlux.Flux_South = impl.Block.TransmisibilityData.GasTransmisibiity.Tg_North *
		(block_n.PressureData.IterationPressureData.GasPressure - impl.Block.PressureData.IterationPressureData.GasPressure)

	impl.Block.Flux.GasFlux.Flux_Top = impl.Block.TransmisibilityData.GasTransmisibiity.Tg_Top *
		(block_t.PressureData.IterationPressureData.GasPressure - impl.Block.PressureData.IterationPressureData.GasPressure)

	impl.Block.Flux.GasFlux.Flux_Bottom = impl.Block.TransmisibilityData.GasTransmisibiity.Tg_Bottom *
		(block_b.PressureData.IterationPressureData.GasPressure - impl.Block.PressureData.IterationPressureData.GasPressure)

	impl.Block.Residual.Gas = impl.Block.Acumulation.GasVolumeStored -
		impl.Block.Flux.GasFlux.Flux_West - impl.Block.Flux.GasFlux.Flux_East -
		impl.Block.Flux.GasFlux.Flux_South - impl.Block.Flux.GasFlux.Flux_North -
		impl.Block.Flux.GasFlux.Flux_Top - impl.Block.Flux.GasFlux.Flux_Bottom

}
