package datastructure

type TransmisibilityData struct {
	GasTransmisibiity    GasTransmisibiity
	WaterTransmisibility WaterTransmisibility
	OilTransmisibiity    OilTransmisibiity
}

func NewTransmisibilityData() TransmisibilityData {
	return TransmisibilityData{
		GasTransmisibiity:    NewGasTransmisibiity(),
		WaterTransmisibility: NewWaterTransmisibility(),
		OilTransmisibiity:    NewOilTransmisibiity(),
	}
}

type OilTransmisibiity struct {
	To_West   float64
	To_East   float64
	To_South  float64
	To_North  float64
	To_Top    float64
	To_Bottom float64
}

func NewOilTransmisibiity() OilTransmisibiity {
	return OilTransmisibiity{}
}

type WaterTransmisibility struct {
	Tw_West   float64
	Tw_East   float64
	Tw_South  float64
	Tw_North  float64
	Tw_Top    float64
	Tw_Bottom float64
}

func NewWaterTransmisibility() WaterTransmisibility {
	return WaterTransmisibility{}
}

type GasTransmisibiity struct {
	Tg_West   float64
	Tg_East   float64
	Tg_South  float64
	Tg_North  float64
	Tg_Top    float64
	Tg_Bottom float64
}

func NewGasTransmisibiity() GasTransmisibiity {
	return GasTransmisibiity{}
}
