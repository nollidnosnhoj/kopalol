package uploads

import (
	"github.com/nollidnosnhoj/kopalol/internal/queries"
)

type FileUpload struct {
	queries.File
	Url   string
	Error error
}
