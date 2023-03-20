package results

import (
	CFDPoroMedia "github.com/GabrielAchumba/school-mgt-backend/Simulation/CFD_PoroMedia"
	DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"
	NumMthds "github.com/GabrielAchumba/school-mgt-backend/Simulation/NumericalMethods"
)

type SimulationLogViewModelImpl struct {
	CG                 CFDPoroMedia.CartGrid
	DIMENS_            DataStructure.DIMENS
	Rock               CFDPoroMedia.Rock
	PVTO               DataStructure.PVTO
	PVTG               DataStructure.PVTG
	PVTW               DataStructure.PVTW
	Times              []float64
	Pressures          []float64
	Wells              map[string]DataStructure.WellData
	Boundaries         map[string]DataStructure.Boundary
	SpaceDistributions []DataStructure.SpaceDistributions
	SimulationLogs     string
}

func New(DIMENS_ DataStructure.DIMENS, Rock CFDPoroMedia.Rock,
	PVTO DataStructure.PVTO, PVTG DataStructure.PVTG,
	PVTW DataStructure.PVTW, Times []float64,
	Pressures []float64, Wells map[string]DataStructure.WellData,
	Boundaries map[string]DataStructure.Boundary) SimulationLogViewModelImpl {

	return SimulationLogViewModelImpl{
		DIMENS_:    DIMENS_,
		Rock:       Rock,
		PVTO:       PVTO,
		PVTG:       PVTG,
		PVTW:       PVTW,
		Times:      Times,
		Pressures:  Pressures,
		Wells:      Wells,
		Boundaries: Boundaries,
	}
}

func (impl *SimulationLogViewModelImpl) InputData_For_Simulation() {

	//=============1. Define geometry=====================//
	nx := impl.DIMENS_.Nx
	ny := impl.DIMENS_.Ny
	nz := impl.DIMENS_.Nz
	impl.CG = CFDPoroMedia.NewCatGrid(nx, ny, nz)
	Lx := impl.DIMENS_.Lx
	Ly := impl.DIMENS_.Ly
	Lz := impl.DIMENS_.Lz
	TVD := impl.DIMENS_.TVD
	nGrid := nx * ny * nz

	//===================================================//

	//=======================2. Process geometry============//
	_DX := NumMthds.NewMatD()
	_DX.EqualSegments(0, Lx, nx)
	_DY := NumMthds.NewMatD()
	_DY.EqualSegments(0, Ly, ny)
	_DZ := NumMthds.NewMatD()
	_DZ.EqualSegments(0, Lz, nz)
	impl.CG.Compute(_DX, _DY, _DZ, TVD)

	//==================================================//

	//===================3. Set rock data============================//

	_PORO := NumMthds.NewMatD()
	_PORO.Duplicate(0, impl.Rock.Porosity, nGrid)
	_PERMX := NumMthds.NewMatD()
	_PERMX.Duplicate(0, impl.Rock.PermeabilityX, nGrid)
	_PERMY := NumMthds.NewMatD()
	_PERMY.Duplicate(0, impl.Rock.PermeabilityY, nGrid)
	_PERMZ := NumMthds.NewMatD()
	_PERMZ.Duplicate(0, impl.Rock.PermeabilityZ, nGrid)
	CR := NumMthds.NewMatD()
	CR.Duplicate(0, impl.Rock.Compressibility, nGrid)
	SRD := CFDPoroMedia.NewRockData(impl.CG)
	SRD.SetRockData(_PORO, _PERMX, _PERMY, _PERMZ, CR)
	impl.CG = SRD.CartGrid

	//=================================================================//

	//===================4. Set Reservoir  Pressure============================//

	Pressures := CFDPoroMedia.NewPressures(impl.CG)
	Pressures.SetOldPressures(impl.Pressures)
	Pressures.SetNewPressures(impl.Pressures)
	impl.CG = Pressures.CartGrid

	//=================================================================//

	//===================5. Set Wells============================//

	WellsData := CFDPoroMedia.NewWellsData(impl.CG, impl.Wells)
	WellsData.SetWellsData()
	impl.CG = WellsData.CartGrid

	//=================================================================//

	//===================5. Rel Perm============================//

	RelPermData := CFDPoroMedia.NewSlightlyCompressibleRelPerm(impl.CG)
	RelPermData.SetSlightlyCompressibleRelPerm()
	impl.CG = RelPermData.CartGrid

	//=================================================================//
}

