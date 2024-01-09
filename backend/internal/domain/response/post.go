package response

type PostResponse struct {
	ID      uint         `json:"ID"`
	Title   string       `json:"Title"`
	Content string       `json:"Content"`
	User    UserResponse `json:"User"`
}

type PostsResponse struct {
	ID      uint         `json:"ID"`
	Title   string       `json:"Title"`
	Content string       `json:"Content"`
	User    UserResponse `json:"User"`
}
