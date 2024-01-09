package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseAuth struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func ResponseToken(w http.ResponseWriter, message string, token string, status int) {
	res := ResponseAuth{
		Status:  status,
		Message: message,
		Token:   token,
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}

func ResponseMessage(w http.ResponseWriter, message string, data interface{}, status int) {
	res := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}

func ResponseError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)

	res := Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func BadRequestError(w http.ResponseWriter, message string) {
	errResponse := ErrorResponse{Message: message}
	jsonResponse, _ := json.Marshal(errResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(jsonResponse)
}

func HandleValidationErrors(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case *validator.InvalidValidationError:
		BadRequestError(w, "validation failed: invalid validator")
	case validator.ValidationErrors:
		var errorMsg string
		for _, fieldErr := range e {
			errorMsg += fmt.Sprintf("field: %s, error: %s\n", fieldErr.Field(), fieldErr.Tag())
		}
		BadRequestError(w, "validation failed: "+errorMsg)
	default:
		BadRequestError(w, "unknown validation error")
	}
}
