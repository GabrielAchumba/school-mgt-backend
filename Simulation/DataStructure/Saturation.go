package datastructure

type SaturationData struct {
	So_old float64
	Sw_old float64
	Sg_old float64
	So_new float64
	Sw_new float64
	Sg_new float64
	So_itr float64
	Sw_itr float64
	Sg_itr float64
}

func NewSaturationData() SaturationData {
	return SaturationData{}
}
