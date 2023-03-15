package dtos

type Pvt struct {
	TypeOfPVT string `json:"typeOfPVT"`
	Oil       Oil    `json:"oil"`
	Gas       Gas    `json:"gas"`
	Water     Water  `json:"water"`
}

type Oil struct {
	Pressure        []float64 `json:"pressure"`
	Viscosity       []float64 `json:"viscosity"`
	FVF             []float64 `json:"fVF"`
	SolutionGOR     []float64 `json:"solutionGOR"`
	Density         []float64 `json:"density"`
	Compressibility []float64 `json:"compressibility"`
}

type Gas struct {
	Pressure        []float64 `json:"pressure"`
	Viscosity       []float64 `json:"viscosity"`
	FVF             []float64 `json:"fVF"`
	Density         []float64 `json:"density"`
	Compressibility []float64 `json:"compressibility"`
}

type Water struct {
	Pressure        []float64 `json:"pressure"`
	Viscosity       []float64 `json:"viscosity"`
	FVF             []float64 `json:"fVF"`
	Density         []float64 `json:"density"`
	Compressibility []float64 `json:"compressibility"`
}
