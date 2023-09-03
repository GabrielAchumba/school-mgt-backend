package services

import (
	"context"
	"log"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	vtu_ng "github.com/GabrielAchumba/school-mgt-backend/vtu_ng"
)

type BalanceService interface {
	GetWalletBalance() (interface{}, error)
}

type BalanceServiceImpl struct {
	ctx        context.Context
	vtngClient *vtu_ng.Client
}

func NewBalanceService(ctx context.Context, config config.Settings) BalanceService {
	vtngClient := vtu_ng.NewClient(nil)

	return &BalanceServiceImpl{
		ctx:        ctx,
		vtngClient: vtngClient,
	}
}

func (impl BalanceServiceImpl) GetWalletBalance() (interface{}, error) {
	log.Print("PurchaseDataBundle called")
	variationCodes, err := impl.vtngClient.Balance.GetWalletBalance()
	if err != nil {
		return nil, err
	}

	log.Print("PurchaseDataBundle completed")
	return variationCodes, err

}
