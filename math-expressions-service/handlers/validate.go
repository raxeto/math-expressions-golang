package handlers

import (
	"encoding/json"
	"net/http"

	"test-domain.com/math-expressions-service/models"
	"test-domain.com/math-expressions-service/models/dto"
	"test-domain.com/math-expressions-service/models/logger"
)

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dto.ValidateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	mathExpr := models.MathExpression{}
	mathExpr.SetExpression(req.Expression)

	valid, err := mathExpr.Validate()

	response := dto.ValidateResponse{
		Valid: valid,
	}

	if err != nil {
		response.Reason = err.Error()
		logger.GetInstance().Error(mathExpr.GetExpression(), "/validate", err)
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the response object as JSON and write it to the response
	encodeErr := json.NewEncoder(w).Encode(response)
	if encodeErr != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
