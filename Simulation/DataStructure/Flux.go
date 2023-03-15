package datastructure

type Flux struct {
	OilFlux   OilFlux
	GasFlux   GasFlux
	WaterFlux WaterFlux
}

func NewFlux() Flux {
	return Flux{
		OilFlux:   NewOilFlux(),
		GasFlux:   NewGasFlux(),
		WaterFlux: NewWaterFlux(),
	}
}

type OilFlux struct {
	Flux_West   float64
	Flux_East   float64
	Flux_South  float64
	Flux_North  float64
	Flux_Top    float64
	Flux_Bottom float64
}

func NewOilFlux() OilFlux {
	return OilFlux{}
}

type GasFlux struct {
	Flux_West   float64
	Flux_East   float64
	Flux_South  float64
	Flux_North  float64
	Flux_Top    float64
	Flux_Bottom float64
}

func NewGasFlux() GasFlux {
	return GasFlux{}
}

type WaterFlux struct {
	Flux_West   float64
	Flux_East   float64
	Flux_South  float64
	Flux_North  float64
	Flux_Top    float64
	Flux_Bottom float64
}

func NewWaterFlux() WaterFlux {
	return WaterFlux{}
}
