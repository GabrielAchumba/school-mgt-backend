package services

import (
	"context"
	"log"

	"github.com/GabrielAchumba/school-mgt-backend/common/config"
	vtpass "github.com/GabrielAchumba/school-mgt-backend/vtpass"
	vtu_ng "github.com/GabrielAchumba/school-mgt-backend/vtu_ng"
)

type DataBundleService interface {
	GetVariationCodes(serviceID string) (interface{}, error)
	PurchaseDataBundle(phone, networkID, amount string) (interface{}, error)
}

type DataBundleServiceImpl struct {
	ctx          context.Context
	vtpassClient *vtpass.Client
	vtngClient   *vtu_ng.Client
}

func NewDataBundleService(ctx context.Context, config config.Settings) DataBundleService {
	vtpassClient := vtpass.NewClient(nil)
	vtngClient := vtu_ng.NewClient(nil)

	return &DataBundleServiceImpl{
		ctx:          ctx,
		vtpassClient: vtpassClient,
		vtngClient:   vtngClient,
	}
}

func (impl DataBundleServiceImpl) GetVariationCodes(serviceID string) (interface{}, error) {

	log.Print("GetVariationCodes called")
	variationCodes, err := impl.vtpassClient.CommonService.GetVariation(serviceID)
	if err != nil {
		return nil, err
	}

	log.Print("GetVariationCodes completed")
	return variationCodes, err

}

func (impl DataBundleServiceImpl) PurchaseDataBundle(phone, networkID, amount string) (interface{}, error) {
	log.Print("PurchaseDataBundle called")
	variationCodes, err := impl.vtngClient.DataBundle.PurchaseDataBundle(phone, networkID, amount)
	if err != nil {
		return nil, err
	}

	log.Print("PurchaseDataBundle completed")
	return variationCodes, err

}
