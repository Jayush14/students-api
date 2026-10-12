package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	statusCode int
	Error      string
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func ValidationError(errs validator.ValidationErrors) ErrorResponse {
	var errorMessages []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errorMessages = append(errorMessages, err.Field()+" is required")
		default:
			errorMessages = append(errorMessages, err.Field()+" is invalid")
		}
	}
	return ErrorResponse{
		statusCode: http.StatusBadRequest,
		Error:      strings.Join(errorMessages, ", "),
	}
}
