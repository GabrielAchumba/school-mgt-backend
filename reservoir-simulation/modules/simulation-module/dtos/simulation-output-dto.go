package dtos

import DataStructure "github.com/GabrielAchumba/school-mgt-backend/Simulation/DataStructure"

type SimulatioOutput struct {
	SpaceDistributions []DataStructure.SpaceDistributions `json:"spaceDistributions"`
	SimulationLogs     string                             `json:"simulationLogs"`
}
