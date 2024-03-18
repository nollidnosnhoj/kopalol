package uploads

import (
	"github.com/nollidnosnhoj/kopalol/queries"
)

type FileUpload struct {
	queries.File
	Url   string
	Error error
}
