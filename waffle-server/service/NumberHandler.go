package service

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func NewInsertHandler(ctx context.Context, r NumberRepository) http.Handler {
	svc := NewNumberService(r)
	return httptransport.NewServer(
		ctx,
		NewInsertEndpoint(svc),
		DecodeInsertRequest,
		EncodeInsertResponse,
	)
}
