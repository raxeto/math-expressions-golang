package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"test-domain.com/math-expressions-service/models"
	"test-domain.com/math-expressions-service/models/dto"
	"test-domain.com/math-expressions-service/models/logger"
)

func EvaluateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dto.EvaluateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	mathExpr := models.MathExpression{}
	mathExpr.SetExpression(req.Expression)

	result, err := mathExpr.Evaluate()

	response := dto.EvaluateResponse{}

	if err != nil {
		logger.GetInstance().Error(mathExpr.GetExpression(), "/evaluate", err)
		response.Result = err.Error()
	} else {
		response.Result = strconv.FormatFloat(result, 'f', -1, 64)
	}
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the response object as JSON and write it to the response
	encodeErr := json.NewEncoder(w).Encode(response)
	if encodeErr != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
