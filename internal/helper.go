package helper

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetAllProblems(domain string) (interface{}, error) {
	resp, err := http.Get(domain + "/api/problems/all/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
