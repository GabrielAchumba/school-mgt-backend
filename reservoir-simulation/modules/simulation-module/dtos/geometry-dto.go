package dtos

type DimensionsDTO struct {
	Nx float64 `json:"nx"`
	Ny float64 `json:"ny"`
	Nz float64 `json:"nz"`
}

type GriddingDTO struct {
	TypeOfGridding string    `json:"typeOfGridding"`
	DxVector       []float64 `json:"dxVector"`
	DyVector       []float64 `json:"dyVector"`
	DzVector       []float64 `json:"dzVector"`
}

type GeometryDTO struct {
	Dimensions DimensionsDTO `json:"dimensions"`
	Gridding   GriddingDTO   `json:"gridding"`
}
