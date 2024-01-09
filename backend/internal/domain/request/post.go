package request

import "github.com/go-playground/validator/v10"

type CreatePostRequest struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID int    `json:"category_id"`
}

type UpdatePostRequest struct {
	PostID     int    `json:"post_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID int    `json:"category_id"`
}

func (request *CreatePostRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(request)

	if err != nil {
		return err
	}

	return nil
}

func (request *UpdatePostRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(request)

	if err != nil {
		return err
	}

	return nil
}
