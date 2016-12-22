package number

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/1ambda/gokit-waffle/waffle-server/service/common"
	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func DecodeInsertRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req InsertRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func NewInsertHandler(ctx context.Context, svc NumberService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewInsertEndpoint(svc),
		DecodeInsertRequest,
		common.EncodeResponse,
	)
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

func NewQueryHandler(ctx context.Context, svc NumberService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewQueryEndpoint(svc),
		DecodeQueryRequest,
		common.EncodeResponse,
	)
}
