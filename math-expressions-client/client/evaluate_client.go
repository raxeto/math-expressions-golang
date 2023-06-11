package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"test-domain.com/math-expressions-client/models/dto"
)

func SendEvaluateRequest(serviceAddress string, expression string) string {
	reqBody, err := json.Marshal(dto.EvaluateRequest{Expression: expression})
	if err != nil {
		return err.Error()
	}

	response, err := http.Post(serviceAddress+"/evaluate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err.Error()
	}
	defer response.Body.Close()

	var result dto.EvaluateResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return err.Error()
	}

	return result.Result
}
