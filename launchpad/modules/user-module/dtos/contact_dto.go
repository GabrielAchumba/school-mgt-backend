package dtos

type ContactDTO struct {
	FullName         string `json:"fullname"`
	Address          string `json:"address"`
	ResidentialCity  string `json:"residentialCity"`
	ResidentialState string `json:"residentialState"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phoneNumber"`
	ContributorId    string `json:"contributorId"`
	Base64String     string `json:"base64String"`
}
