package services

import (
	"context"

	"github.com/GabrielAchumba/school-mgt-backend/reservoir-simulation/modules/simulation-module/dtos"
	utils "github.com/GabrielAchumba/school-mgt-backend/reservoir-simulation/modules/simulation-module/utils"
)

type SimulationService interface {
	Run(inputData dtos.SimulationInputDTO) (interface{}, error)
}

type serviceImpl struct {
	ctx context.Context
}

func New(ctx context.Context) SimulationService {

	return &serviceImpl{
		ctx: ctx,
	}
}

func (impl serviceImpl) Run(inputData dtos.SimulationInputDTO) (interface{}, error) {

	SpaceDistributions, SimulationLogs := utils.Simulate(inputData)
	SimulatioOutput := dtos.SimulatioOutput{
		SpaceDistributions: SpaceDistributions,
		SimulationLogs:     SimulationLogs,
	}
	return SimulatioOutput, nil
}