func (impl *SimulationLogViewModelImpl) InputData_For_Simulation2() {

	//=============1. Define geometry=====================//
	nx := impl.DIMENS_.Nx
	ny := impl.DIMENS_.Ny
	nz := impl.DIMENS_.Nz
	impl.CG = CFDPoroMedia.NewCatGrid(nx, ny, nz)
	Lx := impl.DIMENS_.Lx
	Ly := impl.DIMENS_.Ly
	Lz := impl.DIMENS_.Lz
	TVD := impl.DIMENS_.TVD
	nGrid := nx * ny * nz

	//===================================================//

	//=======================2. Process geometry============//
	_DX := NumMthds.NewMatD()
	_DX.EqualSegments(0, Lx, nx)
	_DY := NumMthds.NewMatD()
	_DY.EqualSegments(0, Ly, ny)
	_DZ := NumMthds.NewMatD()
	_DZ.EqualSegments(0, Lz, nz)
	impl.CG.Compute(_DX, _DY, _DZ, TVD)

	//==================================================//

	//===================3. Set rock data============================//

	_PORO := NumMthds.NewMatD()
	if impl.Rock.PorosityDataType == CFDPoroMedia.Scalar {
		_PORO.Duplicate(0, impl.Rock.Porosity, nGrid)
	} else {
		_PORO.Mat = impl.Rock.Porosity_Vector
	}

	_PERMX := NumMthds.NewMatD()
	if impl.Rock.PermeabilityXDataType == CFDPoroMedia.Scalar {
		_PERMX.Duplicate(0, impl.Rock.PermeabilityX, nGrid)
	} else {
		_PERMX.Mat = impl.Rock.PermeabilityX_Vector
	}

	_PERMY := NumMthds.NewMatD()
	if impl.Rock.PermeabilityYDataType == CFDPoroMedia.Scalar {
		_PERMY.Duplicate(0, impl.Rock.PermeabilityY, nGrid)
	} else {
		_PERMY.Mat = impl.Rock.PermeabilityY_Vector
	}

	_PERMZ := NumMthds.NewMatD()
	if impl.Rock.PermeabilityZDataType == CFDPoroMedia.Scalar {
		_PERMZ.Duplicate(0, impl.Rock.PermeabilityZ, nGrid)
	} else {
		_PERMZ.Mat = impl.Rock.PermeabilityZ_Vector
	}

	CR := NumMthds.NewMatD()
	if impl.Rock.RockCompDataType == CFDPoroMedia.Scalar {
		CR.Duplicate(0, impl.Rock.Compressibility, nGrid)
	} else {
		CR.Mat = impl.Rock.Compressibility_Vector
	}

	SRD := CFDPoroMedia.NewRockData(impl.CG)
	SRD.SetRockData(_PORO, _PERMX, _PERMY, _PERMZ, CR)
	impl.CG = SRD.CartGrid

	//=================================================================//

	//===================4. Set Reservoir  Pressure============================//

	Pressures := CFDPoroMedia.NewPressures(impl.CG)
	Pressures.SetOldPressures(impl.Pressures)
	Pressures.SetNewPressures(impl.Pressures)
	impl.CG = Pressures.CartGrid

	//=================================================================//

	//===================5. Set Wells============================//

	WellsData := CFDPoroMedia.NewWellsData(impl.CG, impl.Wells)
	WellsData.SetWellsData()
	impl.CG = WellsData.CartGrid

	//=================================================================//

	//===================6. Rel Perm============================//

	RelPermData := CFDPoroMedia.NewSlightlyCompressibleRelPerm(impl.CG)
	RelPermData.SetSlightlyCompressibleRelPerm()
	impl.CG = RelPermData.CartGrid

	//=================================================================//

	//===================7. Set Reservoir Outer Boundaries============================//

	BoundariesData := CFDPoroMedia.NewBoundaryData(impl.CG, impl.Boundaries)
	BoundariesData.SetBoundaryDataData()
	impl.CG = WellsData.CartGrid

	//=================================================================//

}

