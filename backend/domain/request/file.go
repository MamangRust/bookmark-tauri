package request

type CreateFileRequest struct {
	Folder  string `json:"folder"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateFileRequest struct {
	Folder  string `json:"folder"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type FileRequest struct {
	Folder   string `json:"folder"`
	FileName string `json:"fileName"`
}
