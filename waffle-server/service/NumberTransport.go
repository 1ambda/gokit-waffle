package service

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
)

func DecodeInsertRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req InsertRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func EncodeInsertResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
