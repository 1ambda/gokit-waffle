package service

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func CreateInsertHandler(ctx context.Context) http.Handler {
	svc := NumberServiceInst{}
	return httptransport.NewServer(
		ctx,
		CreateInsertionEndpoint(svc),
		DecodeInsertRequest,
		EncodeInsertResponse,
	)
}
