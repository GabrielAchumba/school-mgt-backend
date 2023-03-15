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
}

type PerforationInterval struct {
	SegmentLength float64
}
