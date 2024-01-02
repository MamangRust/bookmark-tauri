package response

type FileDataResponse struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Author  string `json:"author"`
	Content string `json:"content"`
}
