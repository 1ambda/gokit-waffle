package user

import (
	"github.com/1ambda/gokit-waffle/waffle-server/service/common"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func NewUserListEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(UserListRequest)

		res := UserListResponse{
			Users:       svc.Users(),
			ErrResponse: *common.NewErrResponse(nil),
		}

		return res, nil
	}
}
