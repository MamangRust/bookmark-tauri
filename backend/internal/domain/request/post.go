package request

type CreatePostRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID int    `json:"category_id"`
	UserID     int    `json:"user_id"`
}

type UpdatePostRequest struct {
	PostID     int    `json:"post_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID int    `json:"category_id"`
	UserID     int    `json:"user_id"`
}
