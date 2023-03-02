package dtos

type BioDataDTO struct {
	FullName             string `json:"fullname"`
	Base64String         string `json:"base64String"`
	IsPhotographUploaded int    `json:"isPhotographUploaded"`
	BloodGroup           string `json:"bloodGroup"`
	Genotype             string `json:"genotype"`
	MaritalStatus        string `json:"maritalStatus"`
	LGAOfOrigin          string `json:"lGAOfOrigin"`
	StateOfOrigin        string `json:"stateOfOrigin"`
	Country              string `json:"country"`
	ContributorId        string `json:"contributorId"`
}
