package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"test-domain.com/math-expressions-client/models/dto"
)

func SendValidateRequest(serviceAddress string, expression string) (bool, string) {
	reqBody, err := json.Marshal(dto.ValidateRequest{Expression: expression})
	if err != nil {
		return false, err.Error()
	}

	response, err := http.Post(serviceAddress+"/validate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return false, err.Error()
	}
	defer response.Body.Close()

	var result dto.ValidateResponse
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return false, err.Error()
	}

	return result.Valid, result.Reason
}
