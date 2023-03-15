package cfdporomedia

import DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"

type BoundaryData struct {
	CartGrid   CartGrid
	Boundaries map[string]DataStructure.Boundary
}

func NewBoundaryData(CartGrid CartGrid, Boundaries map[string]DataStructure.Boundary) BoundaryData {
	return BoundaryData{
		CartGrid:   CartGrid,
		Boundaries: Boundaries,
	}
}

func (impl *BoundaryData) SetBoundaryDataData() {

	for key, val := range impl.Boundaries {
		block := impl.CartGrid.Blocks[key]
		block.Boundary = val
		impl.CartGrid.Blocks[key] = block
	}
}
