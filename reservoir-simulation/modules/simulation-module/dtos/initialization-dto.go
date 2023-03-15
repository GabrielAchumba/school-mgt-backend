package dtos

type Initialization struct {
	TypeOfInitialization string    `json:"typeOfInitialization"`
	Pressure             []float64 `json:"pressure"`
	OilSaturation        []float64 `json:"oilSaturation"`
	WaterSaturation      []float64 `json:"waterSaturation"`
	GasSaturation        []float64 `json:"gasSaturation"`
}
