package service

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func CreateInsertionEndpoint(svc NumberService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertRequest)

		msg, err := svc.Insert(req.User, req.Number)
		res := InsertResponse{
			Message: msg, Error: "",
		}

		if err != nil {
			res.Error = err.Error()
			return res, nil
		}

		return res, nil
	}
}
