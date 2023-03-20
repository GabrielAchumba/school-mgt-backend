package datastructure

type WellType int

const (
	GasProducer WellType = iota + 1
	OilProducer
	GasInjector
	WaterInjector
)

type WellCondition int

const (
	ConstantRate WellType = iota + 1
	ConstantBHP
	LiftCurves
)

type WellData struct {
	WellType                     WellType
	OilRate                      float64
	GasRate                      float64
	Water                        float64
	BottomHolePressureDatumDepth float64
	AveragePermeability          float64
	WellBoreRadius               float64
	BlockDrainageRadius          float64
	PerforationIntervals         []PerforationInterval
	WellCondition                WellCondition
	SkinFactor                   float64
	WellName                     string
}

type PerforationInterval struct {
	SegmentLength float64
}

type WellReport struct {
	ProductionTime            float64 `json:"productionTime"`
	ReservoirPressure         float64 `json:"reservoirPressure"`
	FlowingBottomHolePressure float64 `json:"flowingBottomHolePressure"`
	OilRate                   float64 `json:"oilRate"`
	GasRate                   float64 `json:"gasRate"`
	WaterRate                 float64 `json:"waterRate"`
	WellName                  string  `json:"wellName"`
	ProductivityIndex         float64 `json:"productivityIndex"`
}

func NewWellReport() WellReport {
	return WellReport{}
}
