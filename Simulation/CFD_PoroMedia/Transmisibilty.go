package cfdporomedia

import (
	"math"

	mathematicsLibrary "github.com/GabrielAchumba/school-mgt-backend/MathematicsLibrary"
	DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"
)

func TransmisibilityOil(mBlock DataStructure.BlockData,
	block_w DataStructure.BlockData, block_e DataStructure.BlockData,
	block_s DataStructure.BlockData, block_n DataStructure.BlockData,
	block_t DataStructure.BlockData, block_b DataStructure.BlockData) DataStructure.OilTransmisibiity {

	var averaging mathematicsLibrary.Averaging = mathematicsLibrary.NewAveraging()

	var T_adj float64 = 0

	beta := 1.127 * math.Pow(10, -3)
	OilTransmisibiity := DataStructure.NewOilTransmisibiity()

	//=====West Flow===========================
	T_c := (beta * mBlock.Geometry.Area_x * mBlock.RockData.Kx * mBlock.FluidData.RelPermData.Kro_itr) /
		(mBlock.FluidData.OilData.Viscosity_itr * mBlock.FluidData.OilData.FVF_itr * mBlock.Geometry.Dx)

	T_adj = (beta * block_w.Geometry.Area_x * block_w.RockData.Kx * block_w.FluidData.RelPermData.Kro_itr) /
		(block_w.FluidData.OilData.Viscosity_itr * block_w.FluidData.OilData.FVF_itr * block_w.Geometry.Dx)
	OilTransmisibiity.To_West = averaging.Harmonic([]float64{T_c, T_adj})

	if mBlock.IMinusOneJK == -1 {
		switch mBlock.Boundary.West.BoundaryType {
		case DataStructure.Closed:
			OilTransmisibiity.To_West = 0
		case DataStructure.KnownFlowRate:
			OilTransmisibiity.To_West = 0
		}
	}

	//==============East Flow ============================================
	T_c = (beta * mBlock.Geometry.Area_x * mBlock.RockData.Kx * mBlock.FluidData.RelPermData.Kro_itr) /
		(mBlock.FluidData.OilData.Viscosity_itr * mBlock.FluidData.OilData.FVF_itr * mBlock.Geometry.Dx)

	T_adj = (beta * block_e.Geometry.Area_x * block_e.RockData.Kx * block_e.FluidData.RelPermData.Kro_itr) /
		(block_e.FluidData.OilData.Viscosity_itr * block_e.FluidData.OilData.FVF_itr * block_e.Geometry.Dx)
	OilTransmisibiity.To_East = averaging.Harmonic([]float64{T_c, T_adj})

	if mBlock.IPlusOneJK == -1 {
		switch mBlock.Boundary.East.BoundaryType {
		case DataStructure.Closed:
			OilTransmisibiity.To_East = 0
		case DataStructure.KnownFlowRate:
			OilTransmisibiity.To_East = 0
		}
	}

	//=======================South FLow========================================
	T_c = (beta * mBlock.Geometry.Area_y * mBlock.RockData.Ky * mBlock.FluidData.RelPermData.Kro_itr) /
		(mBlock.FluidData.OilData.Viscosity_itr * mBlock.FluidData.OilData.FVF_itr * mBlock.Geometry.Dy)

	T_adj = (beta * block_s.Geometry.Area_y * block_s.RockData.Ky * block_s.FluidData.RelPermData.Kro_itr) /
		(block_s.FluidData.OilData.Viscosity_itr * block_s.FluidData.OilData.FVF_itr * block_s.Geometry.Dy)
	OilTransmisibiity.To_South = averaging.Harmonic([]float64{T_c, T_adj})

	if mBlock.IJMinusOneK == -1 {
		switch mBlock.Boundary.South.BoundaryType {
		case DataStructure.Closed:
			OilTransmisibiity.To_South = 0
		case DataStructure.KnownFlowRate:
			OilTransmisibiity.To_South = 0
		}
	}

	//=======================North FLow========================================
	T_c = (beta * mBlock.Geometry.Area_y * mBlock.RockData.Ky * mBlock.FluidData.RelPermData.Kro_itr) /
		(mBlock.FluidData.OilData.Viscosity_itr * mBlock.FluidData.OilData.FVF_itr * mBlock.Geometry.Dy)

	T_adj = (beta * block_n.Geometry.Area_y * block_n.RockData.Ky * block_n.FluidData.RelPermData.Kro_itr) /
		(block_n.FluidData.OilData.Viscosity_itr * block_n.FluidData.OilData.FVF_itr * block_n.Geometry.Dy)
	OilTransmisibiity.To_North = averaging.Harmonic([]float64{T_c, T_adj})

	if mBlock.IJPlusOneK == -1 {
		switch mBlock.Boundary.North.BoundaryType {
		case DataStructure.Closed:
			OilTransmisibiity.To_North = 0
		case DataStructure.KnownFlowRate:
			OilTransmisibiity.To_North = 0
		}
	}

	//=======================Top FLow========================================
	T_c = (beta * mBlock.Geometry.Area_z * mBlock.RockData.Kz * mBlock.FluidData.RelPermData.Kro_itr) /
		(mBlock.FluidData.OilData.Viscosity_itr * mBlock.FluidData.OilData.FVF_itr * mBlock.Geometry.Dz)

	T_adj = (beta * block_t.Geometry.Area_z * block_t.RockData.Kz * block_t.FluidData.RelPermData.Kro_itr) /
		(block_t.FluidData.OilData.Viscosity_itr * block_t.FluidData.OilData.FVF_itr * block_t.Geometry.Dz)
	OilTransmisibiity.To_Top = averaging.Harmonic([]float64{T_c, T_adj})

	if mBlock.IJKMinusOne == -1 {
		switch mBlock.Boundary.Top.BoundaryType {
		case DataStructure.Closed:
			OilTransmisibiity.To_Top = 0
		case DataStructure.KnownFlowRate:
			OilTransmisibiity.To_Top = 0
		}
	}

	//=======================Bottom FLow========================================
	T_c = (beta * mBlock.Geometry.Area_z * mBlock.RockData.Kz * mBlock.FluidData.RelPermData.Kro_itr) /
		(mBlock.FluidData.OilData.Viscosity_itr * mBlock.FluidData.OilData.FVF_itr * mBlock.Geometry.Dz)

	T_adj = (beta * block_b.Geometry.Area_z * block_b.RockData.Kz * block_b.FluidData.RelPermData.Kro_itr) /
		(block_b.FluidData.OilData.Viscosity_itr * block_b.FluidData.OilData.FVF_itr * block_b.Geometry.Dz)
	OilTransmisibiity.To_Bottom = averaging.Harmonic([]float64{T_c, T_adj})

	if mBlock.IJKPlusOne == -1 {
		switch mBlock.Boundary.Bottom.BoundaryType {
		case DataStructure.Closed:
			OilTransmisibiity.To_Bottom = 0
		case DataStructure.KnownFlowRate:
			OilTransmisibiity.To_Bottom = 0
		}
	}
	return OilTransmisibiity
}

