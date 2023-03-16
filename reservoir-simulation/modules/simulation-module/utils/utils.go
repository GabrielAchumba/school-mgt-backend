package simulationmodule

import (
	"strconv"

	CFDPoroMedia "github.com/GabrielAchumba/school-mgt-backend/Simulation/CFD_PoroMedia"
	DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"
	numericalmethods "github.com/GabrielAchumba/school-mgt-backend/Simulation/NumericalMethods"
	SimulationResults "github.com/GabrielAchumba/school-mgt-backend/Simulation/Results"
	"github.com/GabrielAchumba/school-mgt-backend/reservoir-simulation/modules/simulation-module/dtos"
)

func Simulate(inputData dtos.SimulationInputDTO) ([]DataStructure.SpaceDistributions, string) {
	Dimensions := DataStructure.DIMENS{
		Nx:    int(inputData.Geometry.Dimensions.Nx),
		Ny:    int(inputData.Geometry.Dimensions.Ny),
		Nz:    int(inputData.Geometry.Dimensions.Nz),
		DxVec: inputData.Geometry.Gridding.DxVector,
		DyVec: inputData.Geometry.Gridding.DyVector,
		DzVec: inputData.Geometry.Gridding.DzVector,
		TVD:   0,
	}

	//var zero float64 = 0 //math.Pow(10, -15)
	rock := CFDPoroMedia.Rock{
		PorosityDataType:       CFDPoroMedia.Vector,
		Porosity_Vector:        inputData.Rock.Porosity.PorosityArray,
		PermeabilityXDataType:  CFDPoroMedia.Vector,
		PermeabilityX_Vector:   inputData.Rock.Permeability.PermeabilityXArray,
		PermeabilityYDataType:  CFDPoroMedia.Vector,
		PermeabilityY_Vector:   inputData.Rock.Permeability.PermeabilityYArray,
		PermeabilityZDataType:  CFDPoroMedia.Vector,
		PermeabilityZ_Vector:   inputData.Rock.Permeability.PermeabilityZArray,
		RockCompDataType:       CFDPoroMedia.Vector,
		Compressibility_Vector: inputData.Rock.Compressibility.CompressibilityArray,
	}

	PVTO := DataStructure.PVTO{
		POIL:            inputData.Pvt.Oil.Pressure,
		RS:              inputData.Pvt.Oil.SolutionGOR,
		FVFO:            inputData.Pvt.Oil.FVF,
		VISO:            inputData.Pvt.Oil.Viscosity,
		DENSITYOIL:      inputData.Pvt.Oil.Density,
		COMPRESSIBILITY: inputData.Pvt.Oil.Compressibility,
	}

	PVTG := DataStructure.PVTG{
		PGAS:            inputData.Pvt.Gas.Pressure,
		BGAS:            inputData.Pvt.Gas.FVF,
		VISGAS:          inputData.Pvt.Gas.Viscosity,
		DENGAS:          inputData.Pvt.Gas.Density,
		COMPRESSIBILITY: inputData.Pvt.Gas.Compressibility,
	}

	PVTW := DataStructure.PVTW{
		PRES:            inputData.Pvt.Water.Pressure,
		FVF:             inputData.Pvt.Water.FVF,
		COMPRESSIBILITY: inputData.Pvt.Water.Compressibility,
		VISCOSITY:       inputData.Pvt.Water.Density,
		DENSITYWATER:    inputData.Pvt.Water.Density,
	}

	var wells map[string]DataStructure.WellData = make(map[string]DataStructure.WellData)

	for i := 0; i < len(inputData.Wells); i++ {
		key := strconv.Itoa(inputData.Wells[i].I) + "_" +
			strconv.Itoa(inputData.Wells[i].J) + "_" +
			strconv.Itoa(inputData.Wells[i].K)
		wells[key] = DataStructure.WellData{
			WellType:                     GetWellType(inputData.Wells[i].WellType),
			BottomHolePressureDatumDepth: inputData.Wells[i].BottomHolePressureDatumDepth,
			WellBoreRadius:               inputData.Wells[i].WellBoreRadius,
			OilRate:                      inputData.Wells[i].OilRate,
			GasRate:                      inputData.Wells[i].GasRate,
			Water:                        inputData.Wells[i].WaterRate,
			SkinFactor:                   0,
			PerforationIntervals: []DataStructure.PerforationInterval{
				{SegmentLength: Dimensions.DzVec[i]},
			},
			WellCondition: GetWellControl(inputData.Wells[i].WellCondition),
		}
	}

	counter := -1
	var boundaries map[string]DataStructure.Boundary = make(map[string]DataStructure.Boundary)
	for k := 0; k < Dimensions.Nz; k++ {
		for j := 0; j < Dimensions.Ny; j++ {
			for i := 0; i < Dimensions.Nx; i++ {
				counter++
				key := strconv.Itoa(i) + "_" + strconv.Itoa(j) + "_" + strconv.Itoa(k)

				boundary := DataStructure.NewBoundary()
				west := DataStructure.NewWest()
				east := DataStructure.NewEast()
				south := DataStructure.NewSouth()
				north := DataStructure.NewNorth()
				top := DataStructure.NewTop()
				bottom := DataStructure.NewBottom()

				boundary.West = west
				boundary.East = east
				boundary.South = south
				boundary.North = north
				boundary.Top = top
				boundary.Bottom = bottom
				boundaries[key] = boundary
			}
		}
	}

	for k := 0; k < len(inputData.Boundaries.West); k++ {
		block := inputData.Boundaries.West[k]
		key := strconv.Itoa(block.I) + "_" + strconv.Itoa(block.J) + "_" + strconv.Itoa(block.K)
		boundary := boundaries[key]
		boundary.West.BoundaryType = GetBoundaryCondition(block.BoundaryCondition)
		boundary.West.FlowRate = block.FlowRate
		boundary.West.PressureGradient = block.PressureGradient
		boundary.West.ConstantPressure = block.ConstantPressure
		boundaries[key] = boundary
	}

	for k := 0; k < len(inputData.Boundaries.East); k++ {
		block := inputData.Boundaries.East[k]
		key := strconv.Itoa(block.I) + "_" + strconv.Itoa(block.J) + "_" + strconv.Itoa(block.K)
		boundary := boundaries[key]
		boundary.East.BoundaryType = GetBoundaryCondition(block.BoundaryCondition)
		boundary.East.FlowRate = block.FlowRate
		boundary.East.PressureGradient = block.PressureGradient
		boundary.East.ConstantPressure = block.ConstantPressure
		boundaries[key] = boundary
	}

	for k := 0; k < len(inputData.Boundaries.South); k++ {
		block := inputData.Boundaries.South[k]
		key := strconv.Itoa(block.I) + "_" + strconv.Itoa(block.J) + "_" + strconv.Itoa(block.K)
		boundary := boundaries[key]
		boundary.South.BoundaryType = GetBoundaryCondition(block.BoundaryCondition)
		boundary.South.FlowRate = block.FlowRate
		boundary.South.PressureGradient = block.PressureGradient
		boundary.South.ConstantPressure = block.ConstantPressure
		boundaries[key] = boundary
	}

	for k := 0; k < len(inputData.Boundaries.North); k++ {
		block := inputData.Boundaries.North[k]
		key := strconv.Itoa(block.I) + "_" + strconv.Itoa(block.J) + "_" + strconv.Itoa(block.K)
		boundary := boundaries[key]
		boundary.North.BoundaryType = GetBoundaryCondition(block.BoundaryCondition)
		boundary.North.FlowRate = block.FlowRate
		boundary.North.PressureGradient = block.PressureGradient
		boundary.North.ConstantPressure = block.ConstantPressure
		boundaries[key] = boundary
	}

	for k := 0; k < len(inputData.Boundaries.Top); k++ {
		block := inputData.Boundaries.Top[k]
		key := strconv.Itoa(block.I) + "_" + strconv.Itoa(block.J) + "_" + strconv.Itoa(block.K)
		boundary := boundaries[key]
		boundary.Top.BoundaryType = GetBoundaryCondition(block.BoundaryCondition)
		boundary.Top.FlowRate = block.FlowRate
		boundary.Top.PressureGradient = block.PressureGradient
		boundary.Top.ConstantPressure = block.ConstantPressure
		boundaries[key] = boundary
	}

	for k := 0; k < len(inputData.Boundaries.Bottom); k++ {
		block := inputData.Boundaries.Bottom[k]
		key := strconv.Itoa(block.I) + "_" + strconv.Itoa(block.J) + "_" + strconv.Itoa(block.K)
		boundary := boundaries[key]
		boundary.Bottom.BoundaryType = GetBoundaryCondition(block.BoundaryCondition)
		boundary.Bottom.FlowRate = block.FlowRate
		boundary.Bottom.PressureGradient = block.PressureGradient
		boundary.Bottom.ConstantPressure = block.ConstantPressure
		boundaries[key] = boundary
	}

	Times := numericalmethods.NewMatD()
	Times.Mat = inputData.Schedule.CumulativeTime

	Pressure := numericalmethods.NewMatD()
	Pressure.Mat = inputData.Initialization.Pressure

	SimRes := SimulationResults.New(Dimensions, rock, PVTO, PVTG, PVTW,
		Times.Mat, Pressure.Mat, wells, boundaries)

	SimRes.InputData_For_Simulation3()
	SimRes.SolveSlightlyCompressible()

	return SimRes.SpaceDistributions, SimRes.SimulationLogs

}

