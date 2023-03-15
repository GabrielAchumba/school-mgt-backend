package dtos

type Porosity struct {
	TypeOfPorositing string    `json:"typeOfPorositing"`
	PorosityArray    []float64 `json:"porosityArray"`
}

type Permeability struct {
	TypeOfPermeability string    `json:"typeOfPermeability"`
	PermeabilityXArray []float64 `json:"permeabilityXArray"`
	PermeabilityYArray []float64 `json:"permeabilityYArray"`
	PermeabilityZArray []float64 `json:"permeabilityZArray"`
}

type Compressibility struct {
	TypeOfCompressibility string    `json:"typeOfCompressibility"`
	CompressibilityArray  []float64 `json:"compressibilityArray"`
}

type Rock struct {
	Porosity        Porosity        `json:"porosity"`
	Permeability    Permeability    `json:"permeability"`
	Compressibility Compressibility `json:"compressibility"`
}
