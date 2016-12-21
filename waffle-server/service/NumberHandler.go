package service

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func NewInsertHandler(ctx context.Context, svc NumberService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewInsertEndpoint(svc),
		DecodeInsertRequest,
		EncodeInsertResponse,
	)
}

func NewQueryHandler(ctx context.Context, svc NumberService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewQueryEndpoint(svc),
		DecodeQueryRequest,
		EncodeQueryResponse,
	)
}
