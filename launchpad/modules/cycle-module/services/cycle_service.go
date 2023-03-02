package services

import (
	"context"

	"github.com/GabrielAchumba/school-mgt-backend/launchpad/helpers"
	"github.com/GabrielAchumba/school-mgt-backend/launchpad/modules/cycle-module/models"
)

type CycleService interface {
	GetCyclesWithLevelsByUserId() []models.Stream
	GetROIs() helpers.ReturnOnInvestment
}

type serviceImpl struct {
	CyclesUtils helpers.CyclesUtils
	ctx         context.Context
}

func New(ctx context.Context) CycleService {

	return &serviceImpl{
		ctx: ctx,
	}
}

func (impl serviceImpl) GetCyclesWithLevelsByUserId() []models.Stream {

	streams := impl.CyclesUtils.GetListOfStream()

	return streams
}

func (impl serviceImpl) GetROIs() helpers.ReturnOnInvestment {

	rOIs := helpers.ReturnOnInvestment{
		N500ROIs:   helpers.ROIs[0],
		N1000ROIs:  helpers.ROIs[1],
		N2000ROIs:  helpers.ROIs[2],
		N5000ROIs:  helpers.ROIs[3],
		N10000ROIs: helpers.ROIs[4],
	}

	return rOIs
}
