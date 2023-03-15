package datastructure

type Acumulation struct {
	OilVolumeStored   float64
	GasVolumeStored   float64
	WaterVolumeStored float64
}

func NewAcumulation() Acumulation {
	return Acumulation{}
}
