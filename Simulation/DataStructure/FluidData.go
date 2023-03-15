package datastructure

type FluidData struct {
	OilData     OilData
	GasData     GasData
	WaterData   WaterData
	RelPermData RelPermData
}

func NewFluidData() FluidData {
	return FluidData{
		OilData:     NewOilData(),
		GasData:     NewGasData(),
		WaterData:   NewWaterData(),
		RelPermData: NewRelPermData(),
	}
}

type OilData struct {
	Pressure_old        float64
	FVF_old             float64
	Viscosity_old       float64
	Compressibility_old float64
	Density_old         float64
	SolutionGOR_old     float64
	Pressure_new        float64
	FVF_new             float64
	Viscosity_new       float64
	Compressibility_new float64
	Density_new         float64
	SolutionGOR_new     float64
	Pressure_itr        float64
	FVF_itr             float64
	Viscosity_itr       float64
	Compressibility_itr float64
	Density_itr         float64
	SolutionGOR_itr     float64
}

func NewOilData() OilData {
	return OilData{}
}

type GasData struct {
	Pressure_old        float64
	FVF_old             float64
	Viscosity_old       float64
	Compressibility_old float64
	Density_old         float64
	Pressure_new        float64
	FVF_new             float64
	Viscosity_new       float64
	Compressibility_new float64
	Density_new         float64
	Pressure_itr        float64
	FVF_itr             float64
	Viscosity_itr       float64
	Compressibility_itr float64
	Density_itr         float64
}

func NewGasData() GasData {
	return GasData{}
}

type WaterData struct {
	Pressure_old        float64
	FVF_old             float64
	Viscosity_old       float64
	Compressibility_old float64
	Density_old         float64
	Pressure_new        float64
	FVF_new             float64
	Viscosity_new       float64
	Compressibility_new float64
	Density_new         float64
	Pressure_itr        float64
	FVF_itr             float64
	Viscosity_itr       float64
	Compressibility_itr float64
	Density_itr         float64
}

func NewWaterData() WaterData {
	return WaterData{}
}

type RelPermData struct {
	Kro_old float64
	Krw_old float64
	Krg_old float64
	Kro_new float64
	Krw_new float64
	Krg_new float64
	Kro_itr float64
	Krw_itr float64
	Krg_itr float64
}

func NewRelPermData() RelPermData {
	return RelPermData{}
}
