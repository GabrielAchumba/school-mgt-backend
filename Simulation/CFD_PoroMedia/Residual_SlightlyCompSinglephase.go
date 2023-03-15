package cfdporomedia

import DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"

type Residual_SlightlyCompSinglephase struct {
	Block DataStructure.BlockData
}

func NewResidual_SlightlyCompSinglephase(Block DataStructure.BlockData) Residual_SlightlyCompSinglephase {
	return Residual_SlightlyCompSinglephase{
		Block: Block,
	}
}

func (impl *Residual_SlightlyCompSinglephase) Oil(blocks []DataStructure.BlockData, mBlock DataStructure.BlockData, dt float64) {

	block_w := DataStructure.NewBlockData()
	block_e := DataStructure.NewBlockData()
	block_s := DataStructure.NewBlockData()
	block_n := DataStructure.NewBlockData()
	block_t := DataStructure.NewBlockData()
	block_b := DataStructure.NewBlockData()

	if mBlock.IMinusOneJK != -1 {
		block_w = blocks[mBlock.IMinusOneJK]
	} else {
		block_w = blocks[mBlock.IJK]
	}

	if mBlock.IPlusOneJK != -1 {
		block_e = blocks[mBlock.IPlusOneJK]
	} else {
		block_e = blocks[mBlock.IJK]
	}

	if mBlock.IJMinusOneK != -1 {
		block_s = blocks[mBlock.IJMinusOneK]
	} else {
		block_s = blocks[mBlock.IJK]
	}

	if mBlock.IJPlusOneK != -1 {
		block_n = blocks[mBlock.IJPlusOneK]
	} else {
		block_n = blocks[mBlock.IJK]
	}

	if mBlock.IJKMinusOne != -1 {
		block_t = blocks[mBlock.IJKMinusOne]
	} else {
		block_t = blocks[mBlock.IJK]
	}

	if mBlock.IJKPlusOne != -1 {
		block_b = blocks[mBlock.IJKPlusOne]
	} else {
		block_b = blocks[mBlock.IJK]
	}

	acc, zigma := SlightlyCompresibleAcumulation(mBlock, dt)
	impl.Block.Acumulation.OilVolumeStored = acc
	impl.Block.Zigma = zigma
	impl.Block.TransmisibilityData.OilTransmisibiity = TransmisibilityOil(mBlock, block_w, block_e,
		block_s, block_n, block_t, block_b)

	impl.Block.Flux.OilFlux.Flux_West = impl.Block.TransmisibilityData.OilTransmisibiity.To_West *
		(block_w.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	impl.Block.Flux.OilFlux.Flux_East = impl.Block.TransmisibilityData.OilTransmisibiity.To_East *
		(block_e.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	impl.Block.Flux.OilFlux.Flux_South = impl.Block.TransmisibilityData.OilTransmisibiity.To_South *
		(block_s.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	impl.Block.Flux.OilFlux.Flux_North = impl.Block.TransmisibilityData.OilTransmisibiity.To_North *
		(block_n.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	impl.Block.Flux.OilFlux.Flux_Top = impl.Block.TransmisibilityData.OilTransmisibiity.To_Top *
		(block_t.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	impl.Block.Flux.OilFlux.Flux_Bottom = impl.Block.TransmisibilityData.OilTransmisibiity.To_Bottom *
		(block_b.PressureData.IterationPressureData.OilPressure - impl.Block.PressureData.IterationPressureData.OilPressure)

	//Calculate water rates from wells

	wells := NewWells(impl.Block)
	wells.SumWellRates()

	impl.Block.Wells = wells.mBlock.Wells
	impl.Block.Qosc = wells.mBlock.Qosc
	impl.Block.Wells_LHS = wells.mBlock.Wells_LHS
	impl.Block.Wells_RHS = wells.mBlock.Wells_RHS
	impl.Block.Residual.Oil = impl.Block.Acumulation.OilVolumeStored -
		impl.Block.Flux.OilFlux.Flux_West - impl.Block.Flux.OilFlux.Flux_East -
		impl.Block.Flux.OilFlux.Flux_South - impl.Block.Flux.OilFlux.Flux_North -
		impl.Block.Flux.OilFlux.Flux_Top - impl.Block.Flux.OilFlux.Flux_Bottom -
		impl.Block.Qosc

}
