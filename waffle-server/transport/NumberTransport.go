package transport

import (
	"encoding/json"
	"net/http"

	"github.com/1ambda/gokit-waffle/waffle-server/message"

	"golang.org/x/net/context"
)

func DecodeGenerateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req message.GenerateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func EncodeGenerateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
