package client

import (
	"encoding/json"
	"net/http"

	"test-domain.com/math-expressions-client/models/dto"
)

func SendErrorsRequest(serviceAddress string) ([]dto.ErrorJson, error) {
	response, err := http.Get(serviceAddress + "/errors")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var result []dto.ErrorJson
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
