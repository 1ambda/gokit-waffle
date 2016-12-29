package number

import (
	. "github.com/1ambda/gokit-waffle/waffle-server/service/common"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func NewInsertEndpoint(svc NumberService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertRequest)

		msg, err := svc.Insert(req.User, req.Number)
		res := InsertResponse{
			Message:     msg,
			ErrResponse: *NewErrResponse(err),
		}

		return res, nil
	}
}
