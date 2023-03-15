package datastructure

type RockData struct {
	Kx                  float64
	Ky                  float64
	Kz                  float64
	Porosity            float64
	RockCompressibility float64
}

func NewRockData() RockData {
	return RockData{}
}
