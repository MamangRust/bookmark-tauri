package request

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type RegisterUserRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (r *RegisterUserRequest) Validate() error {

	if r.Password != r.ConfirmPassword {
		return errors.New("password and confirm_password do not match")
	}

	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil

}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *LoginUserRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil
}
