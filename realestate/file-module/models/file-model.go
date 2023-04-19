package models

import "time"

type File struct {
	CreatedAt        time.Time `json:"createdAt"`
	CreatedBy        string    `json:"createdBy"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	FileUrl          string    `json:"fileUrl"`
	FileName         string    `json:"fileName"`
	OriginalFileName string    `json:"originalFileName"`
	FileCategory     string    `json:"fileCategory"`
	CategoryId       string    `json:"categoryId"`
}
