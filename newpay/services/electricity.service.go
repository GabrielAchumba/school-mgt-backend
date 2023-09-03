package services

import (
	"context"
	"log"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	vtpass "github.com/GabrielAchumba/school-mgt-backend/vtpass"
	vtu_ng "github.com/GabrielAchumba/school-mgt-backend/vtu_ng"
)

type ElectricityService interface {
	MeterNumberVerification(customerID, VariationID, serviceID string) (interface{}, error)
	PurchaseElectricity(phone, meterNumber, serviceID, variationID, amount string) (interface{}, error)
}

type ElectricityServiceImpl struct {
	ctx          context.Context
	vtpassClient *vtpass.Client
	vtngClient   *vtu_ng.Client
}

func NewElectricityService(ctx context.Context, config config.Settings) ElectricityService {
	vtpassClient := vtpass.NewClient(nil)
	vtngClient := vtu_ng.NewClient(nil)

	return &ElectricityServiceImpl{
		ctx:          ctx,
		vtpassClient: vtpassClient,
		vtngClient:   vtngClient,
	}
}

func (impl ElectricityServiceImpl) MeterNumberVerification(customerID, VariationID, serviceID string) (interface{}, error) {
	log.Print("MeterNumberVerification called")
	variationCodes, err := impl.vtngClient.CustomerVerification.MeterNumberVerification(customerID, serviceID, VariationID)
	if err != nil {
		return nil, err
	}

	log.Print("MeterNumberVerification completed")
	return variationCodes, err

}

func (impl ElectricityServiceImpl) PurchaseElectricity(phone, meterNumber, serviceID, variationID, amount string) (interface{}, error) {
	log.Print("PurchaseElectricity called")
	variationCodes, err := impl.vtngClient.Electricity.PurchaseElectricity(phone, meterNumber, serviceID,
		variationID, amount)
	if err != nil {
		return nil, err
	}

	log.Print("PurchaseElectricity completed")
	return variationCodes, err

}
