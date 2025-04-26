package response

import (
	"github.com/labstack/gommon/log"
	"net/http"
)

type BaseResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Error   []string    `json:"error,omitempty"`
	Data    interface{} `json:"data"`
}
type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Errors  []ErrorDetail `json:"errors,omitempty"`
}

type ErrorMapping struct {
	Code    int
	Message string
}

type ErrorHandlerOptions struct {
	LogInternalError bool
}

var DefaultErrorHandlerOptions = ErrorHandlerOptions{
	LogInternalError: true,
}

func HandleError(err error, mapping *ErrorMapping, options ErrorHandlerOptions) *ErrorResponse {
	if options.LogInternalError && err != nil {
		log.Errorf("unhandled error: %+v", err)
	}

	if mapping != nil {
		errResponse := ErrorResponse{
			Code:    mapping.Code,
			Message: mapping.Message,
		}
		return &errResponse
	}

	errResponse := ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}
	return &errResponse
}
