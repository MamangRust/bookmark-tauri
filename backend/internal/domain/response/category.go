package response

type CategoriesResponse struct {
	ID          uint           `json:"ID"`
	Name        string         `json:"Name"`
	Image       string         `json:"Image"`
	Description string         `json:"Description"`
	Posts       []PostResponse `json:"Posts"`
}

type CategoryResponse struct {
	ID          uint   `json:"ID"`
	Name        string `json:"Name"`
	Image       string `json:"Image"`
	Description string `json:"Description"`
}
