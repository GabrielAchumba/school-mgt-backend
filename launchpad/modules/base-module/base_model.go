package basemodule

import "time"

type BaseModel struct {
	Id                   string    `json:"id" bson:"_id"`
	Base64String         string    `json:"base64String"`
	IsPhotographUploaded int       `json:"isPhotographUploaded"`
	CreatedDay           int       `json:"createdDay"`
	CreatedMonth         int       `json:"createdMonth"`
	CreatedYear          int       `json:"createdYear"`
	CreatedAt            time.Time `json:"createdAt"`
}
