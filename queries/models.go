// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package queries

import (
	"time"
)

type File struct {
	ID               string
	FileExtension    string
	FileType         string
	FileName         string
	OriginalFileName string
	FileSize         int64
	Md5Hash          string
	DeletionKey      string
	CreatedAt        time.Time
}
