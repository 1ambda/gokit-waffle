package service

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"golang.org/x/net/context"
)

func DecodeInsertRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req InsertRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, res interface{}) error {
	return json.NewEncoder(w).Encode(res)
}

func DecodeQueryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	u, ok := vars["user"]
	if !ok {
		// TODO error list
		return nil, errors.New("Bad Request for DecodeQueryRequest")
	}

	return QueryRequest{User: u}, nil
}
