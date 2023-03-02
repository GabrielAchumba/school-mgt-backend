package dtos

type NextOfKinDTO struct {
	FullName        string `json:"fullname"`
	NOKNames        string `json:"nOKNames"`
	NOKAddress      string `json:"nOKAddress"`
	NOKPhoneNumber  string `json:"nOKPhoneNumber"`
	NOKRelationship string `json:"nOKRelationship"`
	ContributorId   string `json:"contributorId"`
}
