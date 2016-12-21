package service

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func NewInsertHandler(ctx context.Context) http.Handler {
	svc := numberServiceInst{}
	return httptransport.NewServer(
		ctx,
		NewInsertEndpoint(svc),
		DecodeInsertRequest,
		EncodeInsertResponse,
	)
}
