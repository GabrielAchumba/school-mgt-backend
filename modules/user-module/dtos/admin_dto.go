package dtos

type AdminDTO struct {
	ContributorId string `json:"contributorId"`
	Base64String  string `json:"base64String"`
	FirstName     string `json:"firstName" `
	LastName      string `json:"lastName" `
	UserType      string `json:"userType" `
	Designation   string `json:"designation" `
	UserName      string `json:"userName"`
}
