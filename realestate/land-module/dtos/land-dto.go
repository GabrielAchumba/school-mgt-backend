package dtos

import (
	"time"

	fileDTOsPackage "github.com/GabrielAchumba/school-mgt-backend/realestate/file-module/dtos"
)

type CreateLandRequest struct {
	Title          string `json:"title"`
	WholePlot      string `json:"wholePlot"`
	FractionPlot   string `json:"fractionPlot"`
	PartialAddress string `json:"partialAddress"`
}

type UpdateLandRequest struct {
	Title          string `json:"title"`
	WholePlot      string `json:"wholePlot"`
	FractionPlot   string `json:"fractionPlot"`
	PartialAddress string `json:"partialAddress"`
}

type LandResponse struct {
	Id             string                         `json:"id"  bson:"_id"`
	CreatedAt      time.Time                      `json:"createdAt"`
	CreatedBy      string                         `json:"createdBy"`
	Title          string                         `json:"title"`
	WholePlot      string                         `json:"wholePlot"`
	FractionPlot   string                         `json:"fractionPlot"`
	UserPictureUrl string                         `json:"userPictureUrl"`
	FullName       string                         `json:"fullName"`
	PhoneNumber    string                         `json:"phoneNumber"`
	DateOfCreation string                         `json:"dateOfCreation"`
	Files          []fileDTOsPackage.FileResponse `json:"files"`
	PartialAddress string                         `json:"partialAddress"`
}

type LandResponsePaginated struct {
	TotalNumberOfUsers int            `json:"totalNumberOfUsers"`
	PaginatedLands     []LandResponse `json:"paginatedLands"`
	Limit              int            `json:"limit"`
}
