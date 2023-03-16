package datastructure

type SpaceDistributions struct {
	PressureDistribution        []float64
	OilSaturationDistribution   []float64
	GasSaturationDistribution   []float64
	WaterSaturationDistribution []float64
}

func NewDistributions(P []float64, So []float64, Sg []float64, Sw []float64) SpaceDistributions {
	return SpaceDistributions{
		PressureDistribution:        P,
		OilSaturationDistribution:   So,
		GasSaturationDistribution:   Sg,
		WaterSaturationDistribution: Sw,
	}
}
