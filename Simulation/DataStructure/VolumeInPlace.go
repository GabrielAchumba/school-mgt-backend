package datastructure

type VolumeInPlace struct {
	Oil   float64
	Water float64
	Gas   float64
}

func NewVolumeInPlace() VolumeInPlace {
	return VolumeInPlace{}
}