func GetWellType(wellType string) DataStructure.WellType {
	_wellType := DataStructure.OilProducer

	switch wellType {
	case "Gas Producer":
		_wellType = DataStructure.GasProducer
	case "Oil Producer":
		_wellType = DataStructure.OilProducer
	case "Gas Injector":
		_wellType = DataStructure.GasInjector
	case "Water Injector":
		_wellType = DataStructure.WaterInjector
	}
	return _wellType
}

func GetWellControl(wellControl string) DataStructure.WellCondition {
	_wellControl := DataStructure.WellCondition(DataStructure.ConstantRate)

	switch wellControl {
	case "Constant Rate":
		_wellControl = DataStructure.WellCondition(DataStructure.ConstantRate)
	case "Constant BHP":
		_wellControl = DataStructure.WellCondition(DataStructure.ConstantBHP)
	}
	return _wellControl
}

func GetBoundaryCondition(boundaryType string) DataStructure.BoundaryType {
	_boundaryType := DataStructure.BoundaryType(DataStructure.Closed)

	switch boundaryType {
	case "Closed":
		_boundaryType = DataStructure.BoundaryType(DataStructure.Closed)
	case "Constant Gradient":
		_boundaryType = DataStructure.BoundaryType(DataStructure.ConstantGradient)
	case "Known Flow Rate":
		_boundaryType = DataStructure.BoundaryType(DataStructure.KnownFlowRate)
	case "Constant Pressure":
		_boundaryType = DataStructure.BoundaryType(DataStructure.ConstantPressure)
	}
	return _boundaryType
}
