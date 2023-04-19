package models

import "time"

type Land struct {
	CreatedAt      time.Time `json:"createdAt"`
	CreatedBy      string    `json:"createdBy"`
	Title          string    `json:"title"`
	WholePlot      string    `json:"wholePlot"`
	FractionPlot   string    `json:"fractionPlot"`
	PartialAddress string    `json:"partialAddress"`
}
