package number

import (
	"net/http"

	"github.com/1ambda/gokit-waffle/waffle-server/service/common"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func NewInsertHandler(ctx context.Context, svc NumberService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewInsertEndpoint(svc),
		DecodeInsertRequest,
		common.EncodeResponse,
	)
}

func NewQueryHandler(ctx context.Context, svc NumberService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewQueryEndpoint(svc),
		DecodeQueryRequest,
		common.EncodeResponse,
	)
}
