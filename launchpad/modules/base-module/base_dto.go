package basemodule

import "time"

type BaseDTO struct {
	CreatedAt            time.Time `json:"createdAt,omitempty"`
	Base64String         string    `json:"base64String,omitempty"`
	IsPhotographUploaded int64     `json:"isPhotographUploaded,omitempty"`
}
