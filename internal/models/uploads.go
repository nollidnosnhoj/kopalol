package models

import (
	"time"
)

type UploadFileResponse struct {
	Id          string    `json:"id"`
	ContentType string    `json:"content_type"`
	FileSize    int64     `json:"file_size"`
	Md5Hash     string    `json:"md5_hash"`
	DeletionKey string    `json:"deletion_key"`
	Url         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
}
