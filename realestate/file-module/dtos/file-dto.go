package dtos

import "time"

type CreateFileRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	FileUrl          string `json:"fileUrl"`
	FileName         string `json:"fileName"`
	OriginalFileName string `json:"originalFileName"`
	FileCategory     string `json:"fileCategory"`
	CategoryId       string `json:"categoryId"`
}

type UpdateFileRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	FileUrl          string `json:"fileUrl"`
	FileName         string `json:"fileName"`
	OriginalFileName string `json:"originalFileName"`
	FileCategory     string `json:"fileCategory"`
	CategoryId       string `json:"categoryId"`
}

type FileResponse struct {
	Id               string    `json:"id"  bson:"_id"`
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
