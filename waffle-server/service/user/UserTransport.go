package user

import (
	"net/http"

	"github.com/1ambda/gokit-waffle/waffle-server/service/common"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func DecodeUserListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return UserListRequest{}, nil
}

func NewUserListHandler(ctx context.Context, svc UserService) http.Handler {
	return httptransport.NewServer(
		ctx,
		NewUserListEndpoint(svc),
		DecodeUserListRequest,
		common.EncodeResponse,
	)
}
