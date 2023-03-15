package datastructure

type PressureData struct {
	IterationPressureData IterationPressureData
	NewTimePressureData   NewTimePressureData
	OldTimePressureData   OldTimePressureData
}

func NewPressureData() PressureData {
	return PressureData{
		IterationPressureData: NewIterationPressureData(),
		NewTimePressureData:   NewTimePressureData_(),
		OldTimePressureData:   NewOldTimePressureData(),
	}
}

type OldTimePressureData struct {
	OilPressure   float64
	GasPressure   float64
	WaterPressure float64
}

func NewOldTimePressureData() OldTimePressureData {
	return OldTimePressureData{}
}

type NewTimePressureData struct {
	OilPressure   float64
	GasPressure   float64
	WaterPressure float64
}

func NewTimePressureData_() NewTimePressureData {
	return NewTimePressureData{}
}

type IterationPressureData struct {
	OilPressure   float64
	GasPressure   float64
	WaterPressure float64
}

func NewIterationPressureData() IterationPressureData {
	return IterationPressureData{}
}
