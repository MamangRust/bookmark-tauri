package request

import "github.com/go-playground/validator/v10"

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

func (c *CreateCategoryRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}

func (c *UpdateCategoryRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(c)

	if err != nil {
		return err
	}

	return nil
}
