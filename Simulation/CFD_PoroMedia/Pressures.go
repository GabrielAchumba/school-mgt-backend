package cfdporomedia

import (
	"strconv"

	DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"
)

type Pressures struct {
	CartGrid CartGrid
	Blocks   []DataStructure.BlockData
}

func NewPressures(CartGrid CartGrid) Pressures {
	return Pressures{
		CartGrid: CartGrid,
	}
}

func NewPressures2(Blocks []DataStructure.BlockData) Pressures {
	return Pressures{
		Blocks: Blocks,
	}
}

func (impl *Pressures) SetOldPressures(pressures []float64) {
	counter := -1
	i := 0
	j := 0
	k := 0
	for k = 0; k < impl.CartGrid.Nz; k++ {
		for j = 0; j < impl.CartGrid.Ny; j++ {
			for i = 0; i < impl.CartGrid.Nx; i++ {
				counter++
				index := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k)

				block := impl.CartGrid.Blocks[index]
				block.PressureData.OldTimePressureData.GasPressure = pressures[counter]
				block.PressureData.OldTimePressureData.OilPressure = pressures[counter]
				block.PressureData.OldTimePressureData.WaterPressure = pressures[counter]
				impl.CartGrid.Blocks[index] = block
			}
		}
	}
}

func (impl *Pressures) SetOldPressures2(pressures []float64) {

	for i := 0; i < len(impl.Blocks); i++ {
		impl.Blocks[i].PressureData.OldTimePressureData.GasPressure = pressures[i]
		impl.Blocks[i].PressureData.OldTimePressureData.OilPressure = pressures[i]
		impl.Blocks[i].PressureData.OldTimePressureData.WaterPressure = pressures[i]
	}
}

func (impl *Pressures) SetNewPressures(pressures []float64) {
	counter := -1
	i := 0
	j := 0
	k := 0
	for k = 0; k < impl.CartGrid.Nz; k++ {
		for j = 0; j < impl.CartGrid.Ny; j++ {
			for i = 0; i < impl.CartGrid.Nx; i++ {
				counter++
				index := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k)

				block := impl.CartGrid.Blocks[index]
				block.PressureData.IterationPressureData.GasPressure = pressures[counter]
				block.PressureData.IterationPressureData.OilPressure = pressures[counter]
				block.PressureData.IterationPressureData.WaterPressure = pressures[counter]
				impl.CartGrid.Blocks[index] = block
			}
		}
	}
}

func (impl *Pressures) SetNewPressures2(pressures []float64) {

	for i := 0; i < len(impl.Blocks); i++ {
		impl.Blocks[i].PressureData.IterationPressureData.GasPressure = pressures[i]
		impl.Blocks[i].PressureData.IterationPressureData.OilPressure = pressures[i]
		impl.Blocks[i].PressureData.IterationPressureData.WaterPressure = pressures[i]
	}
}
