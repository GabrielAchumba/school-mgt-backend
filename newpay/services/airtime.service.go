package services

import (
	"context"
	"log"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	vtu_ng "github.com/GabrielAchumba/school-mgt-backend/vtu_ng"
)

type AirtimeService interface {
	PurchaseAirtime(phone, networkID, amount string) (interface{}, error)
}

type AirtimeServiceImpl struct {
	ctx        context.Context
	vtngClient *vtu_ng.Client
}

func NewAirtimeService(ctx context.Context, config config.Settings) AirtimeService {
	vtngClient := vtu_ng.NewClient(nil)

	return &AirtimeServiceImpl{
		ctx:        ctx,
		vtngClient: vtngClient,
	}
}

func (impl AirtimeServiceImpl) PurchaseAirtime(phone, networkID, amount string) (interface{}, error) {
	log.Print("PurchaseAirtime called")
	variationCodes, err := impl.vtngClient.Airtime.PurchaseAirtime(phone, networkID, amount)
	if err != nil {
		return nil, err
	}

	log.Print("PurchaseAirtime completed")
	return variationCodes, err

}
