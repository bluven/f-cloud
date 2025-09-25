package errorx

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ErrorResponse struct {
	Error       string            `json:"error"`
	Detail      string            `json:"detail,omitempty"`
	FieldErrors map[string]string `json:"fieldErrors,omitempty"`
}

// ErrorHandler handles errors and returns the appropriate HTTP status code and response body.
func DefaultErrorHandler(err error) (int, any) {
	switch err := err.(type) {
	case *LogicError:
		return err.Code, err.ToErrorResponse()
	case validator.ValidationErrors:
		return http.StatusBadRequest, translateValidationErrors(err)
	default:
		code := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
		}

		return code, ErrorResponse{
			Error: err.Error(),
		}
	}
}

func translateValidationErrors(err validator.ValidationErrors) ErrorResponse {
	fieldErrors := make(map[string]string, len(err))
	for _, e := range err {
		fieldErrors[e.Field()] = e.Error()
	}

	return ErrorResponse{
		Error:       err.Error(),
		FieldErrors: fieldErrors,
	}
}
