package request

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type UpdateCategoryRequest struct {
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
