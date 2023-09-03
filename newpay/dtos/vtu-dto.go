package dtos

type VTURequest struct {
	ServiceID       string `json:"serviceID"`
	CustomerID      string `json:"customerID"`
	Phone           string `json:"phone"`
	SmartCardNumber string `json:"smartCardNumber"`
	VariationID     string `json:"variationID"`
	MeterNumber     string `json:"meterNumber"`
	Amount          string `json:"amount"`
	NetworkID       string `json:"networkID"`
}
