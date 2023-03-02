package models

type Category struct {
	Id            string `json:"id" bson:"_id"`
	CreatedDay    int    `json:"createdDay"`
	CreatedMonth  int    `json:"createdMonth"`
	CreatedYear   int    `json:"createdYear"`
	ContributorId string `json:"contributorId"`
	UserName      string `json:"username"`
	ParentId      string `json:"parentId"`
	FirstName     string `json:"firstName"`
	MiddleName    string `json:"middleName"`
	LastName      string `json:"lastName"`
	Gender        string `json:"gender"`
	BankName      string `json:"bankName"`
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
	BVN           string `json:"bVN"`

	NLevelOneRoomOneChildren   int `json:"nLevelOneRoomOneChildren"`
	NLevelTwoRoomOneChildren   int `json:"nLevelTwoRoomOneChildren"`
	NLevelThreeRoomOneChildren int `json:"nLevelThreeRoomOneChildren"`
	NLevelFourRoomOneChildren  int `json:"nLevelFourRoomOneChildren"`
	NLevelFiveRoomOneChildren  int `json:"nLevelFiveRoomOneChildren"`
	NLevelSixRoomOneChildren   int `json:"nLevelSixRoomOneChildren"`
	NLevelSevenRoomOneChildren int `json:"nLevelSevenRoomOneChildren"`

	IsNLevelOneRoomOneChildren   bool `json:"isNLevelOneRoomOneChildren"`
	IsNLevelTwoRoomOneChildren   bool `json:"isNLevelTwoRoomOneChildren"`
	IsNLevelThreeRoomOneChildren bool `json:"isNLevelThreeRoomOneChildren"`
	IsNLevelFourRoomOneChildren  bool `json:"isNLevelFourRoomOneChildren"`
	IsNLevelFiveRoomOneChildren  bool `json:"isNLevelFiveRoomOneChildren"`
	IsNLevelSixRoomOneChildren   bool `json:"isNLevelSixRoomOneChildren"`
	IsNLevelSevenRoomOneChildren bool `json:"isNLevelSevenRoomOneChildren"`
}
