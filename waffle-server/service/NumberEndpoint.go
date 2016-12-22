package service

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

func NewInsertEndpoint(svc NumberService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertRequest)

		msg, err := svc.Insert(req.User, req.Number)
		res := InsertResponse{Message: msg, ErrResponse: *NewErrResponse(err)}

		return res, nil
	}
}

func NewQueryEndpoint(svc NumberService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(QueryRequest)

		total, err := svc.Query(req.User)
		res := QueryResponse{User: req.User, Total: total, ErrResponse: *NewErrResponse(err)}

		// TODO: test return error
		return res, nil
	}
}
