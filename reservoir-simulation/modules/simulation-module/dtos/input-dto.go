package dtos

type SimulationInputDTO struct {
	Geometry       GeometryDTO    `json:"geometry"`
	Rock           Rock           `json:"rock"`
	Pvt            Pvt            `json:"pvt"`
	Initialization Initialization `json:"initialization"`
	Schedule       Schedule       `json:"schedule"`
	Boundaries     Boundaries     `json:"boundaries"`
	Wells          []WellData     `json:"wells"`
}
