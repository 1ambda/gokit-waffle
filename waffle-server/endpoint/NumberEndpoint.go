package endpoint

import (
	"github.com/1ambda/gokit-waffle/waffle-server/message"
	"github.com/1ambda/gokit-waffle/waffle-server/service"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func CreateInsertionEndpoint(svc service.NumberService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(message.InsertRequest)

		msg, err := svc.Insert(req.User, req.Number)
		res := message.InsertResponse{
			Message: msg, Error: "",
		}

		if err != nil {
			res.Error = err.Error()
			return res, nil
		}

		return res, nil
	}
}
