package services

import (
	"context"
	"log"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	vtpass "github.com/GabrielAchumba/school-mgt-backend/vtpass"
	vtu_ng "github.com/GabrielAchumba/school-mgt-backend/vtu_ng"
)

type CableTVService interface {
	GetVariationCodes(serviceID string) (interface{}, error)
	SmartCardNumberVerification(customerID string, serviceID string) (interface{}, error)
	PurchaseCableTV(phone, serviceID, smartCardNumber, variationID string) (interface{}, error)
}

type CableTVServiceImpl struct {
	ctx          context.Context
	vtpassClient *vtpass.Client
	vtngClient   *vtu_ng.Client
}

func NewCableTVService(ctx context.Context, config config.Settings) CableTVService {
	vtpassClient := vtpass.NewClient(nil)
	vtngClient := vtu_ng.NewClient(nil)

	return &CableTVServiceImpl{
		ctx:          ctx,
		vtpassClient: vtpassClient,
		vtngClient:   vtngClient,
	}
}

func (impl CableTVServiceImpl) GetVariationCodes(serviceID string) (interface{}, error) {

	log.Print("GetVariationCodes called")
	variationCodes, err := impl.vtpassClient.CommonService.GetVariation(serviceID)
	if err != nil {
		return nil, err
	}

	log.Print("GetVariationCodes completed")
	return variationCodes, err

}

func (impl CableTVServiceImpl) SmartCardNumberVerification(customerID string, serviceID string) (interface{}, error) {
	log.Print("SmartCardNumberVerification called")
	variationCodes, err := impl.vtngClient.CustomerVerification.SmartCardNumberVerification(customerID, serviceID)
	if err != nil {
		return nil, err
	}

	log.Print("SmartCardNumberVerification completed")
	return variationCodes, err

}

func (impl CableTVServiceImpl) PurchaseCableTV(phone, serviceID, smartCardNumber, variationID string) (interface{}, error) {
	log.Print("PurchaseCableTV called")
	variationCodes, err := impl.vtngClient.CableTV.PurchaseCableTV(phone, serviceID, smartCardNumber, variationID)
	if err != nil {
		return nil, err
	}

	log.Print("PurchaseCableTV completed")
	return variationCodes, err

}