func (impl *SimulationLogViewModelImpl) InputData_For_Simulation3() {

	//=============1. Define geometry=====================//
	nx := impl.DIMENS_.Nx
	ny := impl.DIMENS_.Ny
	nz := impl.DIMENS_.Nz
	impl.CG = CFDPoroMedia.NewCatGrid(nx, ny, nz)
	nGrid := nx * ny * nz

	//===================================================//

	//=======================2. Process geometry============//
	_DX := NumMthds.NewMatD()
	_DX.Mat = impl.DIMENS_.DxVec
	_DY := NumMthds.NewMatD()
	_DY.Mat = impl.DIMENS_.DyVec
	_DZ := NumMthds.NewMatD()
	_DZ.Mat = impl.DIMENS_.DzVec
	var TVD float64 = 0 // equal surface depth;
	impl.CG.Compute(_DX, _DY, _DZ, TVD)

	//==================================================//

	//===================3. Set rock data============================//

	_PORO := NumMthds.NewMatD()
	if impl.Rock.PorosityDataType == CFDPoroMedia.Scalar {
		_PORO.Duplicate(0, impl.Rock.Porosity, nGrid)
	} else {
		_PORO.Mat = impl.Rock.Porosity_Vector
	}

	_PERMX := NumMthds.NewMatD()
	if impl.Rock.PermeabilityXDataType == CFDPoroMedia.Scalar {
		_PERMX.Duplicate(0, impl.Rock.PermeabilityX, nGrid)
	} else {
		_PERMX.Mat = impl.Rock.PermeabilityX_Vector
	}

	_PERMY := NumMthds.NewMatD()
	if impl.Rock.PermeabilityYDataType == CFDPoroMedia.Scalar {
		_PERMY.Duplicate(0, impl.Rock.PermeabilityY, nGrid)
	} else {
		_PERMY.Mat = impl.Rock.PermeabilityY_Vector
	}

	_PERMZ := NumMthds.NewMatD()
	if impl.Rock.PermeabilityZDataType == CFDPoroMedia.Scalar {
		_PERMZ.Duplicate(0, impl.Rock.PermeabilityZ, nGrid)
	} else {
		_PERMZ.Mat = impl.Rock.PermeabilityZ_Vector
	}

	CR := NumMthds.NewMatD()
	if impl.Rock.RockCompDataType == CFDPoroMedia.Scalar {
		CR.Duplicate(0, impl.Rock.Compressibility, nGrid)
	} else {
		CR.Mat = impl.Rock.Compressibility_Vector
	}

	SRD := CFDPoroMedia.NewRockData(impl.CG)
	SRD.SetRockData(_PORO, _PERMX, _PERMY, _PERMZ, CR)
	impl.CG = SRD.CartGrid

	//=================================================================//

	//===================4. Set Reservoir  Pressure============================//

	Pressures := CFDPoroMedia.NewPressures(impl.CG)
	Pressures.SetOldPressures(impl.Pressures)
	Pressures.SetNewPressures(impl.Pressures)
	impl.CG = Pressures.CartGrid

	//=================================================================//

	//===================5. Set Wells============================//

	WellsData := CFDPoroMedia.NewWellsData(impl.CG, impl.Wells)
	WellsData.SetWellsData()
	impl.CG = WellsData.CartGrid

	//=================================================================//

	//===================6. Rel Perm============================//

	RelPermData := CFDPoroMedia.NewSlightlyCompressibleRelPerm(impl.CG)
	RelPermData.SetSlightlyCompressibleRelPerm()
	impl.CG = RelPermData.CartGrid

	//=================================================================//

	//===================7. Set Reservoir Outer Boundaries============================//

	BoundariesData := CFDPoroMedia.NewBoundaryData(impl.CG, impl.Boundaries)
	BoundariesData.SetBoundaryDataData()
	impl.CG = WellsData.CartGrid

	//=================================================================//

}

func (impl *SimulationLogViewModelImpl) SolveSlightlyCompressible() {

	SlightlyCompressible := CFDPoroMedia.NewSlightlyCompressible(impl.CG.Blocks, impl.PVTO,
		impl.PVTG, impl.PVTW, impl.Times, impl.Wells)

	SlightlyCompressible.Simulate()
	impl.SpaceDistributions = SlightlyCompressible.SpaceDistributions
	impl.SimulationLogs = SlightlyCompressible.SimulationLogs
}