func TransmisibilityGas(mBlock DataStructure.BlockData,
	block_w DataStructure.BlockData, block_e DataStructure.BlockData,
	block_s DataStructure.BlockData, block_n DataStructure.BlockData,
	block_t DataStructure.BlockData, block_b DataStructure.BlockData) DataStructure.GasTransmisibiity {

	GasTransmisibiity := DataStructure.NewGasTransmisibiity()
	T_c := mBlock.FluidData.RelPermData.Krg_itr /
		(mBlock.FluidData.GasData.Viscosity_itr * mBlock.FluidData.GasData.FVF_itr)

	var T_adj float64 = 0

	T_adj = block_w.FluidData.RelPermData.Krg_itr /
		(block_w.FluidData.GasData.Viscosity_itr * block_w.FluidData.GasData.FVF_itr)
	GasTransmisibiity.Tg_West = (T_c + T_adj) / 2

	T_adj = block_e.FluidData.RelPermData.Krg_itr /
		(block_e.FluidData.GasData.Viscosity_itr * block_e.FluidData.GasData.FVF_itr)
	GasTransmisibiity.Tg_East = (T_c + T_adj) / 2

	T_adj = block_s.FluidData.RelPermData.Krg_itr /
		(block_s.FluidData.GasData.Viscosity_itr * block_s.FluidData.GasData.FVF_itr)
	GasTransmisibiity.Tg_South = (T_c + T_adj) / 2

	T_adj = block_n.FluidData.RelPermData.Krg_itr /
		(block_n.FluidData.GasData.Viscosity_itr * block_n.FluidData.GasData.FVF_itr)
	GasTransmisibiity.Tg_North = (T_c + T_adj) / 2

	T_adj = block_t.FluidData.RelPermData.Krg_itr /
		(block_t.FluidData.GasData.Viscosity_itr * block_t.FluidData.GasData.FVF_itr)
	GasTransmisibiity.Tg_South = (T_c + T_adj) / 2

	T_adj = block_b.FluidData.RelPermData.Krg_itr /
		(block_b.FluidData.GasData.Viscosity_itr * block_b.FluidData.GasData.FVF_itr)
	GasTransmisibiity.Tg_North = (T_c + T_adj) / 2

	return GasTransmisibiity
}

