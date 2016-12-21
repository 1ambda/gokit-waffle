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

func DecodeQueryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	u, ok := vars["user"]
	if !ok {
		// TODO error list
		return nil, errors.New("Bad Request for DecodeQueryRequest")
	}

	return QueryRequest{User: u}, nil
}

type HasError interface {
	error() error
}

func EncodeNumberError(_ context.Context, err error, w http.ResponseWriter) {
	// TODO: switch errors

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func EncodeNumberResponse(ctx context.Context, w http.ResponseWriter, res interface{}) error {
	if e, hasError := res.(HasError); hasError && e.error() != nil {
		EncodeNumberError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(res)
}
