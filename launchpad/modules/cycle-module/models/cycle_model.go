package models

type PsuedoLevel struct {
	Id    string `json:"id" bson:"_id"`
	Label string `json:"label,omitempty" binding:"required"`
	Icon  string `json:"icon,omitempty" binding:"required"`
}

type Cycle struct {
	Id           string `json:"id" bson:"_id"`
	CreatedDay   int    `json:"createdDay"`
	CreatedMonth int    `json:"createdMonth"`
	CreatedYear  int    `json:"createdYear"`

	Contributor_Id      string `json:"contributor_Id"`
	Cycle_Index         int    `json:"cycle_Index,omitempty" binding:"required"`
	IsThreshHoldReached bool   `json:"isThreshHoldReached,omitempty" binding:"required"`
	IsCompleted         bool   `json:"isCompleted,omitempty" binding:"required"`
	CycleName           string `json:"cycleName,omitempty" binding:"required"`
}

type Stream struct {
	Id       int           `json:"id"`
	Label    string        `json:"label,omitempty" binding:"required"`
	Children []PsuedoLevel `json:"children"`
	Avatar   string        `json:"avatar,omitempty" binding:"required"`
}
