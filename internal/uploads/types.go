package uploads

type FileUpload struct {
	ID               string
	FileExtension    string
	FileType         string
	FileName         string
	OriginalFileName string
	FileSize         int64
	Url              string
	Error            error
}
