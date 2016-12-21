package service

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

func NewInsertHandler(ctx context.Context, svc NumberService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewInsertEndpoint(svc),
		DecodeInsertRequest,
		EncodeResponse,
	)
}

func NewQueryHandler(ctx context.Context, svc NumberService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewQueryEndpoint(svc),
		DecodeQueryRequest,
		EncodeResponse,
	)
}

func NewNumberRouter(ctx context.Context) *mux.Router {
	numRepo := NewNumberRepository()
	numSvc := NewNumberService(numRepo)
	numRoute := mux.NewRouter()

	numInsertHandler := NewInsertHandler(ctx, numSvc)
	numRoute.Handle("/api/v1/number/insert", numInsertHandler).Methods("POST")
	numQueryHandler := NewQueryHandler(ctx, numSvc)
	numRoute.Handle("/api/v1/number/query/{user}", numQueryHandler).Methods("GET")

	return numRoute
}
