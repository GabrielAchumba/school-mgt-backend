package dtos

type WellData struct {
	I                            int     `json:"i"`
	J                            int     `json:"j"`
	K                            int     `json:"k"`
	WellType                     string  `json:"wellType"`
	OilRate                      float64 `json:"oilRate"`
	GasRate                      float64 `json:"gasRate"`
	WaterRate                    float64 `json:"waterRate"`
	BottomHolePressureDatumDepth float64 `json:"bottomHolePressureDatumDepth"`
	WellBoreRadius               float64 `json:"wellBoreRadius"`
	WellCondition                string  `json:"wellCondition"`
	SkinFactor                   float64 `json:"skinFactor"`
	Name                         string  `json:"name"`
}