func TransmisibilityWater(mBlock DataStructure.BlockData,
	block_w DataStructure.BlockData, block_e DataStructure.BlockData,
	block_s DataStructure.BlockData, block_n DataStructure.BlockData,
	block_t DataStructure.BlockData, block_b DataStructure.BlockData) DataStructure.WaterTransmisibility {

	WaterTransmisibility := DataStructure.NewWaterTransmisibility()
	T_c := mBlock.FluidData.RelPermData.Krw_itr /
		(mBlock.FluidData.WaterData.Viscosity_itr * mBlock.FluidData.WaterData.FVF_itr)

	var T_adj float64 = 0
	T_adj = block_w.FluidData.RelPermData.Krw_itr /
		(block_w.FluidData.WaterData.Viscosity_itr * block_w.FluidData.WaterData.FVF_itr)
	WaterTransmisibility.Tw_West = (T_c + T_adj) / 2

	T_adj = block_e.FluidData.RelPermData.Krw_itr /
		(block_e.FluidData.WaterData.Viscosity_itr * block_e.FluidData.WaterData.FVF_itr)
	WaterTransmisibility.Tw_East = (T_c + T_adj) / 2

	T_adj = block_s.FluidData.RelPermData.Krw_itr /
		(block_s.FluidData.WaterData.Viscosity_itr * block_s.FluidData.WaterData.FVF_itr)
	WaterTransmisibility.Tw_South = (T_c + T_adj) / 2

	T_adj = block_n.FluidData.RelPermData.Krw_itr /
		(block_n.FluidData.WaterData.Viscosity_itr * block_n.FluidData.WaterData.FVF_itr)
	WaterTransmisibility.Tw_North = (T_c + T_adj) / 2

	T_adj = block_t.FluidData.RelPermData.Krw_itr /
		(block_t.FluidData.WaterData.Viscosity_itr * block_t.FluidData.WaterData.FVF_itr)
	WaterTransmisibility.Tw_South = (T_c + T_adj) / 2

	T_adj = block_b.FluidData.RelPermData.Krw_itr /
		(block_b.FluidData.WaterData.Viscosity_itr * block_b.FluidData.WaterData.FVF_itr)
	WaterTransmisibility.Tw_North = (T_c + T_adj) / 2

	return WaterTransmisibility
}
