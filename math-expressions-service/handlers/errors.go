package handlers

import (
	"encoding/json"
	"net/http"

	"test-domain.com/math-expressions-service/models/logger"
)

func ErrorsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	errors := logger.GetInstance().GetAllErrors()

	w.Header().Set("Content-Type", "application/json")
	encodeErr := json.NewEncoder(w).Encode(errors)
	if encodeErr != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
