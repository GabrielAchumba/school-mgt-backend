package datastructure

type SpaceDistributions struct {
	PressureDistribution        []float64    `json:"pressureDistribution"`
	OilSaturationDistribution   []float64    `json:"oilSaturationDistribution"`
	GasSaturationDistribution   []float64    `json:"gasSaturationDistribution"`
	WaterSaturationDistribution []float64    `json:"waterSaturationDistribution"`
	InitialOilInPlace           float64      `json:"initialOilInPlace"`
	InitialGasInPlace           float64      `json:"initialGasInPlace"`
	InitialWaterInPlace         float64      `json:"initialWaterInPlace"`
	WellReport                  []WellReport `json:"wellReport"`
}

func NewDistributions(P []float64, So []float64, Sg []float64, Sw []float64) SpaceDistributions {
	return SpaceDistributions{
		PressureDistribution:        P,
		OilSaturationDistribution:   So,
		GasSaturationDistribution:   Sg,
		WaterSaturationDistribution: Sw,
		WellReport:                  make([]WellReport, 0),
	}
}
