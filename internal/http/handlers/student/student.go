package student

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/Jayush14/students-api/internal/types"
	"github.com/Jayush14/students-api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
			return
		}

		err = validator.New().Struct(student)

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(err.(validator.ValidationErrors)))
			return
		}
		slog.Info("Received student data", "student", student)

		response.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Student data received successfully"})
	}
}
