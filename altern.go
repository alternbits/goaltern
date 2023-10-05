package altern

import (
	"encoding/json"
	"net/http"
)

const APIEndpoint = "https://api.altern.ai/v1/info.json"

type ApiResponse struct {
	Data   map[string]string `json:"data"`
	Status string            `json:"status"`
}

func IsStatusSuccess() (bool, error) {
	resp, err := http.Get(APIEndpoint)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var response ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return false, err
	}

	return response.Status == "success", nil
}
